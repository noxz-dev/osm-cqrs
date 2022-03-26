package main

import (
	"bytes"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"github.com/robfig/cron"
	"noxz.dev/routing/osrm/config"
	"noxz.dev/routing/osrm/statistics"
	"noxz.dev/routing/osrm/update"
	"os"
	"strconv"
)

func main() {
	defer func(mapUpdateStat *statistics.Statistic) {
		err := mapUpdateStat.Close()
		if err != nil {
			config.Logger.Errorf("Error occurred while closing the map statistics: %s", err.Error())
			return
		}
	}(&config.MapUpdateStat)

	defer func(routingUpdateStat *statistics.Statistic) {
		err := routingUpdateStat.Close()
		if err != nil {
			config.Logger.Errorf("Error occurred while closing the routing statistics: %s", err.Error())
			return
		}
	}(&config.RoutingUpdateStat)

	if !config.CollectStatistics {
		err := config.MapUpdateStat.Pause()
		if err != nil {
			config.Logger.Errorf("Error occurred while pausing the map statistics: %s", err.Error())
			return
		}
		err = config.RoutingUpdateStat.Pause()
		if err != nil {
			config.Logger.Errorf("Error occurred while pausing the routing statistics: %s", err.Error())
			return
		}
	}

	url := os.Getenv("NATS_IP")

	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)

	config.Logger.Info("Connecting to NATS-Server...")

	if err != nil {
		config.Logger.Errorf("Failed to connect to NATS-Server: \n%s \n", err.Error())
		return
	}
	defer nc.Close()

	config.Logger.Info("Connection successful")

	_, err = nc.Subscribe("routing", func(msg *nats.Msg) {
		config.LogIfFailing(config.MapUpdateStat.BeginnColum())

		config.Logger.Info("Received message")

		cloudEvent := cloudevents.NewEvent()
		err = json.Unmarshal(msg.Data, &cloudEvent)

		if err != nil {
			config.Logger.Fatalf("Error while running json unmarshal: %s", err.Error())
			return
		}

		zippedXMLData := bytes.NewBuffer(cloudEvent.Data())

		config.Logger.Infof("Event data size: %s", strconv.Itoa(zippedXMLData.Len()))

		config.LogIfFailing(config.MapUpdateStat.SetValue(config.SizeOfIncomingEvent, strconv.Itoa(zippedXMLData.Len())))

		config.LogIfFailing(config.MapUpdateStat.StartTimer(config.WriteDuration))

		err := os.WriteFile(config.DataDir+"change.osc.gz", zippedXMLData.Bytes(), 0644)

		config.LogIfFailing(config.MapUpdateStat.StopTimerAndSetDuration(config.WriteDuration))
		if err != nil {
			config.Logger.Errorf("Could not write file %s", err.Error())
			return
		}

		update.RunLocalMapUpdate()

		config.LogIfFailing(config.MapUpdateStat.EndColum())
	})

	if err != nil {
		config.Logger.Errorf("Error occurred while subscribing: \n%s \n\n", err.Error())
		return
	}

	StartCronForRoutingServerUpdate()

	select {}
}

// StartCronForRoutingServerUpdate starts a cron job that updates the routing servers every 5 minutes
func StartCronForRoutingServerUpdate() {
	c := cron.New()
	err := c.AddFunc("@every 5m", update.RunRoutingServerUpdate)
	if err != nil {
		config.Logger.Infof("Could not start cron job %s", err.Error())
		return
	}
	c.Start()

	config.Logger.Info("Cron job created")
}
