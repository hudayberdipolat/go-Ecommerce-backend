CREATE TYPE adminStatus as ENUM('ACTIVE', 'PASSIVE');
CREATE TABLE IF NOT EXISTS admins(
    id SERIAL PRIMARY KEY, 
    username VARCHAR(100) NOT NULL UNIQUE,
    full_name        VARCHAR(255) NOT NULL,
    phone_number  VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    admin_image_url VARCHAR(255),
    admin_status adminStatus  DEFAULT 'ACTIVE',
    password   VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);

