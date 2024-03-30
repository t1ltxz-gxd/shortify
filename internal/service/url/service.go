package url

import (
	"github.com/t1ltxz-gxd/shortify/internal/repository"
	def "github.com/t1ltxz-gxd/shortify/internal/service"
)

// Ensure that the service struct implements the URLService interface
var _ def.URLService = (*service)(nil)

// service is a struct that represents a service for URLs.
// It has one field: urlRepository.
// urlRepository is an instance of the URLRepository interface that represents the repository for URLs.
type service struct {
	urlRepository repository.URLRepository // The repository for URLs
}

// NewService is a function that creates a new service for URLs.
// It takes an instance of the URLRepository interface as a parameter.
// The URLRepository instance represents the repository for URLs.
// It returns an instance of the URLService interface.
// The URLService instance represents the service for URLs.
func NewService(
	urlRepository repository.URLRepository, // The repository for URLs
) def.URLService {
	return &service{
		urlRepository: urlRepository, // Set the repository for URLs
	}
}
