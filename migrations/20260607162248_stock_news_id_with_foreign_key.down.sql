


ALTER TABLE stocksentiment
DROP CONSTRAINT IF EXISTS fk_stocksentiment_news_id;

ALTER TABLE stocksentiment
DROP COLUMN IF EXISTS news_id;