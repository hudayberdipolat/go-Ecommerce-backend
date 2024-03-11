CREATE TYPE sliderStatus AS ENUM('ACTIVE', 'PASSIVE');
CREATE TABLE IF NOT EXISTS sliders(
    id SERIAL PRIMARY KEY,
    slider_image_url VARCHAR(255),
    slider_status sliderStatus DEFAULT 'PASSIVE',
    created_at TIMESTAMP, 
    updated_at TIMESTAMP
);

