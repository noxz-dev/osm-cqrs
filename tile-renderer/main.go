package main

import (
	"encoding/json"
	"encoding/xml"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"noxz.dev/tile-renderer/types"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	if err != nil {
		log.Fatalf("Failed to connect to NATS-Server: \n%s \n\n", err.Error())
		return
	}

	log.Printf("Connection succesfull")

	_, err = nc.Subscribe("all", func(msg *nats.Msg) {
		log.Printf("Received message")

		cloudEvent := cloudevents.NewEvent()

		_ = json.Unmarshal(msg.Data, &cloudEvent)

		eventData := types.OsmChangeNormalized{}

		_ = json.Unmarshal(cloudEvent.Data(), &eventData)

		if err != nil {
			log.Fatalf(err.Error())
		}

		createAction := types.Action{
			Nodes:     append(eventData.Create.Nodes, eventData.Reloaded.Nodes...),
			Ways:      append(eventData.Create.Ways, eventData.Reloaded.Ways...),
			Relations: append(eventData.Create.Relations, eventData.Reloaded.Relations...),
		}

		xmlContent := types.OsmChange{
			Create: createAction,
			Delete: eventData.Delete,
			Modify: eventData.Modify,
		}

		xmlData, err := xml.MarshalIndent(xmlContent, " ", "    ")
		xmlData = []byte(xml.Header + string(xmlData))

		if err != nil {
			log.Fatalf(err.Error())
		}

		_ = ioutil.WriteFile("temp.xml", xmlData, 0644)

	})
	if err != nil {
		log.Fatalf("Error occured while subscribing: \n%s \n\n", err.Error())
		return
	}

	select {}
}
