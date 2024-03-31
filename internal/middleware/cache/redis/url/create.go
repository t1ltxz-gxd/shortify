package url

import (
	"context"
	"time"
)

// Create is a method that adds a new URL to the cache.
// It takes a context for managing the lifecycle of the operation,
// a hash which is the unique identifier for the URL,
// the actual URL string, and an expiration time for the cache entry.
// It returns an error if the operation fails.
func (c *cache) Create(_ context.Context, hash, url string, expiration time.Duration) error {
	// Set the URL in the cache with the provided hash and expiration time
	err := c.client.Set(hash, url, expiration).Err()
	// If an error occurs, return the error
	if err != nil {
		return err
	}
	// If the operation is successful, return nil
	return nil
}
