



CREATE TABLE stocknews(
    id SERIAL PRIMARY KEY,
    stock_id INTEGER NOT NULL REFERENCES stocks(id),
    title TEXT NOT NULL ,
    description TEXT,
    link TEXT NOT NULL UNIQUE,
    source VARCHAR(200),
    is_processed BOOLEAN NOT NULL DEFAULT FALSE,
    published_at TIMESTAMPTZ,
    stored_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);