CREATE TABLE IF NOT EXISTS users(
    id         UUID PRIMARY KEY,
    username   TEXT,
    password   TEXT,
    created_at TIMESTAMP,
)