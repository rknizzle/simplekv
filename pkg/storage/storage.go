package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type StorageNode interface {
	Write(key string, value io.Reader) error
	Get(key string) (io.Reader, error)
	GetLabel() (label string)
}

type TestStorageNode struct {
	Label         string
	StorageEngine storageEngine
}

func (s TestStorageNode) Write(key string, value io.Reader) error {
	return s.StorageEngine.write(key, value)
}

func (s TestStorageNode) Get(key string) (io.Reader, error) {
	return s.StorageEngine.get(key)
}

func (s TestStorageNode) GetLabel() (label string) {
	return s.Label
}

type InmemoryStorage struct {
	StorageMap map[string][]byte
}

func NewInmemoryStorage() InmemoryStorage {
	return InmemoryStorage{StorageMap: make(map[string][]byte)}
}

func (ims InmemoryStorage) write(key string, value io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(value)
	if err != nil {
		return err
	}

	ims.StorageMap[key] = buf.Bytes()
	return nil
}

func (ims InmemoryStorage) get(key string) (io.Reader, error) {
	value, ok := ims.StorageMap[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Key: '%s' doesnt exist in storage", key))
	}

	reader := bytes.NewReader(value)
	return reader, nil
}
