package url

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"github.com/t1ltxz-gxd/shortify/internal/models"
	"go.uber.org/zap"
)

// Get is a method of the service struct that retrieves a URL from the service.
// It takes a context and a hash string as parameters.
// The context is used for request-scoped data, cancellation signals, and deadlines.
// The hash string is the hashed version of the URL.
// It first tries to get the original URL from the repository.
// If the retrieval from the repository fails, it logs an error and returns the error.
// If the original URL is not in the repository, it logs an error and returns an invalid URL error.
// If the original URL is in the repository, it returns the original URL.
func (s *service) Get(ctx context.Context, hash string) (*models.URL, error) {
	// Get the original URL from repository
	logger.Debug("Fetching URL from repository", zap.String("hash", hash))
	originalURL, err := s.urlRepository.Get(ctx, hash)
	if err != nil {
		logger.Error("Failed to fetch URL from repository", zap.String("hash", hash), zap.Error(err))
		return nil, err
	}

	// Type assert the original URL to string
	if originalURL == nil {
		logger.Error("Original URL is not found", zap.String("hash", hash))
		return nil, models.ErrorInvalidURL
	}
	return originalURL, nil
}
