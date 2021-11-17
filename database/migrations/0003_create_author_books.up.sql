CREATE TABLE IF NOT EXISTS author_books(
    author_id        BINARY (16) NOT NULL,
    book_id          BINARY (16) NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NULL,
    deleted_at       TIMESTAMP NULL
) ENGINE=InnoDB;
