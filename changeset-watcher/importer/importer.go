package main

import (
	"context"
	"fmt"
	"os"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
)

func main() {
	f, err := os.Open("../data/niedersachsen.pbf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()

	var nc, wc, rc uint64

	for scanner.Scan() {
		switch o := scanner.Object().(type) {
		case *osm.Node:
			nc++

		case *osm.Way:
			wc++

		case *osm.Relation:
			rc++
		}
		
	}

	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)

	scanErr := scanner.Err()
	if scanErr != nil {
		panic(scanErr)
	}
}
