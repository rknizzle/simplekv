package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type InmemoryStorage struct {
	StorageMap map[string][]byte
}

func NewInmemoryStorage() InmemoryStorage {
	return InmemoryStorage{StorageMap: make(map[string][]byte)}
}

func (ims InmemoryStorage) Write(key string, value io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(value)
	if err != nil {
		return err
	}

	ims.StorageMap[key] = buf.Bytes()
	return nil
}

func (ims InmemoryStorage) Get(key string) (io.Reader, error) {
	value, ok := ims.StorageMap[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Key: '%s' doesnt exist in storage", key))
	}

	reader := bytes.NewReader(value)
	return reader, nil
}
