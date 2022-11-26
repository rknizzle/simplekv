package main

import (
	"fmt"
	"net/http"

	"github.com/rknizzle/simplekv/pkg/routing"
)

func main() {
	// setup the routing server based on the config
	port := 8080

	rh := routing.RendezvousHash{
		Nodes: []routing.StorageNode{
			routing.RemoteStorageNode{URL: "http://storage_server_1:8000"},
			routing.RemoteStorageNode{URL: "http://storage_server_2:8000"},
			routing.RemoteStorageNode{URL: "http://storage_server_3:8000"},
		},
	}

	rs := routing.NewRoutingServer(2, rh)
	api := routing.NewRestAPI(rs)

	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}
