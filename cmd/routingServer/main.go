package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/rknizzle/simplekv/pkg/routing"
)

type storageNodeFlags []string

func (s *storageNodeFlags) String() string { return "" }
func (s *storageNodeFlags) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	replicas := flag.Int("replicas", 2, "Number of replicas to save data to")
	port := flag.Int("port", 8080, "Port for the HTTP server to listen on")

	var storageNodeURLs storageNodeFlags
	flag.Var(&storageNodeURLs, "storage", "The URL of a storage node")

	flag.Parse()

	storageNodes := URLsToStorageNodes(storageNodeURLs)

	rh := routing.RendezvousHash{
		Nodes: storageNodes,
	}

	rs := routing.NewRoutingServer(*replicas, rh)
	api := routing.NewRestAPI(rs)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), api)
}

func URLsToStorageNodes(urls []string) []routing.StorageNode {
	var storageNodes []routing.StorageNode
	for _, url := range urls {
		s := routing.RemoteStorageNode{
			URL: url,
		}
		storageNodes = append(storageNodes, s)
	}

	return storageNodes
}
