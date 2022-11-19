package main

import (
	"fmt"
	"net/http"
)

type distributedHashAlgo interface {
	// TODO addNode()
	// TODO removeNode()
	getNodesForKey(key string, numReplicas int) (nodes []string)
}

type routingServer struct {
	// number of nodes that each key should be saved to
	numReplicas int

	hash distributedHashAlgo
}

func newRoutingServer() routingServer {
	return routingServer{hash: rendezvousHash{}}
}

func (rs routingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: handle requests here
	fmt.Println(rs)
}

func (rs routingServer) getNodesForKey(key string) (nodes []string) {
	return rs.hash.getNodesForKey(key, rs.numReplicas)
}
