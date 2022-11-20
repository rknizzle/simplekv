package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type distributedHashAlgo interface {
	// TODO addNode()
	// TODO removeNode()
	getNodesForKey(key string, numReplicas int) (nodes []storageNode)
	getAllNodes() []storageNode
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
	nodes := rs.getNodesForKey(key)

	// save the input value stream to a buffer to be used many times to write to all of the node replicas
	// TODO: optimize this
	var buf bytes.Buffer
	_, err := buf.ReadFrom(value)
	if err != nil {
		return err
	}

	for _, node := range nodes {
		byteReader := bytes.NewReader(buf.Bytes())
		// TODO: probably do these concurrently / in parallel
		// NOTE: this could just be a call to save to a map OR an HTTP request depending on the type
		err := node.write(key, byteReader)
		if err != nil {
			return err
		}
	}

	return nil
}
