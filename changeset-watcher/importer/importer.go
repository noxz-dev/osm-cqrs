package importer

import (
	"context"
	"os"
	"time"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"github.com/withmandala/go-log"
	"noxz.dev/changeset-watcher/types"
)

var logger = log.New(os.Stderr)

func Import(filePath string) (*[]types.OsmChangeNormalized, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()

	var nc, wc, rc uint64

	nodesMap := make(map[int]types.Node, 0)
	ways := make([]types.Way, 0)
	relations := make([]types.Relation, 0)

	for scanner.Scan() {
		o := scanner.Object()
		switch o.(type) {
		case *osm.Node:
			nc++
			tmp := o.(*osm.Node)

			tags := make([]types.Tag, 0)

			for _, t := range tmp.Tags {
				tags = append(tags, types.Tag{
					K: t.Key,
					V: t.Value,
				})
			}

			nodesMap[int(tmp.ID)] = types.Node{
				Id:        int(tmp.ID),
				Timestamp: tmp.Timestamp.String(),
				Lat:       tmp.Lat,
				Lon:       tmp.Lon,
				Tags:      tags,
			}
		case *osm.Way:
			wc++
			tmp := o.(*osm.Way)

			refs := make([]types.NodeRef, 0)

			for _, n := range tmp.Nodes {
				refs = append(refs, types.NodeRef{
					Ref: int(n.ID),
				})
			}
			tags := make([]types.Tag, 0)

			for _, t := range tmp.Tags {
				tags = append(tags, types.Tag{
					K: t.Key,
					V: t.Value,
				})
			}

			ways = append(ways, types.Way{
				Id:        int(tmp.ID),
				Version:   tmp.Version,
				Timestamp: tmp.Timestamp.String(),
				NodeRefs:  refs,
				Tags:      tags,
			})
		case *osm.Relation:
			rc++
			tmp := o.(*osm.Relation)

			members := make([]types.Member, 0)

			for _, m := range tmp.Members {
				members = append(members, types.Member{
					Type: string(m.Type),
					Role: m.Role,
					Ref:  int(m.Ref),
				})
			}

			tags := make([]types.Tag, 0)

			for _, t := range tmp.Tags {
				tags = append(tags, types.Tag{
					K: t.Key,
					V: t.Value,
				})
			}

			relations = append(relations, types.Relation{
				Id:        int(tmp.ID),
				Version:   tmp.Version,
				Timestamp: tmp.Timestamp.String(),
				Member:    members,
				Tags:      tags,
			})
		}
	}

	logger.Infof("found %d Nodes, %d Ways, %d Relations", nc, wc, rc)

	now := time.Now()

	changesets := generateChangeSets(&ways, &nodesMap, &relations, 5000)

	logger.Info("generated", len(*changesets), "changesets, generation took:", time.Since(now), "ms")
	scanErr := scanner.Err()
	if scanErr != nil {
		return nil, scanErr
	}

	return changesets, nil
}

func generateChangeSets(ways *[]types.Way, nodes *map[int]types.Node, relations *[]types.Relation, chunkSize int) *[]types.OsmChangeNormalized {
	changeSets := make([]types.OsmChangeNormalized, 0)
	changeset := types.OsmChangeNormalized{}
	wayCount := 0

	remainingNodes := make([]types.Node, 0, len(*nodes))
	logger.Info("started changeset generation ...")
	for _, way := range *ways {
		if wayCount == chunkSize {
			changeSets = append(changeSets, changeset)
			changeset = types.OsmChangeNormalized{}
			wayCount = 0
		}

		foundNodes := getNodesToWay(&way, nodes)
		
		changeset.Create.Ways = append(changeset.Create.Ways, way)
		changeset.Create.Nodes = append(changeset.Create.Nodes, *foundNodes...)
		wayCount++
	}

	//add the changeset to the list, if the waycount is smaller than the chunkSize and all ways are allready processed
	if wayCount < chunkSize {
		changeSets = append(changeSets, changeset)
	}

	allChangesetNodes := make([]types.Node, 0)

	for _, cs := range changeSets {
		allChangesetNodes = append(allChangesetNodes, cs.Create.Nodes...)
	}

	remNodes := getAllRemainingNodes(&allChangesetNodes, nodes)

	remainingNodes = append(remainingNodes, *remNodes...)
	logger.Debug("Map length", len(*nodes))
	logger.Debug("Remaining Nodes", len(remainingNodes))
	logger.Debug("all found Nodes", len(allChangesetNodes))

	for i := 0; i < len(remainingNodes); i += chunkSize {
		changeset = types.OsmChangeNormalized{}
		end := i + chunkSize
		if end > len(remainingNodes) {
			end = len(remainingNodes)
		}
		changeset.Create.Nodes = append(changeset.Create.Nodes, remainingNodes[i:end]...)

		changeSets = append(changeSets, changeset)
	}

	for i := 0; i < len(*relations); i += chunkSize {
		changeset = types.OsmChangeNormalized{}
		end := i + chunkSize
		if end > len(*relations) {
			end = len(*relations)
		}
		changeset.Create.Relations = append(changeset.Create.Relations, (*relations)[i:end]...)
		changeSets = append(changeSets, changeset)
	}


	return &changeSets
}

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func getAllRemainingNodes(foundNodes *[]types.Node, nodes *map[int]types.Node) *[]types.Node {
	remainingNodes := make([]types.Node, 0, len(*nodes))

	for _, value := range *nodes {
		found := false
		for _, node := range *foundNodes {
			if node.Id == value.Id {
				found = true
			}
		}
		if !found {
			remainingNodes = append(remainingNodes, value)
		}
	}

	return &remainingNodes
}

func getNodesToWay(way *types.Way, nodes *map[int]types.Node) *[]types.Node {
	foundNodes := make([]types.Node, 0)
	for _, nodeRef := range way.NodeRefs {
			
		if node, ok := (*nodes)[nodeRef.Ref]; ok {
			foundNodes = append(foundNodes, node)
		}
	}

	return &foundNodes
}
