package storage

import (
	"fmt"
	"io"
)

type StorageEngine interface {
	Write(key string, value io.Reader) error
	Get(key string) (io.ReadCloser, error)
}

type keyNotFoundError struct {
	key string
}

func (k keyNotFoundError) Error() string {
	return fmt.Sprintf("key: '%s' not found", k.key)
}
