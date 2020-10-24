package cache

import (
	c "breakfaster/config"
	exc "breakfaster/pkg/exception"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache is the redis cache client type
type RedisCache struct {
	client     *redis.ClusterClient
	ctx        context.Context
	expiration time.Duration
}

// RedisOpType is the redis operation type
type RedisOpType int

const (
	// SET represents set operation
	SET RedisOpType = iota
	// DELETE represents delete operation
	DELETE
)

// RedisPayload is a abstract interface for payload type
type RedisPayload interface {
	Payload()
}

// RedisSetPayload is the payload type for set method
type RedisSetPayload struct {
	RedisPayload
	Key string
	Val interface{}
}

// RedisDeletePayload is the payload type for delete method
type RedisDeletePayload struct {
	RedisPayload
	Key string
}

// Payload implements abstract interface
func (RedisSetPayload) Payload() {}

// Payload implements abstract interface
func (RedisDeletePayload) Payload() {}

// Cmd represents an operation and its payload
type Cmd struct {
	OpType  RedisOpType
	Payload RedisPayload
}

type tmpPiplineCmd struct {
	OpType RedisOpType
	Cmd    interface{}
}

// NewRedisCache is the factory for MemCache instance
func NewRedisCache(config *c.Config) (*RedisCache, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:         []string{config.RedisConfig.Addr},
		Password:      config.RedisConfig.Password,
		PoolSize:      config.RedisConfig.PoolSize,
		MaxRetries:    config.RedisConfig.MaxRetries,
		IdleTimeout:   config.RedisConfig.IdleTimeout,
		ReadOnly:      true,
		RouteRandomly: true,
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err == redis.Nil || err != nil {
		return nil, err
	}
	config.Logger.ContextLogger.WithField("type", "setup:redis").Info("successful Redis Connection: " + pong)
	return &RedisCache{
		client:     client,
		ctx:        ctx,
		expiration: config.DefaultCacheExpiration,
	}, nil
}

// Get method returns true if the key already exists and set dst to the corresponding value
func (rc *RedisCache) Get(key string, dst interface{}) (bool, error) {
	val, err := rc.client.Get(rc.ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		json.Unmarshal([]byte(val), dst)
	}
	return true, nil
}

// Set method set a key-value pair
func (rc *RedisCache) Set(key string, val interface{}) error {
	strVal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	if err := rc.client.Set(rc.ctx, key, strVal, rc.expiration).Err(); err != nil {
		return err
	}
	return nil
}

// Delete method deletes a key
func (rc *RedisCache) Delete(key string) error {
	if err := rc.client.Del(rc.ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

// ExecPipeLine execute the given commands in a pipline
func (rc *RedisCache) ExecPipeLine(cmds *[]Cmd) error {
	pipe := rc.client.Pipeline()
	var tmpPiplineCmds []tmpPiplineCmd
	for _, cmd := range *cmds {
		switch cmd.OpType {
		case SET:
			strVal, err := json.Marshal(cmd.Payload.(RedisSetPayload).Val)
			if err != nil {
				return err
			}
			tmpPiplineCmds = append(tmpPiplineCmds, tmpPiplineCmd{
				OpType: SET,
				Cmd:    pipe.Set(rc.ctx, cmd.Payload.(RedisSetPayload).Key, strVal, rc.expiration),
			})
		case DELETE:
			tmpPiplineCmds = append(tmpPiplineCmds, tmpPiplineCmd{
				OpType: DELETE,
				Cmd:    pipe.Del(rc.ctx, cmd.Payload.(RedisDeletePayload).Key),
			})
		default:
			return exc.ErrRedisCmdNotFound
		}
	}
	_, err := pipe.Exec(rc.ctx)
	if err != nil {
		return err
	}

	for _, executedCmd := range tmpPiplineCmds {
		switch executedCmd.OpType {
		case SET:
			if err := executedCmd.Cmd.(*redis.StatusCmd).Err(); err != nil {
				return err
			}
		case DELETE:
			if err := executedCmd.Cmd.(*redis.IntCmd).Err(); err != nil {
				return err
			}
		}
	}
	return nil
}
