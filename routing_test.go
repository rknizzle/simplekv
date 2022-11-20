package main

import (
	"bytes"
	"fmt"
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

	// verify that the value was written to the expected number of nodes
	nodes := rs.hash.getAllNodes()

	for i, node := range nodes {
		// log the value for the key on each node (even if it didnt get written to it)
		valReader, err := node.get(key)
		if err != nil {
			if strings.Contains(err.Error(), "doesnt exist") {
				fmt.Printf("Node %d: this node doesnt have the key & value saved\n", i)
			} else {
				fmt.Printf("Node %d error trying to get the value with message: %s\n", i, err.Error())
			}
			continue
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(valReader)
		fmt.Printf("Node %d: %s\n", i, buf.String())
	}

	// TODO ISSUE: saveValueToKey correctly tried to write the key & value to 2 nodes but the value
	// was only actually written to one of the nodes. The issue is because theres only one read stream
	// so I need to split it into a read stream for each replica

}
