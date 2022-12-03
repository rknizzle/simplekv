package storage

import (
	"io"
	"os"
	"strings"
)

type fileSystemStorage struct {
	// NOTE: There could be an abstraction to the file system here so I could mock it out and test
	// these methods.
	// Alternatively I can just mock at the StorageEngine level and test the system using InmemoryStorage
}

func NewFileSystemStorage() fileSystemStorage {
	return fileSystemStorage{}
}

func (se fileSystemStorage) Write(key string, value io.Reader) error {
	file, err := os.Create(key)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, value)
	if err != nil {
		return err
	}

	return nil
}

func (se fileSystemStorage) Get(key string) (io.ReadCloser, error) {
	file, err := os.Open(key)
	if err != nil {

		// 'key not found' case
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil, keyNotFoundError{key}
		}

		return nil, err
	}

	return file, nil
}
