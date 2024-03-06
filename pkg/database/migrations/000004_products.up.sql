CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    product_name_tk VARCHAR(500) NOT NULL,
    product_name_ru VARCHAR(500) NOT NULL,
    product_name_en VARCHAR(500) NOT NULL,
    product_slug VARCHAR(500) UNIQUE,
    product_short_desc_tk TEXT NOT NULL,
    product_short_desc_ru TEXT NOT NULL,
    product_short_desc_en TEXT NOT NULL,
    product_long_desc_tk TEXT NOT NULL,
    product_long_desc_ru TEXT NOT NULL,
    product_long_desc_en TEXT NOT NULL,
    product_main_image_url VARCHAR(255) NOT NULL,
    product_model VARCHAR(100) NOT NULL,
    product_feature TEXT,
    original_price NUMERIC(10,2) NOT NULL,
    discount_price NUMERIC(10,2),
    discount_time TIMESTAMP,
    total_count int  NOT NULL,
    rest_count int  NOT NULL,
    category_id int NOT NULL,
    sub_category_id int NOT NULL,
    brand_id int NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id),
    CONSTRAINT fk_sub_category FOREIGN KEY(sub_category_id) REFERENCES sub_categories(id),
    CONSTRAINT fk_brand FOREIGN KEY(brand_id) REFERENCES brands(id)
);