package cache

// RedisCache is the interface for redis cache
type RedisCache interface {
	Get(key string, dst interface{}) (bool, error)
	Set(key string, val interface{}) error
	Delete(key string) error
	ExecPipeLine(cmds *[]Cmd) error
}
