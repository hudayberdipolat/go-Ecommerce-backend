CREATE TYPE adminRole as ENUM ('super_admin', 'admin');
CREATE TYPE adminStatus as ENUM ('active', 'passive');
CREATE TABLE IF NOT EXISTS admins(
    id SERIAL PRIMARY KEY,
    firstname   VARCHAR(100) NOT NULL, 
    surname VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) UNIQUE,
    admin_role adminRole DEFAULT 'admin',
    admin_status adminStatus DEFAULT 'active',
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);