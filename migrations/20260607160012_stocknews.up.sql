



CREATE TABLE stocknews(
    id SERIAL PRIMARY KEY,
    stock_id INTEGER NOT NULL REFERENCES stocks(id),
    title TEXT NOT NULL ,
    description TEXT NOT NULL,
    link TEXT NOT NULL UNIQUE,
    source VARCHAR(200),
    published_at TIMESTAMPZ,
    stored_at TIMESTAMPZ NOT NULL DEFAULT NOW()
);