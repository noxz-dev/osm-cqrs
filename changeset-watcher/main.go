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

	"noxz.dev/changeset-watcher/types"
	"noxz.dev/changeset-watcher/utils"
)

var logger = log.New(os.Stderr)

func main() {

	var url string

	url = os.Getenv("NATS_IP")

	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
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
	streets := changeNormalized.Filter([]types.NodeFilter{}, []types.WayFilter{types.NewWayFilter("highway")})
	searchPayload := generateSearchEventPayload(changeNormalized)
	go publishEvent(nc, "all", changeNormalized)
	go publishEvent(nc, "routing", streets)
	go publishEvent(nc, "search", searchPayload)
}

func publishEvent(nc *nats.Conn, subject string, payload interface{}) {
	event := utils.CreateEvent("ChangesetWatcher", payload)
	bytes, err := json.Marshal(event)
	if err != nil {
		logger.Error("Event could not be serialized", err.Error())
		return
	}
	logger.Info("publishing new changeset to " + subject + " ...")
	nc.Publish(subject, bytes)
}

func generateSearchEventPayload(normalized types.OsmChangeNormalized) types.SearchPayload {
	buildings := normalized.Filter(
		[]types.NodeFilter{
			types.NewNodeFilter("building", "name"),
			types.NewNodeFilter("building", "addr:housenumber"),
			types.NewNodeFilter("amenity", "name"),
			types.NewNodeFilter("tourism", "name"),
		},
		[]types.WayFilter{
			types.NewWayFilter("building", "name"),
			types.NewWayFilter("building", "addr:housenumber"),
			types.NewWayFilter("amenity", "name"),
			types.NewWayFilter("tourism", "name"),
		})
	modifySearchPoints := reduceWaysToSearchPoints(buildings.Modify.Ways, append(buildings.Modify.Nodes, buildings.Reloaded.Nodes...))
	createSearchPoints := reduceWaysToSearchPoints(buildings.Create.Ways, append(buildings.Create.Nodes, buildings.Reloaded.Nodes...))
	deleteSearchPoints := reduceWaysToSearchPoints(buildings.Delete.Ways, append(buildings.Delete.Nodes, buildings.Reloaded.Nodes...))
	//TODO: Es werden nur die Wege zu `SearchPoints` umgewandelt. Die gefilterten Nodes müssen auch entsprechend umgewandelt werden!

	payload := types.SearchPayload{
		Modify: modifySearchPoints,
		Create: createSearchPoints,
		Delete: deleteSearchPoints,
	}

	return payload
}

func reduceWaysToSearchPoints(ways []types.Way, nodes []types.Node) []types.SearchPoint {
	searchPoints := make([]types.SearchPoint, 0)

	for _, way := range ways {
		wayNodes := make([]types.Node, 0)
		for _, nr := range way.NodeRefs {
			for _, n := range nodes {
				if n.Id == nr.Ref {
					wayNodes = append(nodes, n)
				}
			}
		}
		centroid := utils.CalculateCentroid(&wayNodes)

		name, _ := way.GetTag("name")

		searchPoints = append(searchPoints, types.SearchPoint{
			Name:     name,
			Location: centroid,
			Tags:     way.Tags,
			Id:       fmt.Sprint("way_", way.Id),
		})

	}
	return searchPoints
}
