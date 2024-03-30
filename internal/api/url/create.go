package url

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/converter"
	desc "github.com/t1ltxz-gxd/shortify/pkg/url_v1"
)

// Implementation is a struct that implements the URL service interface.

// Create is a method on the Implementation struct.
// It takes a context and a CreateRequest as parameters.
// The CreateRequest contains the URL to be shortened.
// This method calls the Create method on the urlService, passing the context and the URL from the request.
// The URL from the request is converted from a descriptor URL to a service URL using the ToURLFromDesc function from the converter package.
// If the Create method on the urlService returns an error, the Create method returns nil and the error.
// If the Create method on the urlService does not return an error, the Create method returns a CreateResponse containing the shortened URL and nil error.
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	// Call the Create method on the urlService, passing the context and the URL from the request.
	// The URL from the request is converted from a descriptor URL to a service URL using the ToURLFromDesc function from the converter package.
	shortURL, err := i.urlService.Create(ctx, converter.ToURLFromDesc(req.Url))
	// If the Create method on the urlService returns an error, return nil and the error.
	if err != nil {
		return nil, err
	}

	// If the Create method on the urlService does not return an error, return a CreateResponse containing the shortened URL and nil error.
	return &desc.CreateResponse{
		ShortUrl: shortURL,
	}, nil
}
