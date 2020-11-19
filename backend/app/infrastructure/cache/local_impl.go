package cache

import (
	c "breakfaster/config"

	goc "github.com/patrickmn/go-cache"
)

// LocalCacheImpl implements the Cache interface
type LocalCacheImpl struct {
	cache *goc.Cache
}

// Get implements the get method
func (lc *LocalCacheImpl) Get(key string) (interface{}, bool) {
	return lc.cache.Get(key)
}

// Set implements the set method
func (lc *LocalCacheImpl) Set(key string, val interface{}) {
	// set to default expiration we specify
	lc.cache.Set(key, val, goc.DefaultExpiration)
}

// Delete implements the delete method
func (lc *LocalCacheImpl) Delete(key string) {
	lc.cache.Delete(key)
}

// NewLocalCache is the factory for LocalCacheImpl
func NewLocalCache(config *c.Config) LocalCache {
	return &LocalCacheImpl{
		cache: goc.New(config.DefaultCacheExpiration, config.CleanCacheInterval),
	}
}
