CREATE TABLE IF NOT EXISTS migrations (
    id SERIAL UNIQUE,
    name TEXT UNIQUE NOT NULL,
    checksum TEXT NOT NULL CHECK (length(checksum) = 64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
