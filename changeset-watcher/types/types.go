package types

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Osm struct {
	Version   string      `xml:"version"`
	ChageSets []ChangeSet `xml:"changeset"`
}

type ChangeSet struct {
	Id         int     `xml:"id,attr"`
	CreatedAt  string  `xml:"created_at,attr"`
	NumChanges int     `xml:"num_changes,attr"`
	MinLat     float32 `xml:"min_lat,attr"`
	MaxLat     float32 `xml:"max_lat,attr"`
	MinLong    float32 `xml:"min_lon,attr"`
	MaxLong    float32 `xml:"max_lon,attr"`

	Tags []Tag `xml:"tag"`
}

type Tag struct {
	K string `xml:"k,attr"`
	V string `xml:"v,attr"`
}

type OsmChange struct {
	Modify []Action `xml:"modify"`
	Create []Action `xml:"create"`
	Delete []Action `xml:"delete"`
}
type OsmChangeNormalized struct {
	Modify   Action `json:"modify"`
	Create   Action `json:"create"`
	Delete   Action `json:"delete"`
	Reloaded Action `json:"reloaded"`
}

type Action struct {
	Nodes     []Node     `xml:"node"`
	Ways      []Way      `xml:"way"`
	Relations []Relation `xml:"relation"`
}

type Node struct {
	Id        int     `xml:"id,attr"`
	Version   int     `xml:"version,attr"`
	Timestamp string  `xml:"timestamp,attr"`
	Lat       float32 `xml:"lat,attr"`
	Lon       float32 `xml:"lon,attr"`
	Tags      []Tag   `xml:"tag"`
}

type Way struct {
	Id        int       `xml:"id,attr"`
	Version   int       `xml:"version,attr"`
	Timestamp string    `xml:"timestamp,attr"`
	NodeRefs  []NodeRef `xml:"nd"`
	Tags      []Tag     `xml:"tag"`
}

type NodeRef struct {
	Ref int `xml:"ref,attr"`
}

type Relation struct {
	Id        int      `xml:"id,attr"`
	Version   int      `xml:"version,attr"`
	Timestamp string   `xml:"timestamp,attr"`
	Member    []Member `xml:"member"`
	Tags      []Tag    `xml:"tag"`
}

type Member struct {
	Type string `xml:"type,attr"`
	Ref  int    `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}

type OverPassAnswer struct {
	Nodes []Node `json:"elements" xml:"node"`
}

type SearchPoint struct {
	Name     string
	Location Location
	Id       string
	Tags     []Tag
}

type Location struct {
	Lat float32
	Lng float32
}

type SearchPayload struct {
	Modify []SearchPoint
	Create []SearchPoint
	Delete []SearchPoint
}

const (
	MODIFY_EVENT = "MODIFY"
	DELETE_EVENT = "DELETE"
	CREATE_EVENT = "CREATE"
)

func (action Action) ContainsNodeByRef(ref NodeRef) bool {
	for _, node := range action.Nodes {
		if node.Id == ref.Ref {
			return true
		}
	}
	return false
}

func (osmChange OsmChange) Normalize() OsmChangeNormalized {
	return OsmChangeNormalized{
		Modify: normalizeActionObject(osmChange.Modify),
		Delete: normalizeActionObject(osmChange.Delete),
		Create: normalizeActionObject(osmChange.Create),
	}

}

func normalizeActionObject(actions []Action) (normalizedAction Action) {
	ways := make([]Way, 0)
	nodes := make([]Node, 0)
	relations := make([]Relation, 0)

	for _, action := range actions {
		ways = append(ways, action.Ways...)
		nodes = append(nodes, action.Nodes...)
		relations = append(relations, action.Relations...)
	}
	return Action{
		Ways:      ways,
		Nodes:     nodes,
		Relations: relations,
	}
}

func (normalized OsmChangeNormalized) ExtractMissingNodes() (nodeIDs map[int]struct{}, missingNodes int, foundNodes int) {
	missingNodes = 0
	foundNodes = 0
	nodeIDs = make(map[int]struct{})
	normalized.Modify.extractMissingNodes(&nodeIDs, &missingNodes, &foundNodes)
	//osmChangeNormalized.Delete.extractMissingNodes(&nodeIDs, &missingNodes, &foundNodes)
	normalized.Create.extractMissingNodes(&nodeIDs, &missingNodes, &foundNodes)
	return
}

func (normalized *OsmChangeNormalized) Reload() (err error) {

	nodeIDs, _, _ := normalized.ExtractMissingNodes()
	reloadedNodes, err := GetNodesByID(nodeIDs)
	if err != nil {
		return err
	}
	normalized.Reloaded.Nodes = reloadedNodes
	return nil
}

func (action Action) extractMissingNodes(nodeIDs *map[int]struct{}, missingNodes *int, foundNodes *int) {
	for _, way := range action.Ways {
		for _, ref := range way.NodeRefs {
			if action.ContainsNodeByRef(ref) {
				*foundNodes++
			} else {
				*missingNodes++
				(*nodeIDs)[ref.Ref] = struct{}{}
			}
		}
	}
}

func (way Way) HasTags(tags ...string) bool {
	for _, tag := range tags {
		_, err := way.GetTag(tag)
		if err != nil {
			return false
		}
	}
	return true

}

func (way Way) GetTag(tagString string) (value string, err error) {
	for _, tag := range way.Tags {
		if tagString == tag.K {
			return tag.V, nil
		}
	}
	return "", errors.New("Tag " + tagString + " not found")

}

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

func GetNodesByID(nodeIDs map[int]struct{}) (nodes []Node, err error) {
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

func (action *Action) FilterWays(filters ...WayFilter) {
	filteredWays := make([]Way, 0)
	for _, way := range action.Ways {
		for _, filter := range filters {
			if way.HasTags(filter.TagKeys...) {
				filteredWays = append(filteredWays, way)
			}
		}
		action.Ways = filteredWays
	}

}

func (action Action) UsedNodes(nodeIDs *map[int]struct{}) {
	for _, way := range action.Ways {
		for _, nodeRef := range way.NodeRefs {
			(*nodeIDs)[nodeRef.Ref] = struct{}{}
		}
	}
}

func (action Action) UsedNodesByFilter(nodeIDs *map[int]struct{}, filters ...NodeFilter) {
	for _, node := range action.Nodes {
		for _, filter := range filters {
			if node.HasTags(filter.TagKeys...) {
				(*nodeIDs)[node.Id] = struct{}{}
			}
		}
	}

}

func (action *Action) DeleteAllNodesExcept(usedNodes map[int]struct{}) {
	filteredNodes := make([]Node, 0)
	for _, node := range action.Nodes {
		_, exists := usedNodes[node.Id]
		if exists {
			filteredNodes = append(filteredNodes, node)
		}
	}
	action.Nodes = filteredNodes
}

func (normalized OsmChangeNormalized) Filter(nodeFilters []NodeFilter, wayFilters []WayFilter) OsmChangeNormalized {
	usedNodes := make(map[int]struct{}, 0)
	normalized.Create.FilterWays(wayFilters...)
	normalized.Modify.FilterWays(wayFilters...)
	normalized.Delete.FilterWays(wayFilters...)

	normalized.Create.UsedNodes(&usedNodes)
	normalized.Delete.UsedNodes(&usedNodes)
	normalized.Modify.UsedNodes(&usedNodes)

	normalized.Modify.UsedNodesByFilter(&usedNodes, nodeFilters...)
	normalized.Delete.UsedNodesByFilter(&usedNodes, nodeFilters...)
	normalized.Reloaded.UsedNodesByFilter(&usedNodes, nodeFilters...)
	normalized.Create.UsedNodesByFilter(&usedNodes, nodeFilters...)

	normalized.Create.DeleteAllNodesExcept(usedNodes)
	normalized.Delete.DeleteAllNodesExcept(usedNodes)
	normalized.Modify.DeleteAllNodesExcept(usedNodes)
	normalized.Reloaded.DeleteAllNodesExcept(usedNodes)

	return normalized

}
