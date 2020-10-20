package cache

// GeneralCache is the interface for memory caching
type GeneralCache interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
	Delete(string)
}
