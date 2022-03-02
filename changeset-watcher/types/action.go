package types

func (action Action) ContainsNodeByRef(ref NodeRef) bool {
	for _, node := range action.Nodes {
		if node.Id == ref.Ref {
			return true
		}
	}
	return false
}

func (action *Action) Size() int {
	return len(action.Ways) + len(action.Relations) + len(action.Nodes)
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
