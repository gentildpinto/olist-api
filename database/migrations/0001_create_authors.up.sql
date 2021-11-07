CREATE TABLE IF NOT EXISTS authors(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR (50) UNIQUE NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NULL,
    deleted_at      TIMESTAMP NULL
);
