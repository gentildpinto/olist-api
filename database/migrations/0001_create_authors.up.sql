CREATE TABLE IF NOT EXISTS `authors` (
    id VARCHAR (256) PRIMARY KEY,
    name VARCHAR (256) UNIQUE NOT NULL,
    created_at DATETIME (3) NULL,
    updated_at DATETIME (3) NULL,
    deleted_at DATETIME (3) NULL
);