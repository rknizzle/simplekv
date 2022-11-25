package routing

import (
	"io"

	"github.com/rknizzle/simplekv/pkg/storage"
)

type StorageNode interface {
	Write(key string, value io.Reader) error
	Get(key string) (io.Reader, error)
	GetLabel() (label string)
}

type TestStorageNode struct {
	Label         string
	StorageEngine storage.StorageEngine
}

func (s TestStorageNode) Write(key string, value io.Reader) error {
	return s.StorageEngine.Write(key, value)
}

func (s TestStorageNode) Get(key string) (io.Reader, error) {
	return s.StorageEngine.Get(key)
}

func (s TestStorageNode) GetLabel() (label string) {
	return s.Label
}
