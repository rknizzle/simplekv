package main

import (
	"io"
	"net/http"
)

type storageEngine interface {
	write(key string, value io.Reader) error
	get(key string) (io.Reader, error)
}

type storageRESTapi struct {
	se storageEngine
}

func newStorageRESTapi(se storageEngine) storageRESTapi {
	return storageRESTapi{se: se}
}

func (api storageRESTapi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := getKeyFromURL(r.URL.Path)

	if r.Method == "GET" {
		valueReader, err := api.se.get(key)
		if err != nil {
			respondWithError(err.Error(), w)
			return
		}

		// stream the value to the response
		io.Copy(w, valueReader)

	} else if r.Method == "POST" {
		err := api.se.write(key, r.Body)
		if err != nil {
			respondWithError(err.Error(), w)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(nil)
	}
}
