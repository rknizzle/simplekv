package routing

import (
	"encoding/json"
	"io"
	"net/http"
)

type restAPI struct {
	rs routingServer
}

func NewRestAPI(rs routingServer) restAPI {
	return restAPI{rs: rs}
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

func (api restAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]

	if r.Method == "GET" {
		valueReader, err := api.rs.get(key)
		if err != nil {
			respondWithError(err.Error(), w)
			return
		}
		defer valueReader.Close()

		// stream the value to the response
		io.Copy(w, valueReader)
	} else if r.Method == "PUT" {
		err := api.rs.saveValueToKey(key, r.Body)
		if err != nil {
			respondWithError(err.Error(), w)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(nil)
	}
}
