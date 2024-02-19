CREATE TYPE  categoryStatus AS ENUM ('active','passive', 'draft');
CREATE TABLE if NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    category_name_tk varchar(500) NOT NULL UNIQUE,
    category_name_ru varchar(500) NOT NULL UNIQUE,
    category_slug varchar(500) NOT NULL UNIQUE,
    category_status categoryStatus DEFAULT 'draft'
);