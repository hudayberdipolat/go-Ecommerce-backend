CREATE TABLE IF NOT EXISTS cart(
    id SERIAL PRIMARY KEY,
    user_id int,
    product_id int,
    product_mukdar int,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);