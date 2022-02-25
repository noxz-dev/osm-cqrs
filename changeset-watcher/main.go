package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	gzip "github.com/klauspost/pgzip"
	"github.com/nats-io/nats.go"
	"github.com/withmandala/go-log"

	"noxz.dev/changeset-watcher/config"
	"noxz.dev/changeset-watcher/types"
	"noxz.dev/changeset-watcher/utils"
)

var logger = log.New(os.Stderr)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	if err != nil {
		logger.Infof("Failed to connect to the NATS-Server: \n%s \n", err.Error())
		return
	}

	var oldSeq = 0

	for {

		resp, err := http.Get("https://planet.openstreetmap.org/replication/minute/state.txt")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body, _ := io.ReadAll(resp.Body)
		stringBody := string(body)

		seq, err := utils.ExtractSeqNumber(&stringBody)

		if oldSeq >= seq {
			logger.Info("no new sequence number found... waiting for 10 sec")
			time.Sleep(10 * time.Second)
			continue
		}
		oldSeq = seq

		logger.Info("new sequence number:" + fmt.Sprint(seq) + " parsing....")

		url, err := utils.BuildChangeSetUrl(seq)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		logger.Info("fetching " + url)

		resp, err = http.Get(url)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		reader, err := gzip.NewReader(resp.Body)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		body, _ = io.ReadAll(reader)

		logger.Info("parsing xml ...")
		osm := types.OsmChange{}
		err = xml.Unmarshal(body, &osm)
		if err != nil {
			logger.Error(err)
			return
		}

		sendNewChangesetNotifcation(nc, &osm)
		// fmt.Printf("%+v\n", osm.ChageSets)
	}
}

func sendNewChangesetNotifcation(nc *nats.Conn, change *types.OsmChange) {
	changeNormalized := change.Normalize()
	err := changeNormalized.Reload()
	if err != nil {
		logger.Error(err.Error())
	}
	streets := extractStreets(changeNormalized)
	utils.WriteObjectToFile(changeNormalized, "all.json")
	utils.WriteObjectToFile(streets, "streets.json")
	publishEvent(nc, utils.GenSubject(config.ModifyEvent), types.MODIFY_EVENT, changeNormalized)
}

func publishEvent(nc *nats.Conn, subject string, EventType string, payload interface{}) {
	event := utils.CreateEvent("ChangesetWatcher", EventType, payload)
	bytes, err := json.Marshal(event)
	if err != nil {
		logger.Error("Event could not be serialized", err.Error())
		return
	}
	logger.Info("publishing new changeset...")
	nc.Publish(subject, bytes)
}

func extractByTag(actions []types.Action, searchTag string) types.Action {
	ways := make([]types.Way, 0)
	nodes := make([]types.Node, 0)
	relations := make([]types.Relation, 0)
	for _, action := range actions {
		for _, way := range action.Ways {
			if hasTag(searchTag, way.Tags) {
				ways = append(ways, way)
			}
		}
		for _, node := range action.Nodes {
			if hasTag(searchTag, node.Tags) {
				nodes = append(nodes, node)
			}
		}
		for _, relation := range action.Relations {
			if hasTag(searchTag, relation.Tags) {
				relations = append(relations, relation)
			}
		}
	}
	return types.Action{
		Ways:      ways,
		Nodes:     nodes,
		Relations: relations,
	}

}

func hasTag(searchTag string, tags []types.Tag) bool {
	for _, tag := range tags {
		if tag.K == searchTag {
			return true
		}
	}
	return false
}

func extractStreets(normalized types.OsmChangeNormalized) (streets types.OsmChangeNormalized) {
	tagName := "highway"
	normalized.Create.FilterWays(tagName)
	normalized.Modify.FilterWays(tagName)
	normalized.Delete.FilterWays(tagName)

	usedNodes := make(map[int]struct{}, 0)
	normalized.Create.UsedNodes(&usedNodes)
	normalized.Delete.UsedNodes(&usedNodes)
	normalized.Modify.UsedNodes(&usedNodes)

	normalized.Create.RemoveUnusedNodes(usedNodes)
	normalized.Delete.RemoveUnusedNodes(usedNodes)
	normalized.Modify.RemoveUnusedNodes(usedNodes)
	normalized.Reloaded.RemoveUnusedNodes(usedNodes)

	return normalized
}
