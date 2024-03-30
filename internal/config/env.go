package config

import "github.com/joho/godotenv"

// LoadDotEnv is a function that loads environment variables from a .env file.
//
// It uses the godotenv package to read the .env file and load the variables into the environment.
// If multiple file names are provided, it will read them in order and later files will overwrite variables from earlier ones.
//
// Parameters:
// filenames - The names of the .env files to load. If no names are provided, it will default to loading ".env".
//
// Returns:
// error - An error object that will be non-nil if an error occurred.
func LoadDotEnv(filenames ...string) error {
	return godotenv.Load(filenames...)
}
