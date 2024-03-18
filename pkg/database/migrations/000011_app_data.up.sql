CREATE TABLE IF NOT EXISTS app_data(
        id SERIAL PRIMARY KEY,
        app_logo VARCHAR(255) NOT NULL,
        app_store VARCHAR(255),
        play_market VARCHAR(255),
        qr_code VARCHAR(255)
);

