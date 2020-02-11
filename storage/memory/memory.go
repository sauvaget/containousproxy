package memory

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sauvaget/containousproxy/models"
)

type cacheStorage struct {
	c *cache.Cache
}

func NewCacheStorage(cttl time.Duration) *cacheStorage {
	c := cache.New(cttl, cttl)
	s := &cacheStorage{
		c: c,
	}
	return s
}

func (s *cacheStorage) Read(key string) (*models.Cacheitem, error) {
	val, found := s.c.Get(key)
	if !found {
		return nil, models.ErrCacheitemNotfound
	}

	return val.(*models.Cacheitem), nil
}

func (s *cacheStorage) Write(ci models.Cacheitem) error {
	s.c.Set(ci.Key, &ci, cache.DefaultExpiration)
	return nil
}
