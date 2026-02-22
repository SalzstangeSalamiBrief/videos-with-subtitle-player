CREATE TABLE migrations (
    name TEXT UNIQUE NOT NULL,
    checksum TEXT NOT NULL CHECK (length(checksum) = 64),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
