package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/kelseyhightower/envconfig"
	"github.com/sauvaget/containousproxy/models"
	"github.com/sauvaget/containousproxy/pkg/proxy"
	"github.com/sauvaget/containousproxy/storage/memory"
	redisStorage "github.com/sauvaget/containousproxy/storage/redis"
)

type server struct {
	config       Config
	proxyService proxy.Service
}

func main() {
	err := runServer()
	if err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	s := &server{}
	err := envconfig.Process("cproxy", &s.config)
	if err != nil {
		return err
	}

	// Proxy config
	s.config.Proxy = make(map[string]string)
	s.config.Proxy["localhost:8888"] = "https://docs.traefik.io"
	s.config.Proxy["127.0.0.1:8888"] = "https://containo.us"

	cacheStorage, err := initCacheStorage(&s.config)
	proxyService := proxy.NewService(&http.Client{}, cacheStorage, s.config.Proxy)
	s.proxyService = proxyService

	log.Printf("Starting Server on :%d\n", s.config.HTTPListenPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.HTTPListenPort), s.router())
}

func initCacheStorage(c *Config) (models.CacheitemRepository, error) {
	switch c.CacheType {
	case "memory":
		storage := memory.NewCacheStorage(c.CacheTTL)
		return storage, nil
	case "redis":
		client := redis.NewClient(&redis.Options{
			Addr:     c.RedisURI,
			Password: "",
			DB:       0,
		})
		storage := redisStorage.NewCacheStorage(client, c.CacheTTL)
		return storage, nil
	default:
		return nil, errors.New("unknown cachetype")
	}
}
