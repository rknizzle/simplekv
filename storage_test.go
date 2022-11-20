package main

import (
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	ims := newInmemoryStorage()

	key := "hello"
	value := "world"
	sr := strings.NewReader(value)

	ims.write(key, sr)

	if string(ims.storageMap[key]) != value {
		t.Fatalf("Expected the value to be saved as %s but got %s", string(ims.storageMap[key]), value)
	}
}
