package types

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

func (normalized OsmChangeNormalized) ExtractMissingNodes() (nodeIDs map[int]struct{}, missingNodes int, foundNodes int) {
	missingNodes = 0
	foundNodes = 0
	nodeIDs = make(map[int]struct{})
	normalized.Modify.extractMissingNodes(&nodeIDs, &missingNodes, &foundNodes)
	//osmChangeNormalized.Delete.extractMissingNodes(&nodeIDs, &missingNodes, &foundNodes)
	normalized.Create.extractMissingNodes(&nodeIDs, &missingNodes, &foundNodes)
	return
}

func (normalized *OsmChangeNormalized) Reload() (reloaded int, err error) {

	nodeIDs, _, _ := normalized.ExtractMissingNodes()
	reloadedNodes, err := getNodesByID(nodeIDs)
	if err != nil {
		return 0, err
	}
	normalized.Reloaded.Nodes = reloadedNodes
	return len(reloadedNodes), nil
}

// Filter returns an filtered OsmChangeNormalized object.
func (normalized OsmChangeNormalized) Filter(nodeFilters []NodeFilter, wayFilters []WayFilter) OsmChangeNormalized {
	if wayFilters == nil && nodeFilters == nil {
		return normalized
	}
	usedNodes := make(map[int]struct{}, 0)
	if wayFilters != nil {
		normalized.Create.FilterWays(wayFilters...)
		normalized.Modify.FilterWays(wayFilters...)
		normalized.Delete.FilterWays(wayFilters...)

		normalized.Create.UsedNodes(&usedNodes)
		normalized.Delete.UsedNodes(&usedNodes)
		normalized.Modify.UsedNodes(&usedNodes)
	}
	if nodeFilters != nil {
		normalized.Modify.UsedNodesByFilter(&usedNodes, nodeFilters...)
		normalized.Delete.UsedNodesByFilter(&usedNodes, nodeFilters...)
		normalized.Reloaded.UsedNodesByFilter(&usedNodes, nodeFilters...)
		normalized.Create.UsedNodesByFilter(&usedNodes, nodeFilters...)
	}
	normalized.Create.DeleteAllNodesExcept(usedNodes)
	normalized.Delete.DeleteAllNodesExcept(usedNodes)
	normalized.Modify.DeleteAllNodesExcept(usedNodes)
	normalized.Reloaded.DeleteAllNodesExcept(usedNodes)

	return normalized

}

// Normalize combines different actions of the same type to one action. Example:
//
// - different Modify sections will be combined to one single Modify section
// - different Delete sections will be combined to one single Delete section
// - different Create sections will be combined to one single Create section
func (osmChange OsmChange) Normalize() OsmChangeNormalized {
	return OsmChangeNormalized{
		Modify: normalizeActionObject(osmChange.Modify),
		Delete: normalizeActionObject(osmChange.Delete),
		Create: normalizeActionObject(osmChange.Create),
	}

}

func (normalized *OsmChangeNormalized) Size() int {
	return normalized.Delete.Size() + normalized.Modify.Size() + normalized.Create.Size() + normalized.Reloaded.Size()
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

func (normalized *OsmChangeNormalized) RemoveDuplicateNodes() {
	newestNodeVersion := make(map[int]time.Time, 0)
	normalized.Delete.getNewestNodeVersions(&newestNodeVersion)
	normalized.Modify.getNewestNodeVersions(&newestNodeVersion)
	normalized.Create.getNewestNodeVersions(&newestNodeVersion)

	normalized.Modify.deleteOldNodeVersions(&newestNodeVersion)
	normalized.Delete.deleteOldNodeVersions(&newestNodeVersion)
	normalized.Create.deleteOldNodeVersions(&newestNodeVersion)
}

func (normalized *OsmChangeNormalized) RemoveDuplicateWays() {
	newestWayVersion := make(map[int]time.Time, 0)
	normalized.Delete.getNewestWayVersions(&newestWayVersion)
	normalized.Modify.getNewestWayVersions(&newestWayVersion)
	normalized.Create.getNewestWayVersions(&newestWayVersion)

	normalized.Modify.deleteOldWayVersions(&newestWayVersion)
	normalized.Delete.deleteOldWayVersions(&newestWayVersion)
	normalized.Create.deleteOldWayVersions(&newestWayVersion)
}

func (normalized *OsmChangeNormalized) RemoveDuplicateRelations() {
	newestRelationVersion := make(map[int]time.Time, 0)
	normalized.Delete.getNewestRelationVersions(&newestRelationVersion)
	normalized.Modify.getNewestRelationVersions(&newestRelationVersion)
	normalized.Create.getNewestRelationVersions(&newestRelationVersion)

	normalized.Modify.deleteOldRelationVersions(&newestRelationVersion)
	normalized.Delete.deleteOldRelationVersions(&newestRelationVersion)
	normalized.Create.deleteOldRelationVersions(&newestRelationVersion)
}

func (normalized *OsmChangeNormalized) RemoveAllDuplicates() {
	normalized.RemoveDuplicateNodes()
	normalized.RemoveDuplicateWays()
	normalized.RemoveDuplicateRelations()
}

func (normalized *OsmChangeNormalized) ToXML() ([]byte, error) {
	createAction := Action{
		Nodes:     append(normalized.Create.Nodes, normalized.Reloaded.Nodes...),
		Ways:      append(normalized.Create.Ways, normalized.Reloaded.Ways...),
		Relations: append(normalized.Create.Relations, normalized.Reloaded.Relations...),
	}
	xmlContent := OsmChangeNormalizedXML{
		Create: createAction,
		Delete: normalized.Delete,
		Modify: normalized.Modify,
	}
	xmlContent.Version = "0.6"
	xmlData, err := xml.MarshalIndent(xmlContent, " ", "    ")
	if err != nil {
		return xmlData, err
	}
	return []byte(xml.Header + string(xmlData)), err
}

func (normalized *OsmChangeNormalized) ToJSON() ([]byte, error) {
	return json.Marshal(normalized)
}
