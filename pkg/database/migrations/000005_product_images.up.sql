CREATE TABLE IF NOT EXISTS product_images(
    id SERIAL PRIMARY KEY,
    image_url VARCHAR(255),
    product_id int,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);sss