package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPgetSuccessful(t *testing.T) {
	key := "hello"
	value := "world"

	// preload the nodes with the 'hello' key
	rh := rendezvousHash{
		nodes: []storageNode{
			testStorageNode{
				label: "localhost:3000",
				storageEngine: inmemoryStorage{
					storageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
			testStorageNode{
				label:         "localhost:3001",
				storageEngine: newInmemoryStorage(),
			},
			testStorageNode{
				label: "localhost:3002",
				storageEngine: inmemoryStorage{
					storageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
		},
	}

	rs := newRoutingServer(2, rh)
	api := newRestAPI(rs)

	// get the hello key
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create the new request")
	}

	rec := httptest.NewRecorder()

	api.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 code but got %d", rec.Code)
	}

	if rec.Body.String() != "world" {
		t.Fatalf("Expected the response body to contain world but got %s", rec.Body.String())
	}
}
