CREATE TABLE IF NOT EXISTS permissions(
    id SERIAL PRIMARY KEY,
    permission_name VARCHAR(255)  UNIQUE
);