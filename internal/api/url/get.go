package url

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/converter"
	desc "github.com/t1ltxz-gxd/shortify/pkg/url_v1"
)

// Get is a method on the Implementation struct.
// It takes a context and a GetRequest as parameters.
// The GetRequest contains the hash of the URL to be retrieved.
// This method calls the Get method on the urlService, passing the context and the hash from the request.
// If the Get method on the urlService returns an error, the Get method returns nil and the error.
// If the Get method on the urlService does not return an error, the Get method returns a GetResponse containing the original URL and nil error.
// The URL returned by the urlService is converted from a service URL to a descriptor URL using the ToURLFromService function from the converter package.
// The original URL from the descriptor URL is then retrieved using the GetOriginalUrl method.
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	// Call the Get method on the urlService, passing the context and the hash from the request.
	url, err := i.urlService.Get(ctx, req.Hash)
	// If the Get method on the urlService returns an error, return nil and the error.
	if err != nil {
		return nil, err
	}
	// If the Get method on the urlService does not return an error, return a GetResponse containing the original URL and nil error.
	// The URL returned by the urlService is converted from a service URL to a descriptor URL using the ToURLFromService function from the converter package.
	// The original URL from the descriptor URL is then retrieved using the GetOriginalUrl method.
	return &desc.GetResponse{
		Url: converter.ToURLFromService(url).GetOriginalUrl(),
	}, nil
}
