package app

import (
	"context"
	// reviving the pq driver
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/t1ltxz-gxd/shortify/internal/config"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	desc "github.com/t1ltxz-gxd/shortify/pkg/url_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"runtime"
)

// App is a struct that holds the dependencies for the application.
// It includes a serviceProvider which provides the services for the application,
// a grpcServer which is the gRPC server for the application,
// a db which is the database connection for the application,
// and a cache which is the Redis client for the application.
type App struct {
	serviceProvider *serviceProvider // serviceProvider provides the services for the application
	grpcServer      *grpc.Server     // grpcServer is the gRPC server for the application
}

// NewApp is a function that creates a new App struct.
// It takes a context as a parameter and returns a pointer to an App struct and an error.
// It initializes the dependencies of the App struct by calling the initDeps method.
// If the initDeps method returns an error, NewApp returns nil and the error.
// If the initDeps method does not return an error, NewApp returns a pointer to the App struct and nil error.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{} // Create a new App struct

	// Initialize the dependencies of the App struct
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err // Return nil and the error if the initDeps method returns an error
	}

	return a, nil // Return the App struct and nil error if the initDeps method does not return an error
}

// Run is a method on the App struct.
// It starts the gRPC server by calling the runGRPCServer method and returns any error that it returns.
func (a *App) Run() error {
	return a.runGRPCServer() // Start the gRPC server and return any error that it returns
}

// initDeps is a method on the App struct.
// It initializes the dependencies of the App struct.
// It takes a context as a parameter and returns an error.
// It creates a slice of functions that initialize the dependencies of the App struct.
// These functions are initConfig, initLogger, initCache, initDatabase, initServiceProvider, and initGRPCServer.
// It then iterates over the slice of functions and calls each function, passing the context as a parameter.
// If any of the functions return an error, initDeps returns the error.
// If none of the functions return an error, initDeps applies the database migrations by calling the applyMigration method.
// If the applyMigration method returns an error, initDeps returns the error.
// If the applyMigration method does not return an error, initDeps returns nil.
func (a *App) initDeps(ctx context.Context) error {
	// Create a slice of functions that initialize the dependencies of the App struct
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	// Iterate over the slice of functions and call each function, passing the context as a parameter
	for _, f := range inits {
		err := f(ctx)
		// If any of the functions return an error, return the error
		if err != nil {
			return err
		}
	}

	// If the applyMigration method does not return an error, return nil
	return nil
}

// initConfig is a method on the App struct.
// It initializes the configuration for the application.
// It takes a context as a parameter and returns an error.
// It calls the LoadDotEnv method from the config package, passing ".env" as the parameter.
// This method loads the environment variables from the .env file.
// If the LoadDotEnv method returns an error, initConfig returns the error.
// If the LoadDotEnv method does not return an error, initConfig calls the LoadConfig method from the config package, passing "config", "config", and "yml" as the parameters.
// This method loads the configuration from the config.yml file.
// If the LoadConfig method returns an error, initConfig returns the error.
// If the LoadConfig method does not return an error, initConfig returns nil.
func (a *App) initConfig(_ context.Context) error {
	err := config.LoadConfig("config", "config", "yml")
	if err != nil {
		return err
	}
	err = config.LoadDotEnv(viper.GetStringSlice("envFiles")...)
	if err != nil {
		return err
	}
	return nil
}

// initLogger is a method on the App struct.
// It initializes the logger for the application.
// It takes a context as a parameter and returns an error.
// It calls the Init method from the logger package, passing the environment variable "ENV" as the parameter.
// This method initializes the logger with the specified environment.
// After the logger is initialized, initLogger logs the system information (OS, architecture, Go version, and environment) using the Info method from the logger package.
// It also logs that debug mode is enabled using the Debug method from the logger package.
// initLogger then returns nil.
func (a *App) initLogger(_ context.Context) error {
	logger.Init(os.Getenv("ENV"))

	// Recording system information after successful logger initialization
	logger.Info("Successfully start!",
		zap.String("OS", runtime.GOOS),
		zap.String("Architecture", runtime.GOARCH),
		zap.String("Go version", runtime.Version()),
		zap.String("Environment", viper.GetString("env")))
	logger.Debug("Debug mode enabled!")

	return nil
}

// initServiceProvider is a method on the App struct.
// It initializes the service provider for the application.
// It takes a context as a parameter and returns an error.
// It creates a new service provider with the database connection and Redis client from the App struct.
// It then assigns the service provider to the serviceProvider field of the App struct.
// initServiceProvider then returns nil.
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// initGRPCServer is a method on the App struct.
// It initializes the gRPC server for the application.
// It takes a context as a parameter and returns an error.
// It creates a new gRPC server with insecure credentials.
// It then registers the gRPC server for reflection and the URL service implementation from the service provider.
// It logs that the gRPC server was initialized.
// initGRPCServer then returns nil.
func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterUrlV1Server(a.grpcServer, a.serviceProvider.URLImpl())

	logger.Info("gRPC server initialized!")

	return nil
}

// runGRPCServer is a method on the App struct.
// It starts the gRPC server for the application.
// It logs that the gRPC server is running with the address from the gRPC configuration of the service provider.
// It then listens for TCP connections on the address from the gRPC configuration of the service provider.
// If the listen returns an error, runGRPCServer returns the error.
// If the listen does not return an error, runGRPCServer serves the gRPC server on the listener.
// If the serve returns an error, runGRPCServer returns the error.
// If the serve does not return an error, runGRPCServer returns nil.
func (a *App) runGRPCServer() error {
	logger.Info("GRPC server is running", zap.String("address", a.serviceProvider.GRPCConfig().Address()))

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
