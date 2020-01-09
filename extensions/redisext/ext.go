package redisext

import (
	"fmt"
	"errors"
	"github.com/go-redis/redis"
	"github.com/shanbay/gobay"
)

// RedisExt redis扩展，处理client的初始化工作
type RedisExt struct {
	NS     string
	app    *gobay.Application
	prefix string
	*redis.Client
}

var _ gobay.Extension = (*RedisExt)(nil)

// Init
func (c *RedisExt) Init(app *gobay.Application) error {
	if c.NS == "" {
		return errors.New("lack of NS")
	}
	c.app = app
	config := gobay.GetConfigByPrefix(app.Config(), c.NS, true)
	host := config.GetString("host")
	password := config.GetString("password")
	dbNum := config.GetInt("db")
	c.prefix = config.GetString("prefix")
	c.Client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       dbNum,
	})
	_, err := c.Client.Ping().Result()
	return err
}

// Object return redis client
func (c *RedisExt) Object() interface{} {
	return c
}

// AddPrefix add prefix to a key
func (c *RedisExt) AddPrefix(key string) string {
	return fmt.Sprintf("%s.%s", c.prefix, key)
}

// Close close redis client
func (c *RedisExt) Close() error {
	return c.Client.Close()
}

// Application
func (c *RedisExt) Application() *gobay.Application {
	return c.app
}
