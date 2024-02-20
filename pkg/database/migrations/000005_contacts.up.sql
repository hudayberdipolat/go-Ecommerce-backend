CREATE TABLE IF NOT EXISTS contacts(
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(100),
    email VARCHAR(255),
    address VARCHAR(1000)
);