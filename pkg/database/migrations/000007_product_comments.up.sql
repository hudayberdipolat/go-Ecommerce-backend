CREATE TABLE IF NOT EXISTS product_comments(
    id SERIAL PRIMARY KEY,
    product_comment text NOT NULL,
    user_id int NOT NULL,
    product_id int NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);