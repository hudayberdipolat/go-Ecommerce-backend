CREATE TYPE status as ENUM ('ACTIVE', 'PASSIVE', 'DRAFT'); 
CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    category_name_tk VARCHAR(255) UNIQUE NOT NULL,
    category_name_ru VARCHAR(255) UNIQUE NOT NULL,
    category_name_en VARCHAR(255) UNIQUE NOT NULL,
    category_slug VARCHAR(255) NOT NULL,
    category_image_url VARCHAR(255) NOT NULL,
    category_status status DEFAULT 'DRAFT',
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);