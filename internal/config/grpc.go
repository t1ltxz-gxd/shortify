package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net"
	"os"
	"strconv"
)

// GRPCConfig is an interface that defines the methods required for a gRPC configuration.
type GRPCConfig interface {
	// Address returns the address of the gRPC server as a string.
	Address() string
}

// grpcConfig is a struct that holds the host and port for a gRPC server.
type grpcConfig struct {
	host string // host is the hostname of the gRPC server.
	port int    // port is the port number on which the gRPC server is running.
}

// NewGRPCConfig is a function that creates a new gRPC configuration.
// It reads the host and port from the environment variables using viper.
// If the host is not found, it returns an error.
// Otherwise, it returns a GRPCConfig interface and nil error.
func NewGRPCConfig() (GRPCConfig, error) {
	host := viper.GetString("host")
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert grpc port to int")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

// Address is a method on the grpcConfig struct.
// It returns the address of the gRPC server by joining the host and port.
func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, strconv.Itoa(cfg.port))
}
