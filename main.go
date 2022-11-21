package main

import (
	"fmt"
	"net/http"
)

func main() {
	// setup the routing server based on the config
	port := 8080

	rh := rendezvousHash{}
	rs := newRoutingServer(2, rh)
	api := newRestAPI(rs)

	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}
