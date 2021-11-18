CREATE TABLE IF NOT EXISTS books(
    id               BINARY (16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
    name             VARCHAR (50) UNIQUE NOT NULL,
    edition          VARCHAR (10) NOT NULL,
    publication_year INTEGER NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NULL,
    deleted_at       TIMESTAMP NULL
) ENGINE=InnoDB;
