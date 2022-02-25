package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	if err != nil {
		fmt.Printf("Failed to connect to NATS-Server: \n%s \n", err.Error())
		return
	}

	nc.Subscribe("change.modify", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	for {
	}
}
