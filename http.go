package main

import "net/http"

type restAPI struct {
	rs routingServer
}

func newRestAPI(rs routingServer) restAPI {
	return restAPI{rs: rs}
}

func (api restAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: handle requests here
}
