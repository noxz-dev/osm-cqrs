package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"noxz.dev/changeset-watcher/config"
	"noxz.dev/changeset-watcher/statistics"
	"os"
	"strconv"
	"sync"
	"time"

	gzip "github.com/klauspost/pgzip"
	"github.com/nats-io/nats.go"
	"github.com/withmandala/go-log"

	"noxz.dev/changeset-watcher/types"
	"noxz.dev/changeset-watcher/utils"
)

var logger = log.New(os.Stderr)
var statistic = statistics.NewStatistic("watcher-statistics.csv",
	config.NumberOfIncomingElements,
	config.DurationChangSetDownload,
	config.DurationNodesReloading,
	config.NumberOfReloadedNodes,
	config.NumberOfPublishedElements,
	config.NumberOfPublishedRoutingElements,
	config.DurationForRoutesFiltering,
	config.NumberOfPublishedSearchElements,
	config.DurationForSearchFiltering)

func main() {
	defer statistic.Close()
	var url string

	url = os.Getenv("NATS_IP")

	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	defer nc.Close()

	if err != nil {
		logger.Errorf("Failed to connect to the NATS-Server: \n%s \n", err.Error())
		return
	}

	var oldSeq = 0

	for {
		statistic.BeginnColum()
		resp, err := http.Get(config.OsmMinuteReplicationStateURL)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body, _ := io.ReadAll(resp.Body)
		stringBody := string(body)

		seq, err := utils.ExtractSeqNumber(&stringBody)

		if oldSeq >= seq {
			logger.Info("no new sequence number found... waiting for ", config.SequenceNumberPollingInterval, " sec")
			time.Sleep(config.SequenceNumberPollingInterval * time.Second)
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
		statistic.StartTimer(config.DurationChangSetDownload)
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
		statistic.StopTimerAndSetDuration(config.DurationChangSetDownload)
		logger.Info("parsing xml ...")
		osm := types.OsmChange{}
		err = xml.Unmarshal(body, &osm)
		if err != nil {
			logger.Error(err)
			return
		}
		osmNormalized := osm.Normalize()
		statistic.SetValue(config.NumberOfIncomingElements, strconv.Itoa(osmNormalized.Size()))
		logger.Info("reloading missing nodes referenced by ways...")
		statistic.StartTimer(config.DurationNodesReloading)
		reloaded, err := osmNormalized.Reload()
		statistic.StopTimerAndSetDuration(config.DurationNodesReloading)
		statistic.SetValue(config.NumberOfReloadedNodes, strconv.Itoa(reloaded))
		if err != nil {
			logger.Error("error while reloading missing nodes: ", err.Error())
			err = nil
		}
		logger.Info("missing nodes reloaded")
		wg := new(sync.WaitGroup)
		wg.Add(3)
		go sendAllChangesets(nc, osmNormalized, wg)
		go sendRoutingChangesets(nc, osmNormalized, wg)
		go sendSearchChangesets(nc, osmNormalized, wg)
		wg.Wait()
		_ = statistic.EndColum()
	}
}

func sendSearchChangesets(nc *nats.Conn, normalized types.OsmChangeNormalized, wg *sync.WaitGroup) {
	defer wg.Done()
	err := statistic.StartTimer(config.DurationForSearchFiltering)
	if err != nil {
		logger.Error(err.Error())
	}
	searchPayload := generateSearchEventPayload(normalized)
	publishEvent(nc, config.SearchSubject, searchPayload)
	statistic.StopTimerAndSetDuration(config.DurationForSearchFiltering)
	statistic.SetValue(config.NumberOfPublishedSearchElements, strconv.Itoa(searchPayload.Size()))

}

func sendRoutingChangesets(nc *nats.Conn, normalized types.OsmChangeNormalized, wg *sync.WaitGroup) {
	defer wg.Done()
	statistic.StartTimer(config.DurationForRoutesFiltering)
	streets := normalized.Filter([]types.NodeFilter{}, []types.WayFilter{types.NewWayFilter("highway")})
	publishEvent(nc, config.RoutingSubject, streets)
	statistic.StopTimerAndSetDuration(config.DurationForRoutesFiltering)
	statistic.SetValue(config.NumberOfPublishedRoutingElements, strconv.Itoa(streets.Size()))
}

func sendAllChangesets(nc *nats.Conn, normalized types.OsmChangeNormalized, wg *sync.WaitGroup) {
	defer wg.Done()
	publishEvent(nc, config.AllSubject, normalized)
	statistic.SetValue(config.NumberOfPublishedElements, strconv.Itoa(normalized.Size()))
}

func publishEvent(nc *nats.Conn, subject string, payload interface{}) {
	event, err := utils.CreateEvent("ChangesetWatcher", payload, subject)
	if err != nil {
		logger.Error("cloudevents wrapper could not be created: ", err.Error())
		return
	}
	bytes, err := json.Marshal(event)
	if err != nil {
		logger.Error("Event could not be serialized", err.Error())
		return
	}
	logger.Info("publishing new changeset to " + subject + " ...")
	err = nc.Publish(subject, bytes)
	if err != nil {
		logger.Error("failed to publish new change set to subject ["+subject+"]: ", err.Error())
	}
}

func generateSearchEventPayload(normalized types.OsmChangeNormalized) types.SearchPayload {
	//TODO: Nodes-Filter nach Vorstellung entfernen. Ist hier nicht notwendig.
	buildings := normalized.Filter(
		[]types.NodeFilter{
			types.NewNodeFilter("building", "name"),
			types.NewNodeFilter("building", "addr:street"),
			types.NewNodeFilter("amenity", "name"),
			types.NewNodeFilter("tourism", "name"),
		},
		[]types.WayFilter{
			types.NewWayFilter("building", "name"),
			types.NewWayFilter("building", "addr:street"),
			types.NewWayFilter("building", "addr:housenumber"),
			types.NewWayFilter("amenity", "name"),
			types.NewWayFilter("tourism", "name"),
		})
	modifySearchPoints := reduceWaysToSearchPoints(buildings.Modify.Ways, append(buildings.Modify.Nodes, buildings.Reloaded.Nodes...))
	createSearchPoints := reduceWaysToSearchPoints(buildings.Create.Ways, append(buildings.Create.Nodes, buildings.Reloaded.Nodes...))
	deleteSearchPoints := reduceWaysToSearchPoints(buildings.Delete.Ways, append(buildings.Delete.Nodes, buildings.Reloaded.Nodes...))
	points := normalized.Filter(
		[]types.NodeFilter{
			types.NewNodeFilter("building", "name"),
			types.NewNodeFilter("addr:housenumber"),
			types.NewNodeFilter("addr:street"),
			types.NewNodeFilter("amenity", "name"),
			types.NewNodeFilter("tourism", "name"),
		},
		[]types.WayFilter{})
	modifySearchPoints = append(modifySearchPoints, reduceNodesToSearchPoints(points.Modify.Nodes)...)
	createSearchPoints = append(createSearchPoints, reduceNodesToSearchPoints(points.Create.Nodes)...)
	deleteSearchPoints = append(deleteSearchPoints, reduceNodesToSearchPoints(points.Delete.Nodes)...)

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

		name, err := way.GetTag("name")
		if err != nil {
			houseNumber, _ := way.GetTag("addr:housenumber")
			street, _ := way.GetTag("addr:street")
			city, _ := way.GetTag("addr:city")
			postcode, _ := way.GetTag("addr:postcode")
			name = fmt.Sprint(street, " ", houseNumber, ", ", postcode, ", ", city)
		}

		searchPoints = append(searchPoints, types.SearchPoint{
			Name:     name,
			Location: centroid,
			Tags:     way.Tags,
			Id:       fmt.Sprint("way_", way.Id),
		})

	}
	return searchPoints
}

func reduceNodesToSearchPoints(nodes []types.Node) []types.SearchPoint {
	searchPoints := make([]types.SearchPoint, 0)
	for _, node := range nodes {
		name, err := node.GetTag("name")
		if err != nil {
			houseNumber, _ := node.GetTag("addr:housenumber")
			street, _ := node.GetTag("addr:street")
			city, _ := node.GetTag("addr:city")
			postcode, _ := node.GetTag("addr:postcode")

			name = fmt.Sprint(street, " ", houseNumber, ", ", postcode, ", ", city)
		}
		location := types.Location{
			Lat: node.Lat,
			Lng: node.Lon,
		}
		searchPoints = append(searchPoints, types.SearchPoint{
			Name:     name,
			Location: location,
			Tags:     node.Tags,
			Id:       fmt.Sprint("node_", node.Id),
		})
	}

	return searchPoints

}
