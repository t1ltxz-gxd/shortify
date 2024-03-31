package app

import (
	"github.com/spf13/viper"
	"github.com/t1ltxz-gxd/shortify/internal/api/url"
	"github.com/t1ltxz-gxd/shortify/internal/config"
	pgURL "github.com/t1ltxz-gxd/shortify/internal/database/postgres/url"
	redisURL "github.com/t1ltxz-gxd/shortify/internal/middleware/cache/redis/url"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"github.com/t1ltxz-gxd/shortify/internal/repository"
	urlRepository "github.com/t1ltxz-gxd/shortify/internal/repository/url"
	"github.com/t1ltxz-gxd/shortify/internal/service"
	urlService "github.com/t1ltxz-gxd/shortify/internal/service/url"
	"go.uber.org/zap"
)

// serviceProvider is a struct that holds the dependencies for the service provider.
// It includes a grpcConfig which holds the gRPC configuration,
// a urlRepository which is the URL repository,
// a urlService which is the URL service,
// a urlImpl which is the URL implementation.
type serviceProvider struct {
	grpcConfig    config.GRPCConfig        // grpcConfig holds the gRPC configuration
	urlRepository repository.URLRepository // urlRepository is the URL repository
	urlService    service.URLService       // urlService is the URL service
	urlImpl       *url.Implementation      // urlImpl is the URL implementation
}

// newServiceProvider is a function that creates a new serviceProvider struct.
// It takes a sqlx.DB pointer and a redis.Client pointer as parameters and returns a pointer to a serviceProvider struct.
// It logs that the service provider was initialized and returns the serviceProvider struct.
func newServiceProvider() *serviceProvider {
	logger.Debug("Service provider initialized!")
	return &serviceProvider{}
}

// GRPCConfig is a method on the serviceProvider struct.
// It gets the gRPC configuration for the service provider.
// If the grpcConfig field of the serviceProvider struct is nil, it creates a new gRPC configuration and assigns it to the grpcConfig field.
// If the creation of the gRPC configuration returns an error, it logs the error and exits the application.
// It logs that the gRPC configuration was initialized and returns the gRPC configuration.
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			logger.Fatal("failed to get grpc config", zap.Error(err)) // Log the error and exit the application if the creation of the gRPC configuration returns an error
		}

		s.grpcConfig = cfg // Set the gRPC configuration
	}
	logger.Debug("GRPC config initialized!")

	return s.grpcConfig // Return the gRPC configuration
}

// URLRepository is a method on the serviceProvider struct.
// It gets the URL repository for the service provider.
// If the urlRepository field of the serviceProvider struct is nil, it creates a new URL repository with the database connection and Redis client from the serviceProvider struct and assigns it to the urlRepository field.
// It logs that the URL repository was initialized and returns the URL repository.
func (s *serviceProvider) URLRepository() repository.URLRepository {
	if s.urlRepository == nil {
		db := pgURL.Init()
		// Apply the database migrations by calling the applyMigration method
		err := db.ApplyMigrations(viper.GetStringSlice("migrationFiles"))
		// If the applyMigration method returns an error, return the error
		if err != nil {
			logger.Fatal("failed to apply migrations", zap.Error(err))
		}
		cache := redisURL.Init()
		s.urlRepository = urlRepository.NewRepository(db, cache)
	}
	logger.Debug("URL repository initialized!")

	return s.urlRepository
}

// URLService is a method on the serviceProvider struct.
// It gets the URL service for the service provider.
// If the urlService field of the serviceProvider struct is nil, it creates a new URL service with the URL repository from the serviceProvider struct and assigns it to the urlService field.
// It logs that the URL service was initialized and returns the URL service.
func (s *serviceProvider) URLService() service.URLService {
	if s.urlService == nil {
		s.urlService = urlService.NewService(
			s.URLRepository(),
		)
	}
	logger.Debug("URL service initialized!")

	return s.urlService
}

// URLImpl is a method on the serviceProvider struct.
// It gets the URL implementation for the service provider.
// If the urlImpl field of the serviceProvider struct is nil, it creates a new URL implementation with the URL service from the serviceProvider struct and assigns it to the urlImpl field.
// It logs that the URL implementation was initialized and returns the URL implementation.
func (s *serviceProvider) URLImpl() *url.Implementation {
	if s.urlImpl == nil {
		s.urlImpl = url.NewImplementation(s.URLService())
	}
	logger.Debug("URL implementation initialized!")
	return s.urlImpl
}
