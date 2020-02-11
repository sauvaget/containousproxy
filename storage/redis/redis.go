package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/sauvaget/containousproxy/models"
)

type cacheStorage struct {
	client *redis.Client
	cttl   time.Duration
}

func NewCacheStorage(client *redis.Client, cttl time.Duration) *cacheStorage {
	s := &cacheStorage{
		client: client,
		cttl:   cttl,
	}
	return s
}

func (s *cacheStorage) Read(key string) (*models.Cacheitem, error) {
	val, err := s.client.Get(key).Result()
	if err != nil {
		return nil, models.ErrCacheitemNotfound
	}

	cacheitem := models.Cacheitem{}
	err = json.Unmarshal([]byte(val), &cacheitem)
	if err != nil {
		return nil, err
	}
	return &cacheitem, nil
}

func (s *cacheStorage) Write(ci models.Cacheitem) error {
	b, err := json.Marshal(ci)
	if err != nil {
		return err
	}

	redisErr := s.client.Set(ci.Key, b, s.cttl)
	if redisErr.Err() != nil {
		return redisErr.Err()
	}
	return nil
}
