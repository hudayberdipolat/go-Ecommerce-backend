CREATE TYPE admin_status_enum AS ENUM('ACTIVE','PASSIVE');
CREATE TYPE admin_role_enum AS ENUM('super_admin','admin');

CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY, 
    username VARCHAR(100) NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    admin_image_url VARCHAR(255),
    admin_status admin_status_enum DEFAULT 'ACTIVE',
    admin_role admin_role_enum DEFAULT 'admin',
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

