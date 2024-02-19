CREATE TYPE brendStatus  AS ENUM ('active','passive', 'draft');
CREATE TABLE IF NOT EXISTS brends(
    id SERIAL PRIMARY KEY,
    brend_name_tk VARCHAR(500) NOT NULL UNIQUE,
    brend_name_ru VARCHAR(500) NOT NULL UNIQUE,
    brend_slug VARCHAR(500) NOT NULL UNIQUE,
    brend_status brendStatus DEFAULT 'draft'
);