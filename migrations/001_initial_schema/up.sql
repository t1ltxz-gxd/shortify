-- This statement creates a new table named 'urls' if it does not already exist.
-- The table has the following columns:
-- 'hash': This is the primary key of the table. It is a variable character string with a maximum length of 10.
-- 'original_url': This is a text column that cannot be null. It stores the original URL.
-- 'added_at': This is a timestamp column. Its default value is the current timestamp. It stores the time when the URL was added.
-- 'updated_at': This is a timestamp column. Its default value is the current timestamp. It stores the time when the URL was last updated.
CREATE TABLE IF NOT EXISTS urls (
    hash VARCHAR(10) PRIMARY KEY, -- The hash of the URL, serves as the primary key
    original_url TEXT NOT NULL, -- The original URL
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- The timestamp when the URL was added
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- The timestamp when the URL was last updated
);