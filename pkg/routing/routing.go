package routing

import (
	"bytes"
	"errors"
	"fmt"
	"io"
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

func (rs routingServer) getNodesForKey(key string) (nodes []storageNode) {
	return rs.hash.getNodesForKey(key, rs.numReplicas)
}

func (rs routingServer) saveValueToKey(key string, value io.Reader) error {
	nodes := rs.getNodesForKey(key)

	// save the input value stream to a buffer to be used many times to write to all of the node replicas
	// TODO: optimize this & send all writes concurrently
	var buf bytes.Buffer
	_, err := buf.ReadFrom(value)
	if err != nil {
		return err
	}

	for _, node := range nodes {
		byteReader := bytes.NewReader(buf.Bytes())
		err := node.write(key, byteReader)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rs routingServer) get(key string) (io.Reader, error) {
	nodes := rs.getNodesForKey(key)

	var valueReader io.Reader
	var err error

	for _, node := range nodes {
		valueReader, err = node.get(key)
		if err != nil {
			// TODO: keep track of the errors from each server
		}
	}

	if valueReader != nil {
		return valueReader, nil
	} else {
		// TODO: include the error messages from the servers
		return nil, errors.New(fmt.Sprintf("Failed to get key %s", key))
	}
}
