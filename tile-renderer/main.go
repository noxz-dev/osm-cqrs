package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"noxz.dev/tile-renderer/types"
	"os/exec"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	if err != nil {
		log.Fatalf("Failed to connect to NATS-Server: \n%s \n\n", err.Error())
		return
	}

	_, err = nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Printf("Received message")

		xmlData := types.Action{}

		_ = json.Unmarshal(msg.Data, &xmlData)

		file, _ := xml.MarshalIndent(xmlData, "", " ")

		_ = ioutil.WriteFile("temp.xml", file, 0644)

		cmd := exec.Command("osm2pgsql", "--append", "-r xml", "-s", "-C 1000", "-G", "--hstore",
			"--number-processes 24",
			"--style openstreetmap-carto.style",
			"--tag-transform-script openstreetmap-carto.lua",
			"-d gis",
			"-H localhost",
			"-U renderer",
			"-W",
			"--password renderer",
			"temp.xml")

		err := cmd.Run()

		if err != nil {
			log.Fatalln(err)
		}

	})
	if err != nil {
		log.Fatalf("Error occured while subscribing: \n%s \n\n", err.Error())
		return
	}

	select {}
}
