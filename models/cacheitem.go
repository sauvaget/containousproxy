package models

import (
	"errors"
	"net/http"
)

var ErrCacheitemNotfound = errors.New("cacheitem not found")

type CacheitemRepository interface {
	Read(string) (*Cacheitem, error)
	Write(Cacheitem) error
}

type Cacheitem struct {
	Key    string      `json="key"`
	Header http.Header `json="header"`
	Value  string      `json="value"`
}
