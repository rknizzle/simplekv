package main

import (
	"fmt"
	"testing"
)

func TestGetNodesForKey(t *testing.T) {
	r := rendezvousHash{
		nodes: []string{"localhost:3000", "localhost:3001", "localhost:3002"},
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
