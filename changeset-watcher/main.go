package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"noxz.dev/changeset-watcher/config"
	"noxz.dev/changeset-watcher/importer"
	"noxz.dev/changeset-watcher/statistics"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	gzip "github.com/klauspost/pgzip"
	"github.com/nats-io/nats.go"
	"github.com/withmandala/go-log"

	"noxz.dev/changeset-watcher/types"
	"noxz.dev/changeset-watcher/utils"
)

var logger = log.New(os.Stderr)
var stat = statistics.NewStatistic("watcher.statistics.csv",
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
	defer stat.Close()
	if !config.CollectStatistics {
		stat.Pause()
	}
	var url string

	if url := os.Getenv("NATS_IP"); url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	defer nc.Close()

	if err != nil {
		logger.Fatalf("Failed to connect to the NATS-Server: \n%s \n", err.Error())
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "--import" {
		if len(os.Args) != 3 {
			logger.Error("no pbf filepath specified")
			return
		}

		changesets, err := importer.Import(os.Args[2])
		if err != nil {
			logger.Fatal("import failed -", err.Error())
		}

		for _, changeset := range *changesets {
			logIfFailing(stat.BeginnColum())
			wg := new(sync.WaitGroup)
			wg.Add(3)
			go SendAllChangesets(nc, changeset, wg)
			go sendRoutingChangesets(nc, changeset, wg)
			go sendSearchChangesets(nc, changeset, wg)
			wg.Wait()
			logIfFailing(stat.EndColum())
		}

		return
	}

	var oldSeq = 0

	for {

		resp, err := http.Get(config.OsmMinuteReplicationStateURL)

		if err != nil {
			logger.Error(err.Error())
			err = nil
			logger.Info("try same http request again...")
			continue
		}

		body, _ := io.ReadAll(resp.Body)
		stringBody := string(body)

		seq, _ := utils.ExtractSeqNumber(&stringBody)

		if oldSeq >= seq {
			logger.Info("no new sequence number found... waiting for ", config.SequenceNumberPollingInterval, " sec")
			time.Sleep(config.SequenceNumberPollingInterval * time.Second)
			continue
		}
		logIfFailing(stat.BeginnColum())
		oldSeq = seq

		logger.Info("new sequence number:" + fmt.Sprint(seq) + " parsing....")

		url := utils.BuildChangeSetUrl(seq)

		logger.Info("fetching " + url)
		logIfFailing(stat.StartTimer(config.DurationChangSetDownload))
		resp, err = http.Get(url)

		if err != nil {
			logger.Error(err.Error())
			err = nil
			logger.Info("try same http request again...")
			resp, err = http.Get(url)
		}

		if err != nil {
			logger.Error(err.Error())
			continue
		}

		reader, err := gzip.NewReader(resp.Body)

		if err != nil {
			logger.Error(err.Error())
			continue
		}

		body, _ = io.ReadAll(reader)
		resp.Body.Close()
		logIfFailing(stat.StopTimerAndSetDuration(config.DurationChangSetDownload))
		logger.Info("parsing xml ...")
		osm := types.OsmChange{}
		err = xml.Unmarshal(body, &osm)
		if err != nil {
			logger.Error(err)
			continue
		}

		osmNormalized := osm.Normalize()
		osmNormalized.RemoveAllDuplicates()

		reloadNodes(&osmNormalized)

		wg := new(sync.WaitGroup)
		wg.Add(3)
		go SendAllChangesets(nc, osmNormalized, wg)
		go sendRoutingChangesets(nc, osmNormalized, wg)
		go sendSearchChangesets(nc, osmNormalized, wg)
		wg.Wait()
		logIfFailing(stat.EndColum())
	}
}

func reloadNodes(osmNormalized *types.OsmChangeNormalized) {
	logIfFailing(stat.SetValue(config.NumberOfIncomingElements, strconv.Itoa(osmNormalized.Size())))
	logger.Info("reloading missing nodes referenced by ways...")
	logIfFailing(stat.StartTimer(config.DurationNodesReloading))
	reloaded, err := osmNormalized.Reload()
	logIfFailing(stat.StopTimerAndSetDuration(config.DurationNodesReloading))
	logIfFailing(stat.SetValue(config.NumberOfReloadedNodes, strconv.Itoa(reloaded)))
	if err != nil {
		logger.Error("error while reloading missing nodes: ", err.Error())
		return
	}
	logger.Info("missing nodes reloaded")
}

func sendSearchChangesets(nc *nats.Conn, normalized types.OsmChangeNormalized, wg *sync.WaitGroup) {
	defer wg.Done()
	err := stat.StartTimer(config.DurationForSearchFiltering)
	if err != nil {
		logger.Error(err.Error())
	}
	searchPayload := generateSearchEventPayload(normalized)
	publishEvent(nc, config.SearchSubject, searchPayload, cloudevents.ApplicationJSON)
	logIfFailing(stat.StopTimerAndSetDuration(config.DurationForSearchFiltering))
	logIfFailing(stat.SetValue(config.NumberOfPublishedSearchElements, strconv.Itoa(searchPayload.Size())))

}

func sendRoutingChangesets(nc *nats.Conn, normalized types.OsmChangeNormalized, wg *sync.WaitGroup) {
	defer wg.Done()
	logIfFailing(stat.StartTimer(config.DurationForRoutesFiltering))

	logger.Info("Anzahl Nodes vor dem Routing Filter: ", len(normalized.Create.Nodes), len(normalized.Reloaded.Nodes))
	streets := normalized.Filter([]types.NodeFilter{}, []types.WayFilter{types.NewWayFilter("highway")})
	logger.Info("Anzahl Nodes nach dem Routing Filter: ", len(streets.Create.Nodes), len(streets.Reloaded.Nodes))

	createAction := types.Action{
		Nodes:     append(streets.Create.Nodes, streets.Reloaded.Nodes...),
		Ways:      append(streets.Create.Ways, streets.Reloaded.Ways...),
		Relations: append(streets.Create.Relations, streets.Reloaded.Relations...),
	}

	xmlContent := types.OsmChangeNormalizedXML{
		Create: createAction,
		Delete: streets.Delete,
		Modify: streets.Modify,
	}

	xmlData, err := xml.MarshalIndent(xmlContent, " ", "    ")
	if err != nil {
		logger.Error(err.Error())
	}
	xmlData = []byte(xml.Header + string(xmlData))

	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	_, err = w.Write(xmlData)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = w.Close()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	zippedBytes := b.Bytes()

	publishEvent(nc, config.RoutingSubject, zippedBytes, "text/plain")
	logIfFailing(stat.StopTimerAndSetDuration(config.DurationForRoutesFiltering))
	logIfFailing(stat.SetValue(config.NumberOfPublishedRoutingElements, strconv.Itoa(streets.Size())))
}

func SendAllChangesets(nc *nats.Conn, normalized types.OsmChangeNormalized, wg *sync.WaitGroup) {
	defer wg.Done()

	createAction := types.Action{
		Nodes:     append(normalized.Create.Nodes, normalized.Reloaded.Nodes...),
		Ways:      append(normalized.Create.Ways, normalized.Reloaded.Ways...),
		Relations: append(normalized.Create.Relations, normalized.Reloaded.Relations...),
	}

	xmlContent := types.OsmChangeNormalizedXML{
		Create: createAction,
		Modify: normalized.Modify,
		Delete: normalized.Delete,
	}

	xmlData, err := xml.MarshalIndent(xmlContent, " ", "    ")
	if err != nil {
		logger.Error(err.Error())
	}
	xmlData = []byte(xml.Header + string(xmlData))

	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	_, err = w.Write(xmlData)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = w.Close()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	zippedBytes := b.Bytes()

	publishEvent(nc, config.AllSubject, zippedBytes, "text/plain")
	logIfFailing(stat.SetValue(config.NumberOfPublishedElements, strconv.Itoa(normalized.Size())))

}

func publishEvent(nc *nats.Conn, subject string, payload interface{}, contentType string) {
	event, err := utils.CreateEvent("ChangesetWatcher", payload, subject, contentType)
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

func publishEventGziped(nc *nats.Conn, subject string, payload interface{}) {
	event, err := utils.CreateEvent("ChangesetWatcher", payload, subject, cloudevents.ApplicationXML)
	if err != nil {
		logger.Error("cloudevents wrapper could not be created: ", err.Error())
		return
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.Error("Event could not be serialized", err.Error())
		return
	}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	w.Write(eventBytes)
	w.Close()

	zippedBytes := b.Bytes()

	logger.Info("publishing new changeset to " + subject + " ...")
	err = nc.Publish(subject, zippedBytes)
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
					wayNodes = append(wayNodes, n)
				}
			}
		}

		centroid := utils.CalculateCentroid(&wayNodes)

		name, exists := way.GetTag("name")
		if !exists {
			name = way.GetAddressString()
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
		name, exists := node.GetTag("name")
		if !exists {
			name = node.GetAddressString()
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

func logIfFailing(err error) {
	if err != nil {
		logger.Error(err.Error())
	}

}
