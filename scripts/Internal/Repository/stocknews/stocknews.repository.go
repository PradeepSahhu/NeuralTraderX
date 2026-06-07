package stocknews

import (
	"github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/models"
	"gorm.io/gorm"
)

type StockNewsRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *StockNewsRepository {
	return &StockNewsRepository{
		db: db,
	}
}

func (r *StockNewsRepository) InsertBatchStockNews(stocknews []*models.StockNews) (bool, error) {
	result := r.db.Create(stocknews)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}
