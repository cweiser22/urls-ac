CREATE TABLE accounts (
                          id SERIAL PRIMARY KEY,
                          email VARCHAR(100) NOT NULL UNIQUE,
                          password_hash TEXT NOT NULL,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);