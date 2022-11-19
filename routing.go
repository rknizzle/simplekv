package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/http"
	"sort"
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

func (rs routingServer) getNodesForKey(key string) (nodes []string) {
	// rendezvous distributed hashing algorithm
	// https://en.wikipedia.org/wiki/Rendezvous_hashing

	// give a score to each node for the given key
	var scores sortableNodeScores
	for _, node := range rs.nodes {
		singleScore := scoreForNode(key, node)
		ns := nodeScore{node, singleScore}
		scores = append(scores, ns)
	}

	// sort the nodes by score
	sort.Sort(scores)

	// Grab the top numReplica scores
	var nodesForKey []string
	for i := 0; i < rs.numReplicas; i++ {
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
