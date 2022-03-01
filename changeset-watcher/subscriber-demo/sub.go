package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	if err != nil {
		fmt.Printf("Failed to connect to NATS-Server: \n%s \n", err.Error())
		return
	}

	nc.Subscribe("routing", func(m *nats.Msg) {

		cloudEvent := cloudevents.NewEvent()

		_ = json.Unmarshal(m.Data, &cloudEvent)

		buff := bytes.NewBuffer(cloudEvent.Data())

		reader, err := gzip.NewReader(buff)

		if err != nil {
			fmt.Println(err.Error())
		}
		msg, _ := io.ReadAll(reader)
		fmt.Printf("Received a message: %s\n", string(msg))
	})

	for {
	}
}
