CREATE TABLE probability_links (
    id SERIAL PRIMARY KEY,
    url_a TEXT NOT NULL,
    url_b TEXT NOT NULL,
    probability_a FLOAT NOT NULL CHECK (probability_a >= 0 AND probability_a <= 1),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    short_code TEXT NOT NULL UNIQUE
);