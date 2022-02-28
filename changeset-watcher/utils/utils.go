package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"noxz.dev/changeset-watcher/config"
	"noxz.dev/changeset-watcher/types"
)

func ExtractSeqNumber(body *string) (int, error) {
	r, _ := regexp.Compile("sequenceNumber=(\\d+)")

	seqString := strings.Split(r.FindString(*body), "=")[1]

	seqNumber, err := strconv.Atoi(seqString)

	if err != nil {
		return 0, err
	}

	return seqNumber, nil
}

func BuildChangeSetUrl(seqNumber int) (string, error) {

	seq := fmt.Sprint("000000000", seqNumber)
	seqShorted := seq[len(seq)-9:]
	var result string
	for i, s := range seqShorted {
		if i%3 == 0 && i != 0 {
			result += "/"
		}
		result += string(s)
	}
	url := "https://planet.openstreetmap.org/replication/minute/" + fmt.Sprint(result) + ".osc.gz"
	return url, nil
}

func CreateEvent(source string, payload interface{}, subject string) (*event.Event, error) {
	natsPublishEvent := cloudevents.NewEvent()
	natsPublishEvent.SetID(uuid.New().String())
	natsPublishEvent.SetSource(source)
	natsPublishEvent.SetType(subject + "Event")
	err := natsPublishEvent.SetData(cloudevents.ApplicationJSON, payload)

	return &natsPublishEvent, err
}

func GenSubject(names ...string) string {
	var subject = config.RootEvent

	for _, name := range names {
		subject += "." + name
	}

	return subject
}

func WriteObjectToFile(object interface{}, filename string) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := jsoniter.NewEncoder(file)
	encoder.Encode(object)
}

func CalculateCentroid(points *[]types.Node) types.Location {
	var xSum float32 = 0.0
	var ySum float32 = 0.0
	var len float32 = 0

	for _, p := range *points {
		xSum += p.Lat
		ySum += p.Lon
		len++
	}

	centroid := types.Location{Lat: xSum / len, Lng: ySum / len}

	return centroid
}

// func GeneratePointsFromNodes(nodes *[]types.Node) []Point {
// 	points := make([]Point, 0)
// 	for _, n := range *nodes {
// 		points = append(points, Point{Lat: n.Lat, Lng: n.Lon})
// 	}

// 	return points
// }
