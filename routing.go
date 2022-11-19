package main

import (
	"fmt"
	"net/http"
)

type routingServer struct {
	// TODO: system config goes here
}

func newRoutingServer() routingServer {
	return routingServer{}
}

func (rs routingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: handle requests here
	fmt.Println(rs)
}
