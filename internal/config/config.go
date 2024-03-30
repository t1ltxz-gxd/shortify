package config

import (
	"github.com/spf13/viper"
	"strings"
)

// Config is a struct that holds the configuration for the application.
// It includes the domain, environment files, logger configuration, application configuration, and port configuration.
type Config struct {
	Host         string   `mapstructure:"host"` // Domain is the domain name for the application.
	PostgresHost string   `mapstructure:"postgresHost"`
	RedisHost    string   `mapstructure:"redisHost"`
	RedisDB      int      `mapstructure:"redisDB"`
	EnvFiles     []string `mapstructure:"env-files"` // EnvFiles is a list of environment files to be loaded.
	Logger       Logger   `mapstructure:"logger"`    // Logger is the logger configuration.
	App          App      `mapstructure:"app"`       // App is the application configuration.
	Ports        Ports    `mapstructure:"ports"`     // Ports is the port configuration.
}

// Ports is a struct that holds the HTTP and gRPC port numbers.
type Ports struct {
	HTTP int `mapstructure:"http"` // HTTP is the HTTP port number.
	GRPC int `mapstructure:"grpc"` // GRPC is the gRPC port number.
}

// Logger is a struct that holds the logger name and file syncer configuration.
type Logger struct {
	Name       string     `mapstructure:"name"`       // Name is the name of the logger.
	FileSyncer FileSyncer `mapstructure:"fileSyncer"` // FileSyncer is the file syncer configuration.
}

// FileSyncer is a struct that holds the file syncer configuration.
type FileSyncer struct {
	Filename   string `mapstructure:"Filename"`   // Filename is the name of the log file.
	MaxSize    int    `mapstructure:"MaxSize"`    // MaxSize is the maximum size of the log file.
	MaxBackups int    `mapstructure:"MaxBackups"` // MaxBackups is the maximum number of backup log files to keep.
	Compress   bool   `mapstructure:"Compress"`   // Compress indicates whether to compress the log files.
	MaxAge     int    `mapstructure:"MaxAge"`     // MaxAge is the maximum age of the log files.
}

// App is a struct that holds the services configuration.
type App struct {
	Services Services `mapstructure:"services"` // Services is the services configuration.
}

// Services is a struct that holds the hash configuration.
type Services struct {
	Hash Hash `mapstructure:"hash"` // Hash is the hash configuration.
}

// Hash is a struct that holds the hash configuration.
type Hash struct {
	TTLCache  int    `mapstructure:"ttlCache"`  // TTLCache is the time-to-live for the cache.
	MinLength int    `mapstructure:"minLength"` // MinLength is the minimum length of the hash.
	Alphabet  string `mapstructure:"alphabet"`  // Alphabet is the set of characters to use in the hash.
}

// LoadConfig is a function that loads the configuration for the application.
// It takes the path, name, and type of the configuration file as parameters and returns an error.
// It sets the path, name, and type of the configuration file using the viper package.
// It then automatically uses environment variables if set and replaces "." in environment variables with "_".
// It reads the configuration file and unmarshals the read configuration into a Config struct.
// If the read or unmarshal returns an error, LoadConfig returns the error.
// If the read and unmarshal do not return an error, LoadConfig returns nil.
func LoadConfig(configPath, configName, configType string) error {
	viper.AddConfigPath(configPath) // Set the path for the config file
	viper.SetConfigName(configName) // Set the name for the config file
	viper.SetConfigType(configType) // Set the type for the config file

	viper.AutomaticEnv()                                   // Automatically use environment variables if set
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`)) // Replace "." in env vars with "_"

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	// Unmarshal the read config into a Config struct
	var cfg *Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}
	return nil
}
