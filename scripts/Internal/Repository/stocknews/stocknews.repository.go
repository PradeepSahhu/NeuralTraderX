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

func (r *StockNewsRepository) GetNewsByStockId(stockId uint) ([]models.StockNews, error) {

	var stockNews []models.StockNews
	result := r.db.Where("stock_id = ?", stockId).Find(&stockNews)

	if result.Error != nil {
		return nil, result.Error
	}

	return stockNews, nil
}

func (r *StockNewsRepository) GetBatchNews() ([]models.StockNews, error) {

	var stockNews []models.StockNews
	result := r.db.Find(&stockNews)

	if result.Error != nil {
		return nil, result.Error
	}

	return stockNews, nil
}
