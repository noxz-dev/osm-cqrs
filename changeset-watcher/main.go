package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	gzip "github.com/klauspost/pgzip"
	"github.com/nats-io/nats.go"

	"noxz.dev/changeset-watcher/types"
	"noxz.dev/changeset-watcher/utils"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	if err != nil {
		fmt.Printf("Failed to connect to NATS-Server: \n%s \n", err.Error())
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
			fmt.Println("same seq .. waiting for 10 sec")
			time.Sleep(10 * time.Second)
			continue
		}
		oldSeq = seq

		fmt.Println("new seq " + fmt.Sprint(seq) + " proceeding")

		url, err := utils.BuildChangeSetUrl(seq)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("fetching " + url)

		resp, err = http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader, err := gzip.NewReader(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body, _ = io.ReadAll(reader)

		fmt.Println("parsing xml ...")
		osm := types.OsmChange{}
		err = xml.Unmarshal(body, &osm)

		// for _, cs := range osm.ChageSets {
		sendNewChangesetNotifcation(nc, &osm)
		// }

		// fmt.Printf("%+v\n", osm.ChageSets)
		//changes.modify
	}
}

func sendNewChangesetNotifcation(nc *nats.Conn, change *types.OsmChange) {
	fmt.Println("publishing new changeset...")
	//changeSetBytes, _ := json.Marshal(change)
	modifyBytes, _ := json.Marshal(normalizeActionObject(change.Modify))

	/*
		extractByTag(change.Modify, "highway")
		extractByTag(change.Delete, "highway")
		extractByTag(change.Create, "highway")

		extractByTag(change.Modify, "building")

		//chnages.modify
		//chnages.*.ways.streets
		//chnages.*.way.buildings
		//chnages.*.way
		//chnages.*.ways.*
		//chnages.*.relations

		nc.Publish(rootEvent, changeSetBytes)
	*/
	nc.Publish(genSub(modifyEvent), modifyBytes)
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

func normalizeActionObject(actions []types.Action) types.Action {
	ways := make([]types.Way, 0)
	nodes := make([]types.Node, 0)
	relations := make([]types.Relation, 0)

	for _, action := range actions {
		ways = append(ways, action.Ways...)
		nodes = append(nodes, action.Nodes...)
		relations = append(relations, action.Relations...)
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

func genSub(names ...string) string {
	var subject = rootEvent

	for _, name := range names {
		subject += "." + name
	}

	return subject
}
