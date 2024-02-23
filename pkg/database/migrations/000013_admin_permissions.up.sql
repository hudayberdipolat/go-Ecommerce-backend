CREATE TABLE IF NOT EXISTS admin_permissions(
    id SERIAL PRIMARY KEY,
    admin_id INT NOT NULL,
    permission_id INT NOT NULL,
    CONSTRAINT fk_admin FOREIGN KEY(admin_id) REFERENCES admins(id),
    CONSTRAINT fk_permission FOREIGN KEY(permission_id) REFERENCES permissions(id)
);