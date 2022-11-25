package storage

import (
	"io"
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	ims := NewInmemoryStorage()

	key := "hello"
	value := "world"
	sr := strings.NewReader(value)

	ims.Write(key, sr)

	if string(ims.StorageMap[key]) != value {
		t.Fatalf("Expected the value to be saved as %s but got %s", string(ims.StorageMap[key]), value)
	}
}

func TestGet(t *testing.T) {
	ims := NewInmemoryStorage()

	key := "hello"
	value := "world"

	ims.StorageMap = map[string][]byte{
		key: []byte(value),
	}

	valueReader, err := ims.Get(key)
	if err != nil {
		t.Fatalf("inmemoryStorage.get() failed to get the reader with message: %s", err.Error())
	}

	val, err := io.ReadAll(valueReader)
	if err != nil {
		t.Fatalf("Failed to read the value from the reader returned by inmemoryStorage.get() with message: %s", err.Error())
	}

	if string(val) != value {
		t.Fatalf("Expected the value %s for key %s, but instead got the value: %s", value, key, string(val))
	}
}
