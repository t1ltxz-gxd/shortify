package cache

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/models"
	"time"
)

// URLCache is an interface that defines the methods for URL caching.
type URLCache interface {
	// Create is a method that adds a new URL to the cache.
	// It takes a context for managing the lifecycle of the operation,
	// a hash which is the unique identifier for the URL,
	// the actual URL string, and an expiration time for the cache entry.
	// It returns an error if the operation fails.
	Create(ctx context.Context, hash, url string, expiration time.Duration) error

	// Get is a method that retrieves a URL from the cache using its hash.
	// It takes a context for managing the lifecycle of the operation,
	// and the hash of the URL to retrieve.
	// It returns a pointer to a URL model if the operation is successful,
	// and an error if the operation fails or if the URL is not found in the cache.
	Get(ctx context.Context, hash string) (*models.URL, error)
}
