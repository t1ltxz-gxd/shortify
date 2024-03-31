package url

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	def "github.com/t1ltxz-gxd/shortify/internal/database"
	"github.com/t1ltxz-gxd/shortify/internal/middleware/logger"
	"go.uber.org/zap"
	"os"
)

var _ def.URLDatabase = (*database)(nil)

type database struct {
	db *sqlx.DB // The database connection
}

func Init() def.URLDatabase {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		viper.GetString("postgresHost"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}
	return &database{
		db: db, // Set the database connection
	}
}

// ApplyMigrations is a method on the App struct.
// It applies the database migrations for the application.
// It takes a sqlx.DB pointer and a slice of migration file paths as parameters and returns an error.
// It iterates over the slice of migration file paths and reads each file.
// If the read returns an error, applyMigration returns the error.
// If the read does not return an error, applyMigration executes the migration on the database.
// If the execution returns an error, applyMigration returns the error.
// If the execution does not return an error, applyMigration logs that the migration was applied successfully.
// After all migrations have been applied, applyMigration returns nil.
func (d *database) ApplyMigrations(migrationFiles []string) error {
	for _, migrationFile := range migrationFiles {
		migration, err := os.ReadFile(migrationFile)
		if err != nil {
			return err
		}
		_, err = d.db.Exec(string(migration))
		if err != nil {
			return err
		}

		logger.Info("Applied migration from file %s", zap.String("file", migrationFile))
	}
	return nil
}
