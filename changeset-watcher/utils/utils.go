package utils

import (
	"bytes"
	"fmt"
	gzip "github.com/klauspost/pgzip"
	"os"
	"regexp"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
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

func BuildChangeSetUrl(seqNumber int) string {

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
	return url
}

func CreateEvent(source string, payload interface{}, subject string, contentType string) (*event.Event, error) {
	natsPublishEvent := cloudevents.NewEvent()
	natsPublishEvent.SetID(uuid.New().String())
	natsPublishEvent.SetSource(source)
	natsPublishEvent.SetType(subject + "Event")
	err := natsPublishEvent.SetData(contentType, payload)

	return &natsPublishEvent, err
}

// WriteObjectToFile : Serialize Struct as JSON-Object. This can be used for debugging.
func WriteObjectToFile(object interface{}, filename string) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := jsoniter.NewEncoder(file)
	encoder.Encode(object)
}

func CalculateCentroid(points *[]types.Node) types.Location {
	var xSum float64 = 0.0
	var ySum float64 = 0.0
	var numberOfNodes float64 = 0

	for _, p := range *points {
		xSum += p.Lat
		ySum += p.Lon
		numberOfNodes++
	}

	centroid := types.Location{Lat: xSum / numberOfNodes, Lng: ySum / numberOfNodes}

	return centroid
}

func Compress(data []byte) (compressedBytes []byte, err error) {

	var buffer bytes.Buffer
	w := gzip.NewWriter(&buffer)
	_, err = w.Write(data)
	if err != nil {
		return
	}

	err = w.Close()
	if err != nil {
		return
	}
	return buffer.Bytes(), nil
}
