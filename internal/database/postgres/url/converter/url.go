package converter

import (
	repoModels "github.com/t1ltxz-gxd/shortify/internal/database/postgres/url/models"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"github.com/t1ltxz-gxd/shortify/internal/models"
	"go.uber.org/zap"
)

// ToURLFromRepo is a function that converts a URL from the repository model to the service model.
// It takes a URL from the repository model as a parameter.
// The URL from the repository model has the original URL, the hash, the time when the URL was added, and the time when the URL was last updated.
// It logs a debug message that it is converting the URL from the repository model to the service model.
// It returns a pointer to a URL from the service model.
// The URL from the service model has the original URL, the hash, the time when the URL was added, and a pointer to the time when the URL was last updated.
// If the URL from the repository model has not been updated, the pointer to the time when the URL was last updated is nil.
func ToURLFromRepo(url repoModels.URL) *models.URL {
	logger.Debug("Converting URL from repository to service", zap.String("original", url.Original), zap.String("short", url.Hash)) // Log the conversion
	return &models.URL{
		Original:  url.Original,        // Set the original URL
		Hash:      url.Hash,            // Set the hash
		AddedAt:   url.AddedAt,         // Set the time when the URL was added
		UpdatedAt: &url.UpdatedAt.Time, // Set the pointer to the time when the URL was last updated
	}
}
