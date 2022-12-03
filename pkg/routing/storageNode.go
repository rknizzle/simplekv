package routing

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rknizzle/simplekv/pkg/storage"
)

type StorageNode interface {
	Write(key string, value io.Reader) error
	Get(key string) (io.ReadCloser, error)
	GetLabel() (label string)
}

type TestStorageNode struct {
	Label         string
	StorageEngine storage.StorageEngine
}

func (s TestStorageNode) Write(key string, value io.Reader) error {
	return s.StorageEngine.Write(key, value)
}

func (s TestStorageNode) Get(key string) (io.ReadCloser, error) {
	return s.StorageEngine.Get(key)
}

func (s TestStorageNode) GetLabel() (label string) {
	return s.Label
}

type RemoteStorageNode struct {
	URL string
}

func (s RemoteStorageNode) Write(key string, value io.Reader) error {
	url := fmt.Sprintf("%s/%s", s.URL, key)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, value)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// TODO: parse out response to check for errors

	return nil
}

func (s RemoteStorageNode) Get(key string) (io.ReadCloser, error) {
	url := fmt.Sprintf("%s/%s", s.URL, key)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//defer resp.Body.Close()
	// TODO: how do I make sure that the response gets closed later?

	return resp.Body, nil
}

func (s RemoteStorageNode) GetLabel() (label string) {
	return s.URL
}
