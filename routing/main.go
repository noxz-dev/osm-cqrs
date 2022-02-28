package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"io"
	"log"
	"noxz.dev/routing/types"
	"os"
	"os/exec"
	"time"
)

func main() {
	url := os.Getenv("NATS_IP")

	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	defer nc.Close()

	log.Printf("Connecting to NATS-Server...")

	if err != nil {
		log.Fatalf("Failed to connect to NATS-Server: \n%s \n", err.Error())
		return
	}

	log.Printf("Connection succesfull")

	_, err = nc.Subscribe("routing", func(msg *nats.Msg) {
		log.Printf("Received message")

		start := time.Now()

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

		file, err := os.Create("temp.osc.gz")
		if err != nil {
			log.Fatalf(err.Error())
		}

		// For running test update.xml and create.xml in testdata folder
		// data, err := ioutil.ReadFile("update.xml")
		// if err != nil {
		//	 log.Fatal(err)
		// }

		err = GzipWrite(file, xmlData)

		if err != nil {
			log.Fatalf(err.Error())
		}

		elapsed := time.Since(start)

		log.Printf("Writing XML took %s", elapsed)

		RunImposmUpdate()

	})

	if err != nil {
		log.Fatalf("Error occured while subscribing: \n%s \n\n", err.Error())
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
	xmlFile := "./temp.osc.gz"

	cmd := exec.Command(imposm, imposmCommand, configCommand, configLocation, xmlFile)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	err := cmd.Run()

	if err != nil {
		log.Fatalf(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	log.Printf(out.String())

	elapsed := time.Since(start)
	log.Printf("Running imposm took %s", elapsed)

}

// GzipWrite gzips the given data and writes it to the disk
func GzipWrite(w io.Writer, data []byte) error {
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer func(gw *gzip.Writer) {
		err := gw.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(gw)
	_, err = gw.Write(data)
	if err != nil {
		return err
	}
	return err
}
