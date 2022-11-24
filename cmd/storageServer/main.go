package main

import (
	"fmt"
	"net/http"

	"github.com/rknizzle/simplekv/pkg/storage"
)

func main() {
	storageEngine := storage.NewFileSystemStorage()
	api := storage.NewStorageRESTapi(storageEngine)

	port := 8000
	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}
