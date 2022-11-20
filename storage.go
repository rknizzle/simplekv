package main

import (
	"io"
	"os"
)

type storage struct {
	writeTo  io.Writer
	readFrom io.Reader
}

func newStorage() storage {
	return storage{
		writeTo:  os.Stdout,
		readFrom: os.Stdin,
	}
}

func (s storage) write(key string, value io.Reader) error {
	// TODO: i want this to block, but not read everything into memeory (stream it instead). Can
	// io.Copy do this? Or do I need to use a pipe or something?
	io.Copy(s.writeTo, value)
	return nil
}
