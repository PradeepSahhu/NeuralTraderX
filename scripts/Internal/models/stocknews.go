package models

import "time"

type StockNews struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"`
	StockId     uint      `gorm:"column:stock_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Link        string    `gorm:"column:link"`
	Source      string    `gorm:"column:source"`
	IsProcessed bool      `gorm:"column:is_processed"`
	PublishedAt time.Time `gorm:"column:published_at"`
	StoredAt    time.Time `gorm:"column:stored_at"`
}

func (s *StockNews) TableName() string {
	return "stocknews"
}
