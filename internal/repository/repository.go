package repository

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/models"
)

// URLRepository is an interface that represents a repository for URLs.
// It has two methods: Create and Get.
type URLRepository interface {
	// Create is a method that creates a new URL in the repository.
	// It takes a context, a hash string, and a URL string as parameters.
	// The context is used for request-scoped data, cancellation signals, and deadlines.
	// The hash string is the hashed version of the URL.
	// The URL string is the original URL.
	// It returns an error if the creation fails.
	Create(ctx context.Context, hash string, url string) error

	// Get is a method that retrieves a URL from the repository.
	// It takes a context and a hash string as parameters.
	// The context is used for request-scoped data, cancellation signals, and deadlines.
	// The hash string is the hashed version of the URL.
	// It returns a pointer to a URL model and an error.
	// If the retrieval is successful, the error is nil.
	// If the retrieval fails, the URL model is nil and the error contains the failure reason.
	Get(ctx context.Context, hash string) (*models.URL, error)
}
