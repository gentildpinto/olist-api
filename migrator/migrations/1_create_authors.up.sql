CREATE TABLE IF NOT EXISTS "authors" (
    "id" varchar(256),
    "name" varchar(256) UNIQUE,
    "created_at" datetime(3) NULL,
    "updated_at" datetime(3) NULL,
    "deleted_at" datetime(3) NULL,
    PRIMARY KEY ("id"),
    INDEX idx_authors_deleted_at ("deleted_at")
);