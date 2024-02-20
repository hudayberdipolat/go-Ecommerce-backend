CREATE TABLE IF NOT EXISTS about(
    id SERIAL PRIMARY KEY,
    about_desc TEXT NOT NULL, 
    location_map_html TEXT NOT NULL
);