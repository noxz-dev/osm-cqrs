package utils

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
	"regexp"
	"strconv"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"noxz.dev/changeset-watcher/config"
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

func CreateEvent(source string, eventType string, payload interface{}) *event.Event {
	event := cloudevents.NewEvent()
	event.SetID(uuid.New().String())
	event.SetSource(source)
	event.SetType(eventType)
	event.SetData(cloudevents.ApplicationJSON, payload)

	return &event
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
