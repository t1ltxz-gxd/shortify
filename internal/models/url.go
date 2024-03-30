package models

import "time"

// URL is a struct that represents a URL in the application.
// It has four fields: Original, Hash, AddedAt, and UpdatedAt.
// Original is a string that holds the original URL.
// Hash is a string that holds the hashed version of the original URL.
// AddedAt is a time.Time value that holds the time when the URL was added to the application.
// UpdatedAt is a pointer to a time.Time value that holds the time when the URL was last updated in the application.
// If the URL has not been updated, UpdatedAt is nil.
type URL struct {
	Original  string     // The original URL
	Hash      string     // The hashed version of the original URL
	AddedAt   time.Time  // The time when the URL was added
	UpdatedAt *time.Time // The time when the URL was last updated, nil if not updated
}
