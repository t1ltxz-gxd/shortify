package logger

import (
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Constants for environment names
const (
	envDev         = "dev"         // Development environment short name
	envDevelopment = "development" // Development environment full name
	envProd        = "prod"        // Production environment short name
	envProduction  = "production"  // Production environment full name
)

// zapLog is a global variable that holds the logger instance
var zapLog *zap.Logger

// Init is a function that initializes the logger.
// It takes the environment name as a parameter.
// Depending on the environment, it sets up the logger for development or production.
// For development, it uses the built-in development configuration of zap.
// For production, it creates a new logger with a file and console syncer, and a JSON encoder.
// If the environment is not recognized, it logs a fatal error.
// Finally, it names the logger with the name from the configuration.
func Init(env string) {
	var w zapcore.WriteSyncer
	var encoder zapcore.Encoder
	switch env {
	case envDev, envDevelopment:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		zapLog, _ = config.Build(zap.AddCallerSkip(1))
	case envProd, envProduction:
		fileSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   viper.GetString("log.file"),
			MaxSize:    viper.GetInt("logger"), // megabytes
			MaxBackups: viper.GetInt("logger.maxBackups"),
			Compress:   viper.GetBool("logger.compress"),
			MaxAge:     viper.GetInt("logger.maxAge"), // days
		})
		consoleSyncer := zapcore.AddSync(os.Stdout)
		w = zapcore.NewMultiWriteSyncer(fileSyncer, consoleSyncer)
		config := zap.NewProductionConfig()
		encoder = zapcore.NewJSONEncoder(config.EncoderConfig)
		core := zapcore.NewCore(encoder, w, config.Level)
		zapLog = zap.New(core, zap.AddCallerSkip(1))
	default:
		err := errors.New("unknown environment")
		if err != nil {
			zapLog.Fatal("Failed to initialize logger", zap.Error(err))
		}
	}
	zapLog = zapLog.Named(viper.GetString("logger.name"))
}

// Sync is a function that syncs the logger.
// It returns an error if the sync fails.
func Sync() error {
	return zapLog.Sync()
}

// Info is a function that logs an info level message.
// It takes a message and a variadic parameter of fields.
func Info(message string, fields ...zap.Field) {
	zapLog.Info(message, fields...)
}

// Debug is a function that logs a debug level message.
// It takes a message and a variadic parameter of fields.
func Debug(message string, fields ...zap.Field) {
	zapLog.Debug(message, fields...)
}

// Error is a function that logs an error level message.
// It takes a message and a variadic parameter of fields.
func Error(message string, fields ...zap.Field) {
	zapLog.Error(message, fields...)
}

// Fatal is a function that logs a fatal level message.
// It takes a message and a variadic parameter of fields.
// After logging the message, it calls os.Exit(1) to terminate the program.
func Fatal(message string, fields ...zap.Field) {
	zapLog.Fatal(message, fields...)
}
