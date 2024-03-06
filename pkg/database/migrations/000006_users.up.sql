CREATE TYPE userStatus as ENUM ('ACTIVE','PASSIVE');
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(100),
    phone_number VARCHAR(30) UNIQUE,
    email VARCHAR(100) UNIQUE,
    user_status userStatus DEFAULT 'ACTIVE',
    password VARCHAR(255),
    created_at TIMESTAMP, 
    updated_at TIMESTAMP
);