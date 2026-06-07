package stocknews

import (
	"github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if len(stocknews) == 0 {
		return false, nil
	}

	result := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(stocknews)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}
