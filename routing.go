package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/http"
	"sort"
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

type rendezvousHash struct {
	nodes []string
}

func (r rendezvousHash) getNodesForKey(key string, numReplicas int) (nodes []string) {
	// rendezvous distributed hashing algorithm
	// https://en.wikipedia.org/wiki/Rendezvous_hashing

	// give a score to each node for the given key
	var scores sortableNodeScores
	for _, node := range r.nodes {
		singleScore := scoreForNode(key, node)
		ns := nodeScore{node, singleScore}
		scores = append(scores, ns)
	}

	// sort the nodes by score
	sort.Sort(scores)

	// Grab the top numReplica scores
	var nodesForKey []string
	for i := 0; i < numReplicas; i++ {
		nodesForKey = append(nodesForKey, scores[i].node)
	}

	return nodesForKey
}

func scoreForNode(key string, node string) []byte {
	hash := md5.New()
	hash.Write([]byte(key))
	hash.Write([]byte(node))
	score := hash.Sum(nil)
	return score
}

// sortable list of nodes by hash score
type nodeScore struct {
	// label for the node
	node string
	// hash score given to the node for a particular key
	score []byte
}

type sortableNodeScores []nodeScore

func (s sortableNodeScores) Len() int      { return len(s) }
func (s sortableNodeScores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortableNodeScores) Less(i, j int) bool {
	return bytes.Compare(s[i].score, s[j].score) == 1
}
