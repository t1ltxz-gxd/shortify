package url

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	def "github.com/t1ltxz-gxd/shortify/internal/middleware/cache"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"go.uber.org/zap"
	"os"
)

// URLCache is an interface that defines the methods for URL caching.
var _ def.URLCache = (*cache)(nil)

// cache is a struct that implements the URLCache interface.
// It contains a client for interacting with the Redis server.
type cache struct {
	client *redis.Client
}

// Init is a function that initializes a new cache.
// It creates a new Redis client with the server address, password, and database number
// specified in the application's configuration.
// It then pings the Redis server to ensure the connection is successful.
// If the connection fails, it logs a fatal error.
// If the connection is successful, it returns a new cache with the Redis client.
func Init() def.URLCache {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redisHost"), viper.GetInt("ports.redis")), // the address of the Redis server
		Password: os.Getenv("REDIS_PASS"),                                                         // password (if required)
		DB:       viper.GetInt("RedisDB"),                                                         // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		logger.Fatal("failed to connect to Redis", zap.Error(err))
	}
	return &cache{
		client: client,
	}
}
