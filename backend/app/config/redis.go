package config

import "time"

// RedisConfig is the configuration for redis
type RedisConfig struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	MaxRetries  int
	IdleTimeout time.Duration
}
