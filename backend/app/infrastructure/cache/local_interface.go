package cache

// LocalCache is the interface for local memory caching
type LocalCache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
	Delete(key string)
}
