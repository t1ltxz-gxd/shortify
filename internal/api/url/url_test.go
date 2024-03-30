package url_test

import (
	"context"
	"errors"
	"github.com/t1ltxz-gxd/shortify/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/t1ltxz-gxd/shortify/internal/api/url"
	desc "github.com/t1ltxz-gxd/shortify/pkg/url_v1"
)

// MockURLService is a struct that mocks the URLService interface for testing.
// It embeds the mock.Mock struct from the testify/mock package.
type MockURLService struct {
	mock.Mock
}

// Create is a method that mocks the Create method of the URLService interface.
// It takes a context and a URL string as parameters.
// The context is used for request-scoped data, cancellation signals, and deadlines.
// The URL string is the original URL.
// It returns a hash string that represents the hashed version of the URL and an error.
// The hash string and the error are the return values of the Called method of the mock.Mock struct.
func (m *MockURLService) Create(ctx context.Context, url string) (string, error) {
	args := m.Called(ctx, url)
	return args.String(0), args.Error(1)
}

// Get is a method that mocks the Get method of the URLService interface.
// It takes a context and a hash string as parameters.
// The context is used for request-scoped data, cancellation signals, and deadlines.
// The hash string is the hashed version of the URL.
// It returns a pointer to a URL model and an error.
// The URL model and the error are the return values of the Called method of the mock.Mock struct.
func (m *MockURLService) Get(ctx context.Context, hash string) (*models.URL, error) {
	args := m.Called(ctx, hash)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

// TestGet_Success is a test function that tests the successful retrieval of a URL from the service.
// It creates a new MockURLService and sets the expected return value of the Get method to a URL model and nil.
// It creates a new Implementation with the MockURLService and a GetRequest with a valid hash.
// It calls the Get method of the Implementation with the GetRequest and checks if the returned URL is the expected URL and if the error is nil.
// It checks if the expectations of the MockURLService were met.
func TestGet_Success(t *testing.T) {
	mockService := new(MockURLService)
	mockService.On("Get", mock.Anything, "validHash").Return(&models.URL{Original: "https://example.com"}, nil)

	impl := url.NewImplementation(mockService)
	req := &desc.GetRequest{Hash: "validHash"}

	resp, err := impl.Get(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", resp.Url)
	mockService.AssertExpectations(t)
}

// TestGet_Error is a test function that tests the failed retrieval of a URL from the service.
// It creates a new MockURLService and sets the expected return value of the Get method to nil and an error.
// It creates a new Implementation with the MockURLService and a GetRequest with an invalid hash.
// It calls the Get method of the Implementation with the GetRequest and checks if the returned URL is nil and if the error is not nil.
// It checks if the expectations of the MockURLService were met.
func TestGet_Error(t *testing.T) {
	mockService := new(MockURLService)
	mockService.On("Get", mock.Anything, "invalidHash").Return(nil, errors.New("error"))

	impl := url.NewImplementation(mockService)
	req := &desc.GetRequest{Hash: "invalidHash"}

	resp, err := impl.Get(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockService.AssertExpectations(t)
}

// TestCreate_Success is a test function that tests the successful creation of a URL in the service.
// It creates a new MockURLService and sets the expected return value of the Create method to a hash string and nil.
// It creates a new Implementation with the MockURLService and a CreateRequest with a valid URL.
// It calls the Create method of the Implementation with the CreateRequest and checks if the returned short URL is the expected short URL and if the error is nil.
// It checks if the expectations of the MockURLService were met.
func TestCreate_Success(t *testing.T) {
	mockService := new(MockURLService)
	mockService.On("Create", mock.Anything, "https://example.com").Return("hash123", nil)

	impl := url.NewImplementation(mockService)
	req := &desc.CreateRequest{Url: "https://example.com"}

	resp, err := impl.Create(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "hash123", resp.ShortUrl)
	mockService.AssertExpectations(t)
}

// TestCreate_Error is a test function that tests the failed creation of a URL in the service.
// It creates a new MockURLService and sets the expected return value of the Create method to an empty string and an error.
// It creates a new Implementation with the MockURLService and a CreateRequest with an invalid URL.
// It calls the Create method of the Implementation with the CreateRequest and checks if the returned short URL is an empty string and if the error is not nil.
// It checks if the expectations of the MockURLService were met.
func TestCreate_Error(t *testing.T) {
	mockService := new(MockURLService)
	mockService.On("Create", mock.Anything, "https://invalid.com").Return("", errors.New("error"))

	impl := url.NewImplementation(mockService)
	req := &desc.CreateRequest{Url: "https://invalid.com"}

	resp, err := impl.Create(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockService.AssertExpectations(t)
}

// BenchmarkCreate is a benchmark test for the Create method of the Implementation struct.
// It measures the performance of the Create method by calling it B.N times in a loop.
// B.N is automatically adjusted by the testing package to get meaningful results.
// The Create method is called with a fixed CreateRequest, so the benchmark measures only the performance of the method itself, not the performance of its dependencies.
// The dependencies of the Create method are mocked using the MockURLService.
func BenchmarkCreate(b *testing.B) {
	// Create a new MockURLService
	mockService := new(MockURLService)
	// Set up the Create method of the mock service to return a fixed hash and no error
	mockService.On("Create", mock.Anything, "https://example.com").Return("hash123", nil)

	// Create an Implementation instance with the mock service
	impl := url.NewImplementation(mockService)

	// Create an instance of CreateRequest
	req := &desc.CreateRequest{
		Url: "https://example.com",
	}

	// Start a loop that will be executed B.N times
	for i := 0; i < b.N; i++ {
		// Call the Create method of the Implementation and discard the results
		_, _ = impl.Create(context.Background(), req)
	}
}

// BenchmarkGet is a benchmark test for the Get method of the Implementation struct.
// It measures the performance of the Get method by calling it B.N times in a loop.
// B.N is automatically adjusted by the testing package to get meaningful results.
// The Get method is called with a fixed GetRequest, so the benchmark measures only the performance of the method itself, not the performance of its dependencies.
// The dependencies of the Get method are mocked using the MockURLService.
func BenchmarkGet(b *testing.B) {
	// Create a new MockURLService
	mockService := new(MockURLService)
	// Set up the Get method of the mock service to return a fixed URL and no error
	mockService.On("Get", mock.Anything, "exampleHash").Return(&models.URL{Original: "https://example.com"}, nil)

	// Create an Implementation instance with the mock service
	impl := url.NewImplementation(mockService)

	// Create a GetRequest instance
	req := &desc.GetRequest{
		Hash: "exampleHash",
	}

	// Start a loop that will be executed B.N times
	for i := 0; i < b.N; i++ {
		// Call the Get method of the Implementation and discard the results
		_, _ = impl.Get(context.Background(), req)
	}
}
