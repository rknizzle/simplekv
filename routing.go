package main

import (
	"fmt"
	"io"
	"net/http"
)

type distributedHashAlgo interface {
	// TODO addNode()
	// TODO removeNode()
	getNodesForKey(key string, numReplicas int) (nodes []storageNode)
}

type routingServer struct {
	// number of nodes that each key should be saved to
	numReplicas int

	hash distributedHashAlgo
}

func newRoutingServer(numReplicas int, hash distributedHashAlgo) routingServer {
	return routingServer{numReplicas: numReplicas, hash: hash}
}

func (rs routingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: handle requests here
	fmt.Println(rs)
}

func (rs routingServer) getNodesForKey(key string) (nodes []storageNode) {
	return rs.hash.getNodesForKey(key, rs.numReplicas)
}

func (rs routingServer) saveValueToKey(key string, value io.Reader) error {
	// TODO: should this return a slice of nodes rather than a slice of strings labels for nodes?
	nodes := rs.getNodesForKey(key)

	for _, node := range nodes {
		// TODO: probably do these concurrently / in parallel
		// NOTE: this could just be a call to save to a map OR an HTTP request depending on the type
		err := node.write(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}
