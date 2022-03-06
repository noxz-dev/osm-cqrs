package types

import (
	"time"
)

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

func (action *Action) getNewestNodeVersions(newestVersionOfNode *map[int]time.Time) {
	for _, node := range action.Nodes {
		creationTime, err := node.getCreationTime()
		if err != nil {
			continue
		}
		t, exists := (*newestVersionOfNode)[node.Id]
		if exists && t.After(creationTime) {
			(*newestVersionOfNode)[node.Id] = t
		} else {
			(*newestVersionOfNode)[node.Id] = creationTime
		}
	}
}

func (action *Action) deleteOldNodeVersions(newestVersionOfNode *map[int]time.Time) {
	newestNodes := make([]Node, 0)
	for _, node := range action.Nodes {
		creationTime, err := node.getCreationTime()
		if err != nil {
			continue
		}
		newestTime, exists := (*newestVersionOfNode)[node.Id]

		if exists && creationTime.Equal(newestTime) {
			newestNodes = append(newestNodes, node)
		}
	}
	action.Nodes = newestNodes
}

func (action *Action) getNewestWayVersions(newestVersionOfWay *map[int]time.Time) {
	for _, way := range action.Ways {
		creationTime, err := way.getCreationTime()
		if err != nil {
			continue
		}
		t, exists := (*newestVersionOfWay)[way.Id]
		if exists && t.After(creationTime) {
			(*newestVersionOfWay)[way.Id] = t
		} else {
			(*newestVersionOfWay)[way.Id] = creationTime
		}
	}
}

func (action *Action) deleteOldWayVersions(newestVersionOfWay *map[int]time.Time) {
	newestWay := make([]Way, 0)
	for _, way := range action.Ways {
		creationTime, err := way.getCreationTime()
		if err != nil {
			continue
		}
		newestTime, exists := (*newestVersionOfWay)[way.Id]

		if exists && creationTime.Equal(newestTime) {
			newestWay = append(newestWay, way)
		}
	}
	action.Ways = newestWay
}

func (action *Action) getNewestRelationVersions(newestVersionOfRelation *map[int]time.Time) {
	for _, relation := range action.Relations {
		creationTime, err := relation.getCreationTime()
		if err != nil {
			continue
		}
		t, exists := (*newestVersionOfRelation)[relation.Id]
		if exists && t.After(creationTime) {
			(*newestVersionOfRelation)[relation.Id] = t
		} else {
			(*newestVersionOfRelation)[relation.Id] = creationTime
		}
	}
}

func (action *Action) deleteOldRelationVersions(newestVersionOfRelation *map[int]time.Time) {
	newestRelations := make([]Relation, 0)
	for _, relation := range action.Relations {
		creationTime, err := relation.getCreationTime()
		if err != nil {
			continue
		}
		newestTime, exists := (*newestVersionOfRelation)[relation.Id]

		if exists && creationTime.Equal(newestTime) {
			newestRelations = append(newestRelations, relation)
		}
	}
	action.Relations = newestRelations
}
