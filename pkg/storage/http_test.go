package storage

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStorageHTTPgetSuccessful(t *testing.T) {

	// preload the node with the key 'hello'
	storageEngine := NewInmemoryStorage()
	storageEngine.StorageMap["hello"] = []byte("world")

	api := NewStorageRESTapi(storageEngine)

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

func TestStorageHTTPgetMissingKey(t *testing.T) {
	storageEngine := NewInmemoryStorage()
	api := NewStorageRESTapi(storageEngine)

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create the new request")
	}

	rec := httptest.NewRecorder()

	api.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("Expected 500 code but got %d", rec.Code)
	}

	// TODO: standardize  this error message with the routing servers
	if !strings.Contains(rec.Body.String(), "doesnt exist") {
		t.Fatalf("Expected an error about the key not existing but instead got %s", rec.Body.String())
	}
}

func TestStorageHTTPwriteSuccessful(t *testing.T) {
	storageEngine := NewInmemoryStorage()
	api := NewStorageRESTapi(storageEngine)

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
