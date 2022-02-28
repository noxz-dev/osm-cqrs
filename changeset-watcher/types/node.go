package types

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (node Node) HasTags(tags ...string) bool {
	for _, tag := range tags {
		_, err := node.GetTag(tag)
		if err != nil {
			return false
		}
	}
	return true

}

func (node Node) GetTag(tagString string) (value string, err error) {
	for _, tag := range node.Tags {
		if tagString == tag.K {
			return tag.V, nil
		}
	}
	return "", errors.New("Tag " + tagString + " not found")

}

func getNodesByID(nodeIDs map[int]struct{}) (nodes []Node, err error) {
	nodes = make([]Node, 0)
	var overpassAnswer OverPassAnswer
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
