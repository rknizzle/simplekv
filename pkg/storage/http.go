package storage

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type storageRESTapi struct {
	se StorageEngine
}

func NewStorageRESTapi(se StorageEngine) storageRESTapi {
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
		valueReader, err := api.se.Get(key)
		if err != nil {
			respondWithError(err.Error(), w)
			return
		}
		defer valueReader.Close()

		// stream the value to the response
		io.Copy(w, valueReader)
	} else if r.Method == "PUT" {
		err := api.se.Write(key, r.Body)
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
