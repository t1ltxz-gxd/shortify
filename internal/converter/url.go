package converter

import (
	"github.com/t1ltxz-gxd/shortify/internal/models"
	desc "github.com/t1ltxz-gxd/shortify/pkg/url_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToURLFromService is a function that converts a URL model to a URL protobuf message.
// It takes a pointer to a URL model as a parameter and returns a pointer to a URL protobuf message.
// It creates a timestamp for the UpdatedAt field of the URL protobuf message if the UpdatedAt field of the URL model is not nil.
// It then creates a new URL protobuf message with the OriginalUrl, ShortUrl, CreatedAt, and UpdatedAt fields from the URL model and returns it.
func ToURLFromService(url *models.URL) *desc.Url {
	var updatedAt *timestamppb.Timestamp
	if url.UpdatedAt != nil {
		updatedAt = timestamppb.New(*url.UpdatedAt)
	}

	return &desc.Url{
		OriginalUrl: url.Original,
		ShortUrl:    url.Hash,
		CreatedAt:   timestamppb.New(url.AddedAt),
		UpdatedAt:   updatedAt,
	}
}

// ToURLFromDesc is a function that converts a URL string to a URL string.
// It takes a URL string as a parameter and returns a URL string.
// It simply returns the URL string that was passed as a parameter.
func ToURLFromDesc(url string) string {
	return url
}
