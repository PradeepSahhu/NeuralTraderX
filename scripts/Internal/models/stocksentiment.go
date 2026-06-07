package models

import "time"

type StockSentiment struct {
	Id             uint      `gorm:"primaryKey;autoIncrement"`
	StockId        uint      `gorm:"column:stock_id"`
	NewsId         uint      `gorm:"column:news_id"`
	Sentiment      string    `gorm:"column:sentiment"`
	SentimentScore float32   `gorm:"column:sentiment_score"`
	Description    string    `gorm:"column:description"`
	Source         string    `gorm:"column:source"`
	UpdatedOn      time.Time `gorm:"column:updated_on"`
}
