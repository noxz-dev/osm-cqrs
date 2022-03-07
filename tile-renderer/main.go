package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"log"
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

	_, err = nc.Subscribe("all", func(msg *nats.Msg) {
		log.Printf("Received message")

		cloudEvent := cloudevents.NewEvent()

		_ = json.Unmarshal(msg.Data, &cloudEvent)

		xmlDataZipped := bytes.NewBuffer(cloudEvent.Data())

		log.Printf("Event data size: %d", len(xmlDataZipped.Bytes()))

		err := os.WriteFile("temp.osc.gz", xmlDataZipped.Bytes(), 0644)
		if err != nil {
			log.Fatalf("Could not write file %s", err.Error())
			return
		}

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
