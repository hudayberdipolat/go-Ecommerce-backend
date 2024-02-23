CREATE TABLE IF NOT EXISTS admin_roles(
    id  SERIAL PRIMARY KEY,
    admin_id  INT NOT NULL,
    role_id INT NOT NULL,
    CONSTRAINT fk_admin FOREIGN KEY(admin_id) REFERENCES admins(id),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(id)

);