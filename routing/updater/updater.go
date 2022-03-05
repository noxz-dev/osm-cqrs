package updater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
)


func RunUpdate() {
	url := os.Getenv("NATS_IP")

	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)

	log.Printf("Connecting to NATS-Server...")

	if err != nil {
		log.Fatalf("Failed to connect to NATS-Server: \n%s \n", err.Error())
		return
	}

	defer nc.Close()


	log.Printf("Connection successful")

	_, err = nc.Subscribe("routing", func(msg *nats.Msg) {
		log.Printf("Received message")

		cloudEvent := cloudevents.NewEvent()

		err = json.Unmarshal(msg.Data, &cloudEvent)

		if err != nil {
			return
		}

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

	err := cmd.Run()

	if err != nil {
		log.Fatalf(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	log.Printf(out.String())

	elapsed := time.Since(start)
	log.Printf("Running imposm took %s", elapsed)

}