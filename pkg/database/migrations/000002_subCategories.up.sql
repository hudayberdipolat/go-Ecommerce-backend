CREATE TYPE subCategoryStatus as ENUM ('ACTIVE', 'PASSIVE', 'DRAFT'); 
CREATE TABLE IF NOT EXISTS sub_categories(
    id SERIAL PRIMARY KEY,
    sub_category_name_tk VARCHAR(255) UNIQUE NOT NULL,
    sub_category_name_ru VARCHAR(255) UNIQUE NOT NULL,
    sub_category_name_en VARCHAR(255) UNIQUE NOT NULL,
    sub_category_slug VARCHAR(255) NOT NULL,
    sub_category_image_url VARCHAR(255) NOT NULL,
    sub_category_status subCategoryStatus DEFAULT 'DRAFT',
    category_id int,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id)
);   