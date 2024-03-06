CREATE TABLE IF NOT EXISTS product_features(
    id SERIAL PRIMARY KEY,
    feature TEXT,
    product_id INT,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);