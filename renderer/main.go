package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"github.com/withmandala/go-log"
	"noxz.dev/tile-renderer/statistics"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var logger = log.New(os.Stderr)

const (
	ImposmDuration      = "duration of imposm update (in ms)"
	WriteDuration       = "duration of writing zip file to disk (in ms)"
	SizeOfIncomingEvent = "size of incoming event (in bytes)"
	CollectStatistics   = true
)

var stat = statistics.NewStatistic("/src/data/renderer.statistics.csv",
	ImposmDuration,
	WriteDuration,
	SizeOfIncomingEvent,
)

func main() {

	defer func(stat *statistics.Statistic) {
		err := stat.Close()
		if err != nil {
			logger.Fatalf("Error occurred while closing the statistics: %s", err.Error())
		}
	}(&stat)

	if !CollectStatistics {
		err := stat.Pause()
		if err != nil {
			logger.Fatalf("Error occurred while pausing the statistics: %s", err.Error())
			return
		}
	}

	url := os.Getenv("NATS_IP")

	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	defer nc.Close()

	logger.Info("Connecting to NATS-Server...")

	if err != nil {
		logger.Fatalf("Failed to connect to NATS-Server: \n%s \n", err.Error())
		return
	}

	logger.Info("Connection successful")

	_, err = nc.Subscribe("all", func(msg *nats.Msg) {
		logIfFailing(stat.BeginnColum())

		logger.Info("Received message")

		cloudEvent := cloudevents.NewEvent()

		_ = json.Unmarshal(msg.Data, &cloudEvent)

		xmlDataZipped := bytes.NewBuffer(cloudEvent.Data())

		logger.Infof("Event data size: %d", len(xmlDataZipped.Bytes()))

		logIfFailing(stat.SetValue(SizeOfIncomingEvent, strconv.Itoa(xmlDataZipped.Len())))

		logIfFailing(stat.StartTimer(WriteDuration))

		err := os.WriteFile("change.osc.gz", xmlDataZipped.Bytes(), 0644)

		logIfFailing(stat.StopTimerAndSetDuration(WriteDuration))

		if err != nil {
			logger.Fatalf("Could not write file %s", err.Error())
			return
		}

		RunImposmUpdate()

		logIfFailing(stat.EndColum())
	})
	if err != nil {
		logger.Fatalf("Error occurred while subscribing: \n%s", err.Error())
		return
	}

	select {}
}

// RunImposmUpdate writes the gzip osmChange file to the database via imposm
func RunImposmUpdate() {
	imposm := "imposm"
	imposmCommand := "diff"
	configCommand := "-config"
	configLocation := "/src/imposm/config.json"
	xmlFile := "./change.osc.gz"

	cmd := exec.Command(imposm, imposmCommand, configCommand, configLocation, xmlFile)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logIfFailing(stat.StartTimer(ImposmDuration))

	err := cmd.Run()

	if err != nil {
		logger.Fatalf(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	logIfFailing(stat.StopTimerAndSetDuration(ImposmDuration))

	logger.Info(out.String())

	elapsed := time.Since(start)

	logger.Infof("Running imposm took %s", elapsed)
}

func logIfFailing(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
