CREATE TYPE productStatus AS ENUM ('active', 'passive', 'draft');
CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    product_name_tk varchar(500) NOT NULL , 
    product_name_ru varchar(500) NOT NULL , 
    product_slug varchar(500) NOT NULL , 
    product_desc_tk TEXT NOT NULL , 
    product_desc_ru TEXT NOT NULL , 
    product_status productStatus DEFAULT 'draft',
    main_image varchar(255) NOT NULL,
    price float NOT NULL,
    total_count int,
    galan_sany int,
    category_id int,
    brend_id int,
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id),
    CONSTRAINT fk_brend FOREIGN KEY(brend_id) REFERENCES brends(id)
);
