package url

import (
	"context"
	"fmt"
	"github.com/speps/go-hashids"
	"github.com/spf13/viper"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"go.uber.org/zap"
)

// Create is a method of the service struct that creates a new URL in the service.
// It takes a context and a URL string as parameters.
// The context is used for request-scoped data, cancellation signals, and deadlines.
// The URL string is the original URL.
// It first logs a debug message that it is creating a new short for the URL.
// It then creates a new hash data with the URL as the salt, the minimum length from the configuration, and the alphabet from the configuration.
// It logs the minimum length and the alphabet.
// It creates a new hash with the hash data.
// It uses the length of the URL as the ID.
// In a real-world application, you would want to use a database auto-incremented ID.
// It generates a unique hash for the ID and logs a debug message that it is generating a hash for the ID.
// It checks if the hash is already in use and logs a debug message that it is checking if the hash is already in use.
// If the hash is already in use, it logs a debug message that the hash is already in use and returns an error.
// If the hash is not in use, it creates a short URL with the host and the gRPC port from the configuration and the hash.
// It returns the short URL and nil.
func (s *service) Create(ctx context.Context, url string) (string, error) {
	logger.Debug("Creating a new short for URL...", zap.String("url", url)) // Log the creation
	hd := hashids.NewData()                                                 // Create new hash data
	hd.Salt = url                                                           // Set the salt
	hd.MinLength = viper.GetInt("app.services.hash.minLength")              // Set the minimum length
	hd.Alphabet = viper.GetString("app.services.hash.alphabet")             // Set the alphabet
	logger.Debug("minLength", zap.Int("minLength", hd.MinLength))           // Log the minimum length
	logger.Debug("alphabet", zap.String("alphabet", hd.Alphabet))           // Log the alphabet
	h, _ := hashids.NewWithData(hd)                                         // Create new hash

	// Here we are using the length of the originalURL as the ID.
	// In a real-world application, you would want to use a database auto-incremented ID.
	id := len(url) // Set the ID

	// Generate a unique hash for the ID
	logger.Debug("Generating a hash for the ID...", zap.Int("id", id)) // Log the generation
	hash, _ := h.Encode([]int{id})                                     // Generate the hash

	// Check if the hash is already in use
	logger.Debug("Checking if the hash is already in use...", zap.String("hash", hash)) // Log the check
	err := s.urlRepository.Create(ctx, hash, url)                                       // Create the URL
	if err != nil {
		logger.Debug("The hash is already in use!", zap.String("hash", hash)) // Log the error
		return "", err                                                        // Return the error
	}
	shortURL := fmt.Sprintf("http://%s:%d/%s", viper.GetString("host"), viper.GetInt("ports.grpc"), hash) // Create the short URL

	return shortURL, nil // Return the short URL
}
