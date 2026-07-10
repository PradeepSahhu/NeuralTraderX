

CREATE TABLE tradesignals(
    id serial PRIMARY KEY,
    stock_id INTEGER NOT NULL REFERENCES stocks(id),
    sentiment_id BIGINT NOT NULL REFERENCES stocksentiment(id),
    trade_date TIMESTAMP NOT NULL,
    trade_percentage DECIMAL(5,2),
    llm_score DECIMAL(5,2),
    signal VARCHAR(20),
    reason TEXT
);