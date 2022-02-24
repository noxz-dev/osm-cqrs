package nodes_reloading

import (
	"encoding/xml"
	"io"
	"net/http"
	"noxz.dev/changeset-watcher/types"
	"strconv"
	"strings"
)

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
