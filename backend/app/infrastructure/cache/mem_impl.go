package cache

import (
	c "breakfaster/config"

	goc "github.com/patrickmn/go-cache"
)

// MemCache implements the Cache interface
type MemCache struct {
	cache *goc.Cache
}

// Get implements the get method
func (mc *MemCache) Get(key string) (interface{}, bool) {
	return mc.cache.Get(key)
}

// Set implements the set method
func (mc *MemCache) Set(key string, val interface{}) {
	// set to default expiration we specify
	mc.cache.Set(key, val, goc.DefaultExpiration)
}

// Delete implements the delete method
func (mc *MemCache) Delete(key string) {
	mc.cache.Delete(key)
}

// NewMemCache is the factory for MemCache instance
func NewMemCache(config *c.Config) GeneralCache {
	return &MemCache{
		cache: goc.New(config.DefaultCacheExpiration, config.CleanCacheInterval),
	}
}
