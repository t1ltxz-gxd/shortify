package url

import (
	"github.com/t1ltxz-gxd/shortify/internal/service"
	desc "github.com/t1ltxz-gxd/shortify/pkg/url_v1"
)

// Implementation is a struct that embeds the UnimplementedUrlV1Server interface from the url_v1 package
// and includes a URLService from the internal service package.
// This struct is used to implement the methods defined in the UnimplementedUrlV1Server interface.
type Implementation struct {
	desc.UnimplementedUrlV1Server                    // Embedding the UnimplementedUrlV1Server interface
	urlService                    service.URLService // URLService from the internal service package
}

// NewImplementation is a function that creates a new Implementation struct.
// It takes a URLService as a parameter and returns a pointer to an Implementation struct.
// The URLService is assigned to the urlService field of the Implementation struct.
func NewImplementation(urlService service.URLService) *Implementation {
	return &Implementation{
		urlService: urlService, // Assigning the URLService to the urlService field of the Implementation struct
	}
}
