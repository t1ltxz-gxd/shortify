package url

import (
	"context"
	"database/sql"
	"errors"
	"github.com/t1ltxz-gxd/shortify/internal/database/postgres/url/converter"
	repoModel "github.com/t1ltxz-gxd/shortify/internal/database/postgres/url/models"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"github.com/t1ltxz-gxd/shortify/internal/models"
	"go.uber.org/zap"
)

// Get is a method that retrieves a URL from the database using its hash.
// It takes a context for managing the lifecycle of the operation,
// and the hash of the URL to retrieve.
// It first logs a debug message indicating that it is fetching the URL from the database.
// It then attempts to retrieve the URL from the database using the provided hash.
// If an error occurs, it checks if the error is due to the URL not being found in the database.
// If the URL is not found, it logs an error message and returns nil for both the URL and the error.
// If the error is due to another issue, it logs an error message and returns nil for the URL and the error.
// If the URL is successfully retrieved, it logs a debug message indicating that the URL was fetched from the database,
// and it converts the retrieved URL from the repository model to the application model.
// It returns a pointer to the URL model if the operation is successful,
// and an error if the operation fails or if the URL is not found in the database.
func (d *database) Get(_ context.Context, hash string) (*models.URL, error) {
	var url repoModel.URL
	logger.Debug("Fetching URL from database", zap.String("hash", hash))
	err := d.db.Get(&url, "SELECT * FROM urls WHERE hash = $1", hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// If the URL is not in the database, return nil
			logger.Error("URL is not found in the database", zap.String("hash", hash))
			return nil, nil
		}
		logger.Error("Failed to fetch URL from the database", zap.String("hash", hash), zap.Error(err))
		return nil, err
	}
	logger.Debug("URL is fetched from the database", zap.String("url", url.Original))
	return converter.ToURLFromRepo(url), nil
}
