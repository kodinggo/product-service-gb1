
-- +migrate Up

CREATE TABLE categories (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


-- +migrate Down
DROP TABLE IF EXISTS categories;
