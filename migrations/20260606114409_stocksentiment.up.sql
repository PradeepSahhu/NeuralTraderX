


CREATE TABLE stocksentiment (
    id BIGSERIAL PRIMARY KEY,
    stock_id INTEGER NOT NULL REFERENCES stocks(id),
    sentiment VARCHAR(200),
    sentiment_score NUMERIC(5,2),
    description TEXT,
    source VARCHAR(200),
    updated_on TIMESTAMP NOT NULL

);