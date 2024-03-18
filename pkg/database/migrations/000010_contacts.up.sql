CREATE TABLE IF NOT EXISTS contacts(
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(30),
    youtube_account   VARCHAR(255),
    instagram_account   VARCHAR(255),
    tiktok_account   VARCHAR(255),
    imo_account   VARCHAR(255),
    address   TEXT,
    play_market VARCHAR(255),
    app_store VARCHAR(255),
    qr_code  VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);