package updater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"noxz.dev/routing/statistics"
	"os"
	"os/exec"
	"strconv"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"github.com/withmandala/go-log"
)

var logger = log.New(os.Stderr)

var stat = statistics.NewStatistic("./out-files/routing.event.statistics.csv",
	statistics.SizeOfIncomingEvent,
	statistics.WriteDuration,
	statistics.InsertDuration,
)

func RunUpdate() {
	defer stat.Close()

	if !statistics.CollectStatistics {
		stat.Pause()
	}

	url := os.Getenv("NATS_IP")

	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)

	logger.Info("Connecting to NATS-Server...")

	if err != nil {
		logger.Errorf("Failed to connect to NATS-Server: \n%s \n", err.Error())
		return
	}
	defer nc.Close()

	logger.Info("Connection successful")

	_, err = nc.Subscribe("all", func(msg *nats.Msg) {

		logIfFailing(stat.BeginnColum())

		logger.Info("Received message")

		cloudEvent := cloudevents.NewEvent()

		err = json.Unmarshal(msg.Data, &cloudEvent)

		if err != nil {
			return
		}

		zippedXMLData := bytes.NewBuffer(cloudEvent.Data())

		logger.Infof("Event data size: %d", strconv.Itoa(zippedXMLData.Len()))

		logIfFailing(stat.SetValue(statistics.SizeOfIncomingEvent, strconv.Itoa(zippedXMLData.Len())))
		logIfFailing(stat.StartTimer(statistics.WriteDuration))
		err := os.WriteFile("temp.osc.gz", zippedXMLData.Bytes(), 0644)
		logIfFailing(stat.StopTimerAndSetDuration(statistics.WriteDuration))
		if err != nil {
			logger.Errorf("Could not write file %s", err.Error())
			return
		}

		RunImposmUpdate()
		logIfFailing(stat.EndColum())
	})

	if err != nil {
		logger.Errorf("Error occured while subscribing: \n%s \n\n", err.Error())
		return
	}

	//block forever
	select {}
}

// RunImposmUpdate writes the gzip osmChange file to the database via imposm
func RunImposmUpdate() {
	imposm := "imposm"
	imposmCommand := "diff"
	configCommand := "-config"
	configLocation := "/src/imposm/config.json"
	xmlFile := "./temp.osc.gz"

	cmd := exec.Command(imposm, imposmCommand, configCommand, configLocation, xmlFile)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logIfFailing(stat.StartTimer(statistics.InsertDuration))

	err := cmd.Run()

	if err != nil {
		logger.Error(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	logIfFailing(stat.StopTimerAndSetDuration(statistics.InsertDuration))

	logger.Info(out.String())

	elapsed := time.Since(start)
	logger.Infof("Running imposm took %s", elapsed)

}

func logIfFailing(err error) {
	if err != nil {
		logger.Error(err.Error())
	}

}
