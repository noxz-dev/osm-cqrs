package main

import (
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
	"noxz.dev/changeset-watcher/statistics"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	gzip "github.com/klauspost/pgzip"
	"github.com/nats-io/nats.go"
	"github.com/withmandala/go-log"

	"noxz.dev/changeset-watcher/types"
	"noxz.dev/changeset-watcher/utils"
)

var logger = log.New(os.Stderr)

func main() {
	var url string

	if url = os.Getenv("NATS_IP"); url == "" {
		url = nats.DefaultURL
	}
	nc, err := nats.Connect(url)
	defer nc.Close()

	if err != nil {
		logger.Fatalf("Failed to connect to the NATS-Server: \n%s \n", err.Error())
		return
	}

	/*
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
				//logIfFailing(stat.BeginnColum())
				logIfFailing(filterFromConfig(nc, config.PathOfFilterConfig, &changeset))
				//logIfFailing(stat.EndColum())
				logger.Info("Waiting for 3 seconds..")
				time.Sleep(time.Second * 3)
			}

			return
		}
	*/

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
		oldSeq = seq

		logger.Info("new sequence number:" + fmt.Sprint(seq) + " parsing....")

		url := utils.BuildChangeSetUrl(seq)

		logger.Info("fetching " + url)
		startTime := time.Now()
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

		logger.Info("READ FROM FILTER CONFIG")
		logIfFailing(filterFromConfig(nc, config.PathOfFilterConfig, &osmNormalized, startTime))
	}
}

func filterFromConfig(nc *nats.Conn, filename string, normalized *types.OsmChangeNormalized, startTime time.Time) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	jsonDecoder := json.NewDecoder(file)
	filterConfig := new(types.FilterConfig)
	err = jsonDecoder.Decode(filterConfig)
	if err != nil {
		return err
	}
	subjectCount := len(filterConfig.Subjects)
	wg := new(sync.WaitGroup)
	wg.Add(subjectCount)
	for _, subject := range filterConfig.Subjects {
		go startAsyncSpecificProcessing(nc, subject, normalized, wg, startTime)
	}
	wg.Wait()
	return nil
}

func startAsyncSpecificProcessing(nc *nats.Conn, subject types.Subject, normalized *types.OsmChangeNormalized, wg *sync.WaitGroup, startTime time.Time) {
	defer wg.Done()
	var subjectStat = statistics.NewStatistic("/watcher-config/"+subject.Name+".statistics.csv",
		config.DurationForFiltering,
		config.NumberOfPublishedElements,
		config.NumberOfIncomingElements,
		config.DurationTotal)
	defer subjectStat.Close()
	logIfFailing(subjectStat.BeginnColum())
	logIfFailing(subjectStat.SetValue(config.NumberOfIncomingElements, strconv.Itoa(normalized.Size())))
	logIfFailing(subjectStat.StartTimer(config.DurationForFiltering))
	publishedElementsCount := startSpecificProcessing(nc, subject, normalized)
	logIfFailing(subjectStat.SetValue(config.NumberOfPublishedElements, strconv.Itoa(publishedElementsCount)))
	logIfFailing(subjectStat.StopTimerAndSetDuration(config.DurationForFiltering))
	duration := time.Since(startTime)
	logIfFailing(subjectStat.SetValue(config.DurationTotal, strconv.FormatInt(duration.Milliseconds(), 10)))
	logIfFailing(subjectStat.EndColum())
}

func startSpecificProcessing(nc *nats.Conn, subject types.Subject, normalized *types.OsmChangeNormalized) (elementCount int) {
	if subject.ReduceToPoints {
		elementCount = applyFilterReduceAndSend(nc, subject, *normalized)
	} else {
		elementCount = applyFilterAndSend(nc, subject, *normalized)
	}
	return
}

func applyFilterReduceAndSend(nc *nats.Conn, subject types.Subject, normalized types.OsmChangeNormalized) int {
	//Filtering
	payload := reduceToSearchPoints(normalized, subject.NodeFilters, subject.WayFilters)
	//Sending
	if subject.Compress {
		payloadBytes, _ := json.Marshal(&payload)
		zippedBytes, _ := utils.Compress(payloadBytes)
		publishEvent(nc, subject.Name, zippedBytes, "application/gzip")
	} else {
		publishEvent(nc, subject.Name, payload, cloudevents.ApplicationJSON)
	}
	return payload.Size()
}

func applyFilterAndSend(nc *nats.Conn, subject types.Subject, normalized types.OsmChangeNormalized) int {
	//Filtering
	filtered := normalized.Filter(subject.NodeFilters, subject.WayFilters)

	//Generating Bytes
	var payloadBytes []byte
	var err error
	var contentTyp string

	switch subject.Format {
	case types.FormatXML:
		payloadBytes, err = filtered.ToXML()
		contentTyp = cloudevents.ApplicationXML
	case types.FormatJSON:
		fallthrough
	default:
		payloadBytes, err = filtered.ToJSON()
		contentTyp = cloudevents.ApplicationJSON
	}

	if err != nil {
		logger.Error(err.Error())
		return 0
	}
	//Sending
	if subject.Compress {
		zippedBytes, _ := utils.Compress(payloadBytes)
		publishEvent(nc, subject.Name, zippedBytes, "application/gzip")
	} else if subject.Format == types.FormatJSON || subject.Format == "" {
		publishEvent(nc, subject.Name, filtered, cloudevents.ApplicationJSON)
	} else {
		publishEvent(nc, subject.Name, payloadBytes, contentTyp)
	}
	return filtered.Size()
}

func reduceToSearchPoints(t types.OsmChangeNormalized, nodeFilters []types.NodeFilter, wayFilters []types.WayFilter) types.SearchPayload {
	filteredWays := t.Filter(nil, wayFilters)

	tmp := append(filteredWays.Create.Nodes, filteredWays.Modify.Nodes...)
	tmp = append(tmp, filteredWays.Delete.Nodes...)
	tmp = append(tmp, filteredWays.Reloaded.Nodes...)

	modifySearchPoints := reduceWaysToSearchPoints(filteredWays.Modify.Ways, tmp)
	createSearchPoints := reduceWaysToSearchPoints(filteredWays.Create.Ways, tmp)
	deleteSearchPoints := reduceWaysToSearchPoints(filteredWays.Delete.Ways, tmp)

	filteredNodes := t.Filter(nodeFilters, []types.WayFilter{})

	modifySearchPoints = append(modifySearchPoints, reduceNodesToSearchPoints(filteredNodes.Modify.Nodes)...)
	createSearchPoints = append(createSearchPoints, reduceNodesToSearchPoints(filteredNodes.Create.Nodes)...)
	deleteSearchPoints = append(deleteSearchPoints, reduceNodesToSearchPoints(filteredNodes.Delete.Nodes)...)

	payload := types.SearchPayload{
		Modify: modifySearchPoints,
		Create: createSearchPoints,
		Delete: deleteSearchPoints,
	}
	return payload

}

func reloadNodes(osmNormalized *types.OsmChangeNormalized) {
	logger.Info("reloading missing nodes referenced by ways...")
	_, err := osmNormalized.Reload()
	if err != nil {
		logger.Error("error while reloading missing nodes: ", err.Error())
		return
	}
	logger.Info("missing nodes reloaded")
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
