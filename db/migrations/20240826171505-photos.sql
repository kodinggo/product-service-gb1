
-- +migrate Up
CREATE TABLE photos (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS photos;