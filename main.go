package main

import (
	"fmt"
	"net/http"
)

func main() {
	// setup the routing server based on the config
	port := 8080

	key := "hello"
	value := "world"

	// preload a key in the nodes just for testing
	rh := rendezvousHash{
		nodes: []storageNode{
			testStorageNode{
				label: "localhost:3000",
				storageEngine: inmemoryStorage{
					storageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
			testStorageNode{
				label:         "localhost:3001",
				storageEngine: newInmemoryStorage(),
			},
			testStorageNode{
				label: "localhost:3002",
				storageEngine: inmemoryStorage{
					storageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
		},
	}

	rs := newRoutingServer(2, rh)
	api := newRestAPI(rs)

	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}
