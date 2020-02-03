package models

import "errors"

var ErrCacheitemNotfound = errors.New("cacheitem not found")

type CacheitemRepository interface {
	Read(string) (Cacheitem, error)
	Write(Cacheitem) error
}

type Cacheitem struct {
	Key   string
	Value string
}
