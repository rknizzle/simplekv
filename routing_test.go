package main

import (
	"strings"
	"testing"
)

func TestSaveValueToKey(t *testing.T) {
	rh := rendezvousHash{
		nodes: []storageNode{
			storageNode{
				label:         "localhost:3000",
				storageEngine: newInmemoryStorage(),
			},
			storageNode{
				label:         "localhost:3001",
				storageEngine: newInmemoryStorage(),
			},
			storageNode{
				label:         "localhost:3002",
				storageEngine: newInmemoryStorage(),
			},
		},
	}

	rs := newRoutingServer(2, rh)

	key := "hello"
	value := "world"
	valueReader := strings.NewReader(value)

	err := rs.saveValueToKey(key, valueReader)
	if err != nil {
		t.Fatalf("Failed to save the value for key: %s with message: %s", key, err.Error())
	}

	// TODO: iterate through the nodes and check which/how many got the the value written to them
}
