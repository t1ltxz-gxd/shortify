package url

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/t1ltxz-gxd/shortify/internal/database"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/cache"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"github.com/t1ltxz-gxd/shortify/internal/models"
	def "github.com/t1ltxz-gxd/shortify/internal/repository"
	"go.uber.org/zap"
	"sync"
	"time"
)

// Ensure that the repository struct implements the URLRepository interface
var _ def.URLRepository = (*repository)(nil)

// repository is a struct that represents a repository for URLs.
// It has three fields: db, cache, and m.
// db is a pointer to a sqlx.DB instance that represents the database connection.
// cache is a pointer to a redis.Client instance that represents the Redis cache.
// m is a sync.RWMutex instance that is used for read/write locking to ensure thread safety.
type repository struct {
	db    database.URLDatabase // The database connection
	cache cache.URLCache       // The cache
	m     sync.RWMutex         // The read/write mutex
}

// NewRepository is a function that creates a new repository.
// It takes a pointer to a sqlx.DB instance and a pointer to a redis.Client instance as parameters.
// The sqlx.DB instance represents the database connection.
// The redis.Client instance represents the Redis cache.
// It returns a pointer to a repository instance.
func NewRepository(db database.URLDatabase, cache cache.URLCache) def.URLRepository {
	return &repository{
		db:    db,    // Set the database connection
		cache: cache, // Set the Redis cache
	}
}

// Create is a method of the repository struct that creates a new URL in the repository.
// It takes a context, a hash string, and a URL string as parameters.
// The context is used for request-scoped data, cancellation signals, and deadlines.
// The hash string is the hashed version of the URL.
// The URL string is the original URL.
// It locks the mutex before creating the URL and unlocks it after the creation.
// It returns an error if the creation fails.
func (r *repository) Create(_ context.Context, hash string, url string) error {
	r.m.Lock()         // Lock the mutex
	defer r.m.Unlock() // Unlock the mutex after the creation

	// The SQL query to insert the URL into the database
	err := r.db.Create(context.Background(), url, hash)
	if err != nil {
		logger.Error("Failed to insert URL into the database", zap.Error(err)) // Log the error if the creation fails
	}
	return err // Return the error
}

// Get is a method of the repository struct that retrieves a URL from the repository.
// It takes a context and a hash string as parameters.
// The context is used for request-scoped data, cancellation signals, and deadlines.
// The hash string is the hashed version of the URL.
// It locks the mutex for reading before retrieving the URL and unlocks it after the retrieval.
// It first tries to get the URL from the Redis cache.
// If the URL is not in the cache, it gets it from the Postgres database.
// If the URL is in the database, it saves it in the Redis cache and returns it.
// If the URL is not in the database, it returns nil.
// If the retrieval from the cache or the database fails, it logs an error and returns the error.
func (r *repository) Get(_ context.Context, hash string) (*models.URL, error) {
	r.m.RLock()         // Lock the mutex for reading
	defer r.m.RUnlock() // Unlock the mutex after the retrieval

	// Try to get the URL from the Redis cache
	logger.Debug("Fetching URL from cache", zap.String("hash", hash))
	val, err := r.cache.Get(context.Background(), hash)
	if errors.Is(err, redis.Nil) {
		// If the URL is not in the cache, get it from the Postgres database
		logger.Debug("Fetching URL from database", zap.String("hash", hash))
		url, err := r.db.Get(context.Background(), hash)
		//err := r.db.Get(&url, "SELECT * FROM urls WHERE hash = $1", hash)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// If the URL is not in the database, return nil
				logger.Error("URL is not found in the database", zap.String("hash", hash))
				return nil, nil
			}
			logger.Error("Failed to fetch URL from the database", zap.String("hash", hash), zap.Error(err))
			return nil, err
		}

		// Save the URL in the Redis cache
		ttl := viper.GetUint("app.services.hash.ttlCache")
		err = r.cache.Create(context.Background(), hash, url.Original, time.Duration(ttl)*time.Second)
		if err != nil {
			logger.Error("Failed to save URL in the cache", zap.Error(err))
			return nil, err
		}

		// Return the URL
		logger.Debug("URL is fetched from the database", zap.String("url", url.Original))
		return url, nil
	} else if err != nil {
		logger.Error("Failed to fetch URL from the cache", zap.String("hash", hash), zap.Error(err))
		return nil, err
	}

	// If the URL is in the cache, return it
	logger.Debug("URL is fetched from the cache",
		zap.String("hash", val.Hash),
		zap.String("url", val.Original))
	return val, nil
}
