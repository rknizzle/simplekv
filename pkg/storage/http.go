package storage

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type storageEngine interface {
	write(key string, value io.Reader) error
	get(key string) (io.Reader, error)
}

type storageRESTapi struct {
	se storageEngine
}

func NewStorageRESTapi(se storageEngine) storageRESTapi {
	return storageRESTapi{se: se}
}

func respondWithError(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = message

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("{\"message\":\"Error\"}"))
		return
	}

	w.Write(jsonResp)
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

func getKeyFromURL(path string) string {
	indexOfLastSlash := strings.LastIndex(path, "/")
	key := path[indexOfLastSlash+1:]
	return key
}
