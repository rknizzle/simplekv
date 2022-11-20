package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type inmemoryStorage struct {
	storageMap map[string][]byte
}

func newInmemoryStorage() inmemoryStorage {
	return inmemoryStorage{storageMap: make(map[string][]byte)}
}

func (ims inmemoryStorage) write(key string, value io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(value)
	if err != nil {
		return err
	}

	ims.storageMap[key] = buf.Bytes()
	return nil
}

func (ims inmemoryStorage) get(key string) (io.Reader, error) {
	value, ok := ims.storageMap[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Key: %s doesnt exist in storage", key))
	}

	reader := bytes.NewReader(value)
	return reader, nil
}
