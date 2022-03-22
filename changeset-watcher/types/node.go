package types

import (
	"encoding/xml"
	"io"
	"net/http"
	"noxz.dev/changeset-watcher/config"
	"strconv"
	"strings"
)

func (node Node) HasTags(tags ...string) bool {
	for _, tag := range tags {
		_, exists := node.GetTag(tag)
		if !exists {
			return false
		}
	}
	return true

}

func (node Node) GetTag(tagString string) (value string, exits bool) {
	for _, tag := range node.Tags {
		if tagString == tag.K {
			return tag.V, true
		}
	}
	return "", false

}

func getNodesByID(nodeIDs map[int]struct{}) (nodes []Node, err error) {
	nodes = make([]Node, 0)
	var overpassAnswer OverPassAnswer
	prefixString := "[out:xml][timeout:500];node(id: "
	postfixString := "0);out meta;"

	bodyBuilder := strings.Builder{}
	bodyBuilder.WriteString(prefixString)

	for i := range nodeIDs {
		bodyBuilder.WriteString(strconv.Itoa(i) + ",")
	}

	bodyBuilder.WriteString(postfixString)
	//structure of requestBody: "[out:xml][timeout:500];node(id: 9309596758, 9519334485, ... ); out;"
	requestBody := strings.NewReader(bodyBuilder.String())
	resp, err := http.Post(config.OverpassApiURL, "x-www-form-urlencoded", requestBody)

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

func (node *Node) GetAddressString() string {
	sb := strings.Builder{}
	street, useSeparator := node.GetTag("addr:street")
	sb.WriteString(street)

	if houseNumber, exists := node.GetTag("addr:housenumber"); exists {
		sb.WriteString(" " + houseNumber)
		useSeparator = true
	}

	if postCode, exists := node.GetTag("addr:postcode"); exists {
		if useSeparator {
			sb.WriteString(", ")
		}
		sb.WriteString(postCode)
		useSeparator = true
	}

	if city, exists := node.GetTag("addr:city"); exists {
		if useSeparator {
			sb.WriteString(", ")
		}
		sb.WriteString(city)
		useSeparator = true
	}

	if country, exists := node.GetTag("addr:country"); exists {
		if useSeparator {
			sb.WriteString(", ")
		}
		sb.WriteString(country)
		useSeparator = true
	}

	return sb.String()
}
