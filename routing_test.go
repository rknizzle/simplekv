package main

import (
	"fmt"
	"testing"
)

func TestGetNodesForKey(t *testing.T) {
	rs := routingServer{
		nodes:       []string{"localhost:3000", "localhost:3001", "localhost:3002"},
		numReplicas: 2,
	}

	key := "test.txt"
	nodesForKey := rs.getNodesForKey(key)

	fmt.Println(nodesForKey)
	// assert that numReplicas nodes are returned
	if len(nodesForKey) != rs.numReplicas {
		t.Fatalf("Expected getNodesForKey to return %d nodes but it returned %d", rs.numReplicas, len(nodesForKey))
	}
}
