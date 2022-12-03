package storage

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type fileSystemStorage struct {
	dir string
}

func NewFileSystemStorage() (fileSystemStorage, error) {
	err := os.Mkdir("./store", os.FileMode(0777))
	if err != nil {
		return fileSystemStorage{}, err
	}

	return fileSystemStorage{dir: "store"}, nil
}

func (se fileSystemStorage) Write(key string, value io.Reader) error {
	file, err := os.Create(fmt.Sprintf("./%s/%s", se.dir, key))
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
	file, err := os.Open(fmt.Sprintf("./%s/%s", se.dir, key))
	if err != nil {

		// 'key not found' case
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil, keyNotFoundError{key}
		}

		return nil, err
	}

	return file, nil
}
