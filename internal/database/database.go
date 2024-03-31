package database

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/models"
)

// URLDatabase is an interface that defines the methods for URL database operations.
type URLDatabase interface {
	// ApplyMigrations is a method that applies database migrations.
	// It takes a slice of strings representing the migration files to be applied.
	// It returns an error if the operation fails.
	ApplyMigrations(migrationFiles []string) error

	// Create is a method that adds a new URL to the database.
	// It takes a context for managing the lifecycle of the operation,
	// a url which is the actual URL string, and a hash which is the unique identifier for the URL.
	// It returns an error if the operation fails.
	Create(ctx context.Context, url, hash string) error

	// Get is a method that retrieves a URL from the database using its hash.
	// It takes a context for managing the lifecycle of the operation,
	// and the hash of the URL to retrieve.
	// It returns a pointer to a URL model if the operation is successful,
	// and an error if the operation fails or if the URL is not found in the database.
	Get(ctx context.Context, hash string) (*models.URL, error)
}
