package main

import (
	"fmt"
	"net/http"
)

func main() {
	// setup the routing server based on the config
	port := 8080

	rh := rendezvousHash{}
	rs := newRoutingServer(rh)
	http.ListenAndServe(fmt.Sprintf(":%d", port), &rs)
}
