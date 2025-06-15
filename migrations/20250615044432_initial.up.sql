CREATE TABLE url_mappings (
                                            id SERIAL PRIMARY KEY,
                                            long_url TEXT NOT NULL,
                                            short_code TEXT NOT NULL UNIQUE,
                                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX idx_short_code ON url_mappings (short_code);
