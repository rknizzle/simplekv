package main

import (
	"fmt"
	"net/http"
)

type routingServer struct {
	nodes []string
	// number of nodes that each key should be saved to
	numReplicas int
}

func newRoutingServer() routingServer {
	return routingServer{}
}

func (rs routingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: handle requests here
	fmt.Println(rs)
}
