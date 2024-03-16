CREATE TYPE brandStatus as ENUM ('ACTIVE', 'PASSIVE', 'DRAFT'); 
CREATE TABLE IF NOT EXISTS brands(
    id SERIAL PRIMARY KEY,
    brand_name_tk VARCHAR(100) UNIQUE NOT NULL,
    brand_name_ru VARCHAR(100) UNIQUE NOT NULL,
    brand_name_en VARCHAR(100) UNIQUE NOT NULL,
    brand_slug VARCHAR(200) NOT NULL, 
    brand_image_url VARCHAR(100) NOT NULL,
    brand_status brandStatus DEFAULT 'DRAFT',
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);