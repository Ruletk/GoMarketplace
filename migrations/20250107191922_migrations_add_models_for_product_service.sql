-- +goose Up

SELECT 'up SQL query';

CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    parent_id BIGINT
);

CREATE TABLE inventories (
     inventory_id SERIAL PRIMARY KEY,
     quantity BIGINT NOT NULL
);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    company_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category_id BIGINT NOT NULL,
    inventory_id BIGINT NOT NULL,
    discount_id BIGINT
);

-- +goose Down
-- +goose StatementBegin

DROP TABLE products;
DROP TABLE inventories;
DROP TABLE categories;

-- +goose StatementEnd