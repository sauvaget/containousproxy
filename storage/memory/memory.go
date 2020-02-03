package memory

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sauvaget/containousproxy/models"
)

type storage struct {
	c *cache.Cache
}

func NewStorage(cttl time.Duration) *storage {
	c := cache.New(cttl, cttl)
	s := &storage{
		c: c,
	}
	return s
}

func (s *storage) Read(key string) (*models.Cacheitem, error) {
	val, found := s.c.Get(key)
	if !found {
		return nil, models.ErrCacheitemNotfound
	}
	ci := &models.Cacheitem{
		Key:   key,
		Value: val.(string),
	}
	return ci, nil
}

func (s *storage) Write(ci models.Cacheitem) error {
	s.c.Set(ci.Key, ci.Value, cache.DefaultExpiration)
	return nil
}
