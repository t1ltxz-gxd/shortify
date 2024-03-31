package url

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"github.com/t1ltxz-gxd/shortify/internal/models"
	"go.uber.org/zap"
)

// Get is a method that retrieves a URL from the cache using its hash.
// It takes a context for managing the lifecycle of the operation,
// and the hash of the URL to retrieve.
// It returns a pointer to a URL model if the operation is successful,
// and an error if the operation fails or if the URL is not found in the cache.
func (c *cache) Get(_ context.Context, hash string) (*models.URL, error) {
	// Attempt to get the URL from the cache using the provided hash
	val, err := c.client.Get(hash).Result()
	// If an error occurs, log the error and return nil and the error
	if err != nil {
		logger.Error("Failed to fetch URL from the cache", zap.Error(err))
		return nil, err
	}
	// If the URL is successfully retrieved, log the URL and return a pointer to the URL model and nil for the error
	logger.Debug("URL is fetched from the cache", zap.String("url", val))
	return &models.URL{Original: val}, nil
}
