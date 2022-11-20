package main

import (
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	s := newStorage()
	sr := strings.NewReader("hello world")
	s.write("doesntmatteryet", sr)
}
