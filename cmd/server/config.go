package main

//Config is the main configuration struct.
type Config struct {
	Debug          bool   `default:"false"`
	HTTPListenPort int    `default:"8888"`
	CacheType      string `default:"memory"`
	RedisURI       string `default:"localhost"`
	CacheTTL       string `default:"300000"`
}
