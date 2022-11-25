package routing

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/rknizzle/simplekv/pkg/storage"
)

func TestSaveValueToKey(t *testing.T) {
	rh := RendezvousHash{
		Nodes: []StorageNode{
			TestStorageNode{
				Label:         "localhost:3000",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3002",
				StorageEngine: storage.NewInmemoryStorage(),
			},
		},
	}

	rs := NewRoutingServer(2, rh)

	key := "hello"
	value := "world"
	valueReader := strings.NewReader(value)

	err := rs.saveValueToKey(key, valueReader)
	if err != nil {
		t.Fatalf("Failed to save the value for key: %s with message: %s", key, err.Error())
	}

	// verify that the value was written to the expected number of nodes
	nodes := rs.hash.getAllNodes()

	// TODO: clean this up to verify that the value was correctly written numReplicas times but with
	// less noise
	for i, node := range nodes {
		// log the value for the key on each node (even if it didnt get written to it)
		valReader, err := node.Get(key)
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
}

func TestSuccessfulGet(t *testing.T) {
	// NOTE: the key 'hello' maps to nodes 0 and 2 when there are 3 nodes
	key := "hello"
	value := "world"

	rh := RendezvousHash{
		Nodes: []StorageNode{
			TestStorageNode{
				Label: "localhost:3000",
				StorageEngine: storage.InmemoryStorage{
					StorageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
			TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label: "localhost:3002",
				StorageEngine: storage.InmemoryStorage{
					StorageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
		},
	}

	rs := NewRoutingServer(2, rh)

	valueReader, err := rs.get(key)
	if err != nil {
		t.Fatalf("Failed to get the value with message: %s", err.Error())
	}

	data, err := io.ReadAll(valueReader)
	if err != nil {
		t.Fatalf("Failed to read the data from the io.Reader with message: %s", err.Error())
	}

	if string(data) != value {
		t.Fatalf("Expected: %s got: %s", value, string(data))
	}
}

func TestWithMissingKey(t *testing.T) {
	key := "doesntExistOnAnyNode"

	rh := RendezvousHash{
		Nodes: []StorageNode{
			TestStorageNode{
				Label:         "localhost:3000",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3002",
				StorageEngine: storage.NewInmemoryStorage(),
			},
		},
	}

	rs := NewRoutingServer(2, rh)

	valueReader, err := rs.get(key)
	if err == nil {
		t.Fatalf("Expected to get an error trying to find the key on any node")
	}

	if valueReader != nil {
		t.Fatalf("Expected not to get a valueReader back from any node")
	}

	if !strings.Contains(err.Error(), "Failed to get key") {
		t.Fatalf(fmt.Sprintf("Expected to get an error about a missing key but instead got the error: %s", err.Error()))
	}
}
