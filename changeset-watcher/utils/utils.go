package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"noxz.dev/changeset-watcher/types"
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

func ExtractMissingNodes(action *types.Action) (nodeIDs map[int]struct{}, missingNodes int, foundNodes int) {
	missingNodes = 0
	foundNodes = 0
	nodeIDs = make(map[int]struct{})
	for _, way := range action.Ways {
		for _, ref := range way.NodeRefs {
			if action.ContainsNodeByRef(ref) {
				foundNodes++
			} else {
				missingNodes++
				nodeIDs[ref.Ref] = struct{}{}
			}
		}
	}
	return
}

func GetNodesByID(nodeIDs map[int]struct{}) (nodes []types.Node, err error) {
	nodes = make([]types.Node, 0)
	var overpassAnswer types.OverPassAnswer
	prefixString := "[out:xml][timeout:500];node(id: "
	postfixString := "0);out;"

	bodyBuilder := strings.Builder{}
	bodyBuilder.WriteString(prefixString)

	for i := range nodeIDs {
		bodyBuilder.WriteString(strconv.Itoa(i) + ",")
	}

	bodyBuilder.WriteString(postfixString)
	//structure of requestBody: "[out:xml][timeout:500];node(id: 9309596758, 9519334485, ... ); out;"
	requestBody := strings.NewReader(bodyBuilder.String())
	resp, err := http.Post("https://overpass-api.de/api/interpreter", "x-www-form-urlencoded", requestBody)

	if err != nil {
		return nodes, err
	}

	responseBody, _ := io.ReadAll(resp.Body)
	err = xml.Unmarshal(responseBody, &overpassAnswer)
	if err != nil {
		return nodes, err
	}

	return overpassAnswer.Nodes, nil
}
