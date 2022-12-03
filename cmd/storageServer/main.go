package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rknizzle/simplekv/pkg/storage"
)

func main() {
	storageEngine, err := storage.NewFileSystemStorage()
	if err != nil {
		panic(err)
	}

	api := storage.NewStorageRESTapi(storageEngine)

	port, err := getPortFromEnv()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}

func getPortFromEnv() (int, error) {
	var port int
	var err error
	envVar, present := os.LookupEnv("PORT")
	if present {
		port, err = strconv.Atoi(envVar)
		if err != nil {
			return 0, err
		}
	} else {
		port = 8000
	}

	return port, nil
}
