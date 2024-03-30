package service

import (
	"context"

	"github.com/t1ltxz-gxd/shortify/internal/models"
)

// URLService is an interface that represents a service for URLs.
// It has two methods: Create and Get.
type URLService interface {
	// Create is a method that creates a new URL in the service.
	// It takes a context and a URL string as parameters.
	// The context is used for request-scoped data, cancellation signals, and deadlines.
	// The URL string is the original URL.
	// It returns a hash string that represents the hashed version of the URL and an error.
	// If the creation is successful, the error is nil.
	// If the creation fails, the hash string is empty and the error contains the failure reason.
	Create(ctx context.Context, url string) (string, error)

	// Get is a method that retrieves a URL from the service.
	// It takes a context and a hash string as parameters.
	// The context is used for request-scoped data, cancellation signals, and deadlines.
	// The hash string is the hashed version of the URL.
	// It returns a pointer to a URL model and an error.
	// If the retrieval is successful, the error is nil.
	// If the retrieval fails, the URL model is nil and the error contains the failure reason.
	Get(ctx context.Context, hash string) (*models.URL, error)
}
