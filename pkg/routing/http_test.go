package routing

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rknizzle/simplekv/pkg/storage"
)

func TestHTTPgetSuccessful(t *testing.T) {
	key := "hello"
	value := "world"

	// preload the nodes with the 'hello' key
	rh := RendezvousHash{
		Nodes: []StorageNode{
			TestStorageNode{
				Label: "localhost:3000",
				StorageEngine: storage.InmemoryStorage{
					StorageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
			TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label: "localhost:3002",
				StorageEngine: storage.InmemoryStorage{
					StorageMap: map[string][]byte{
						key: []byte(value),
					},
				},
			},
		},
	}

	rs := NewRoutingServer(2, rh)
	api := NewRestAPI(rs)

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

func TestHTTPgetMissingKey(t *testing.T) {
	rh := RendezvousHash{
		Nodes: []StorageNode{
			TestStorageNode{
				Label:         "localhost:3000",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3002",
				StorageEngine: storage.NewInmemoryStorage(),
			},
		},
	}

	rs := NewRoutingServer(2, rh)

	api := NewRestAPI(rs)

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create the new request")
	}

	rec := httptest.NewRecorder()

	api.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("Expected 500 code but got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Failed to get key") {
		t.Fatalf("Expected an error about failing to get key but instead got %s", rec.Body.String())
	}
}

func TestHTTPwriteSuccessful(t *testing.T) {
	rh := RendezvousHash{
		Nodes: []StorageNode{
			TestStorageNode{
				Label:         "localhost:3000",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3001",
				StorageEngine: storage.NewInmemoryStorage(),
			},
			TestStorageNode{
				Label:         "localhost:3002",
				StorageEngine: storage.NewInmemoryStorage(),
			},
		},
	}

	rs := NewRoutingServer(2, rh)

	api := NewRestAPI(rs)

	reqBody := strings.NewReader("world")
	req, err := http.NewRequest("POST", "/hello", reqBody)
	if err != nil {
		t.Fatalf("Failed to create the new request")
	}

	rec := httptest.NewRecorder()

	api.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("Expected 201 code but got %d", rec.Code)
	}
}
