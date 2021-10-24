package persistent

import (
	"github.com/go-redis/redis/v8"
	"shopping-cart/common"
)

// Database shares global database instance
var (
	Database Redis
)

// Redis manages Redis connection
type Redis struct {
	Client *redis.Client
}

// Init initializes redis database
func (r *Redis) Init() error {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     common.K8sConfig.Out.Redis.Host,
		Password: common.K8sConfig.Out.Redis.Password,
		DB:       0, // use default DB
	})
	return nil
}

// Close the existing connection
func (r *Redis) Close() {
	if r.Client != nil {
		r.Client.Close()
	}
}
