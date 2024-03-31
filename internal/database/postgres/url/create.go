package url

import (
	"context"
	"database/sql"
	repoModel "github.com/t1ltxz-gxd/shortify/internal/database/postgres/url/models"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"go.uber.org/zap"
	"time"
)

// Create is a method that adds a new URL to the database.
// It takes a context for managing the lifecycle of the operation,
// a url which is the actual URL string, and a hash which is the unique identifier for the URL.
// It first constructs the SQL query to insert the URL into the database.
// It then executes the query, passing in a new URL model with the provided URL, hash, and the current time for the added and updated timestamps.
// If an error occurs during the execution of the query, it logs an error message and returns the error.
// If the operation is successful, it returns nil.
func (d *database) Create(_ context.Context, url string, hash string) error {
	// The SQL query to insert the URL into the database
	query := `INSERT INTO urls (original_url, hash) VALUES (:original_url, :hash)`
	_, err := d.db.NamedExec(query, &repoModel.URL{
		Original:  url,                                         // Set the original URL
		Hash:      hash,                                        // Set the hash
		AddedAt:   time.Now(),                                  // Set the time when the URL was added
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true}, // Set the time when the URL was updated
	})
	if err != nil {
		logger.Error("Failed to insert URL into the database", zap.Error(err)) // Log the error if the creation fails
		return err
	}
	return nil
}
