package routing

import (
	"testing"

	"github.com/rknizzle/simplekv/pkg/storage"
)

func TestGetNodesForKey(t *testing.T) {
	r := RendezvousHash{
		Nodes: []storage.StorageNode{
			storage.TestStorageNode{
				Label:         "localhost:3000",
				StorageEngine: storage.InmemoryStorage{},
			},
			storage.TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.InmemoryStorage{},
			},
			storage.TestStorageNode{
				Label:         "localhost:3002",
				StorageEngine: storage.InmemoryStorage{},
			},
		},
	}

	key := "test.txt"
	numReplicas := 2
	nodesForKey := r.getNodesForKey(key, numReplicas)

	// assert that numReplicas nodes are returned
	if len(nodesForKey) != numReplicas {
		t.Fatalf("Expected getNodesForKey to return %d nodes but it returned %d", numReplicas, len(nodesForKey))
	}
}
