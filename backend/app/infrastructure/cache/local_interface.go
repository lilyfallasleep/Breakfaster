package cache

// GeneralCache is the interface for memory caching
type GeneralCache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
	Delete(key string)
}
