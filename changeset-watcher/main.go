package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	gzip "github.com/klauspost/pgzip"
	"github.com/nats-io/nats.go"

	"noxz.dev/changeset-watcher/types"
	"noxz.dev/changeset-watcher/utils"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)

	var oldSeq = 0

	for {

		resp, err := http.Get("https://planet.openstreetmap.org/replication/minute/state.txt")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body, _ := io.ReadAll(resp.Body)
		stringBody := string(body)

		seq, err := utils.ExtractSeqNumber(&stringBody)

		if oldSeq >= seq {
			fmt.Println("same seq .. waiting for 10 sec")
			time.Sleep(10 * time.Second)
			continue
		}
		oldSeq = seq

		fmt.Println("new seq " + fmt.Sprint(seq) + " proceeding")

		url, err := utils.BuildChangeSetUrl(seq)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("fetching " + url)

		resp, err = http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader, err := gzip.NewReader(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body, _ = io.ReadAll(reader)

		fmt.Println("parsing xml ...")
		osm := types.OsmChange{}
		err = xml.Unmarshal(body, &osm)

		// for _, cs := range osm.ChageSets {
		sendNewChangesetNotifcation(nc, &osm)
		// }

		// fmt.Printf("%+v\n", osm.ChageSets)
	}
}

func sendNewChangesetNotifcation(nc *nats.Conn, change *types.OsmChange) {
	fmt.Println("publishing new changeset...")
	changeSetBytes, _ := json.Marshal(change)

	nc.Publish("foo", changeSetBytes)
}
