package main

import "time"

//Config is the main configuration struct.
type Config struct {
	Debug          bool          `default:"false"`
	HTTPListenPort int           `default:"8888"`
	CacheType      string        `default:"memory"`
	RedisURI       string        `default:"redis"`
	CacheTTL       time.Duration `default:"5m"`
	Proxy          map[string]string
}
