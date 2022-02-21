package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	if err != nil {
		fmt.Printf("Failed to connect to NATS-Server: %s \n", err.Error())
	}

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	for {
	}
}
