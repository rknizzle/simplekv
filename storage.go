package main

import (
	"bytes"
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
