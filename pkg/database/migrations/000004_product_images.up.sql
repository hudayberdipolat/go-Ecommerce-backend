CREATE TABLE IF NOT EXISTS product_images(
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    product_image varchar(255) NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);