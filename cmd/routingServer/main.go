package main

import (
	"fmt"
	"net/http"

	"github.com/rknizzle/simplekv/pkg/routing"
	"github.com/rknizzle/simplekv/pkg/storage"
)

func main() {
	// setup the routing server based on the config
	port := 8080

	key := "hello"
	value := "world"

	// preload a key in the nodes just for testing
	rh := routing.RendezvousHash{
		Nodes: []routing.StorageNode{
			routing.TestStorageNode{
				Label: "localhost:3000",
				StorageEngine: storage.InmemoryStorage{
					StorageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
			routing.TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			routing.TestStorageNode{
				Label: "localhost:3002",
				StorageEngine: storage.InmemoryStorage{
					StorageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
		},
	}

	rs := routing.NewRoutingServer(2, rh)
	api := routing.NewRestAPI(rs)

	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}
