

ALTER TABLE stocksentiment
ADD COLUMN IF NOT EXISTS news_id INTEGER NOT NULL;


ALTER TABLE stocksentiment
ADD CONSTRAINT fk_stocksentiment_news_id
FOREIGN KEY (news_id) REFERENCES stocknews(id);