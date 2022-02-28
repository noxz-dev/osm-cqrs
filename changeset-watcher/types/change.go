package types

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
	reloadedNodes, err := getNodesByID(nodeIDs)
	if err != nil {
		return err
	}
	normalized.Reloaded.Nodes = reloadedNodes
	return nil
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
