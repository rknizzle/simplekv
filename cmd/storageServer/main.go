package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rknizzle/simplekv/pkg/storage"
)

func main() {
	storageEngine := storage.NewFileSystemStorage()
	api := storage.NewStorageRESTapi(storageEngine)

	var port int
	var err error
	envVar, present := os.LookupEnv("PORT")
	if present {
		port, err = strconv.Atoi(envVar)
		if err != nil {
			panic(err)
		}
	} else {
		port = 8000
	}

	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}
