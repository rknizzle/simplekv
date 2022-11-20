package main

import (
	"fmt"
	"testing"
)

func TestGetNodesForKey(t *testing.T) {
	r := rendezvousHash{
		nodes: []storageNode{
			storageNode{
				label:         "localhost:3000",
				storageEngine: inmemoryStorage{},
			},
			storageNode{
				label:         "localhost:3001",
				storageEngine: inmemoryStorage{},
			},
			storageNode{
				label:         "localhost:3002",
				storageEngine: inmemoryStorage{},
			},
		},
	}

	key := "test.txt"
	numReplicas := 2
	nodesForKey := r.getNodesForKey(key, numReplicas)

	fmt.Println(nodesForKey)
	// assert that numReplicas nodes are returned
	if len(nodesForKey) != numReplicas {
		t.Fatalf("Expected getNodesForKey to return %d nodes but it returned %d", numReplicas, len(nodesForKey))
	}
}
