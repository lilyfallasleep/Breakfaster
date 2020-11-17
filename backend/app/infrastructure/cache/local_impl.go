package cache

import (
	c "breakfaster/config"

	goc "github.com/patrickmn/go-cache"
)

// LocalCache implements the Cache interface
type LocalCache struct {
	cache *goc.Cache
}

// Get implements the get method
func (lc *LocalCache) Get(key string) (interface{}, bool) {
	return lc.cache.Get(key)
}

// Set implements the set method
func (lc *LocalCache) Set(key string, val interface{}) {
	// set to default expiration we specify
	lc.cache.Set(key, val, goc.DefaultExpiration)
}

// Delete implements the delete method
func (lc *LocalCache) Delete(key string) {
	lc.cache.Delete(key)
}

// NewLocalCache is the factory for MemCache instance
func NewLocalCache(config *c.Config) GeneralCache {
	return &LocalCache{
		cache: goc.New(config.DefaultCacheExpiration, config.CleanCacheInterval),
	}
}
