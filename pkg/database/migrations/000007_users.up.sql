CREATE TYPE userStatus as ENUM ('active', 'passive');
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone_number VARCHAR(50) UNIQUE  NOT NULL,
    user_status userStatus DEFAULT 'active',
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL

);
