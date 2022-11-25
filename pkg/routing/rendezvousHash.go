package routing

import (
	"bytes"
	"crypto/md5"
	"sort"
)

// rendezvousHash implements the rendezvous distributed hashing algorithm.
// Visit https://en.wikipedia.org/wiki/Rendezvous_hashing for more info
type RendezvousHash struct {
	Nodes []StorageNode
}

func (r RendezvousHash) getAllNodes() (nodes []StorageNode) {
	return r.Nodes
}

func (r RendezvousHash) getNodesForKey(key string, numReplicas int) (nodes []StorageNode) {
	// give a score to each node for the given key
	var scores sortableNodeScores
	for _, node := range r.Nodes {
		singleScore := scoreForNode(key, node)
		ns := nodeScore{node, singleScore}
		scores = append(scores, ns)
	}

	// sort the nodes by score
	sort.Sort(scores)

	// Grab the top numReplica scores
	var nodesForKey []StorageNode
	for i := 0; i < numReplicas; i++ {
		nodesForKey = append(nodesForKey, scores[i].node)
	}

	return nodesForKey
}

func scoreForNode(key string, node StorageNode) []byte {
	// TODO: decouple the specific hashing method from the rendezvousHash algorithm probably by
	// passing adding a hashMethod func as a type to the struct
	hash := md5.New()
	hash.Write([]byte(key))
	hash.Write([]byte(node.GetLabel()))
	score := hash.Sum(nil)
	return score
}

// sortable list of nodes by hash score
type nodeScore struct {
	// label for the node
	node StorageNode
	// hash score given to the node for a particular key
	score []byte
}

type sortableNodeScores []nodeScore

func (s sortableNodeScores) Len() int      { return len(s) }
func (s sortableNodeScores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortableNodeScores) Less(i, j int) bool {
	return bytes.Compare(s[i].score, s[j].score) == 1
}
