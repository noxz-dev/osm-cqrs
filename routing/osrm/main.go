package main

import (
	"bytes"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"github.com/withmandala/go-log"
	"noxz.dev/routing/osrm/config"
	"noxz.dev/routing/osrm/update"
	"os"
	"strconv"
)

var logger = log.New(os.Stderr)

func main() {
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

		logger.Info("Received message")

		cloudEvent := cloudevents.NewEvent()

		err = json.Unmarshal(msg.Data, &cloudEvent)

		if err != nil {
			return
		}

		zippedXMLData := bytes.NewBuffer(cloudEvent.Data())

		logger.Infof("Event data size: %s", strconv.Itoa(zippedXMLData.Len()))

		err := os.WriteFile(config.DataDir+"change.osc.gz", zippedXMLData.Bytes(), 0644)

		if err != nil {
			logger.Errorf("Could not write file %s", err.Error())
			return
		}

		update.RunLocalMapUpdate()
	})

	if err != nil {
		logger.Errorf("Error occured while subscribing: \n%s \n\n", err.Error())
		return
	}

	//block forever
	select {}
}
