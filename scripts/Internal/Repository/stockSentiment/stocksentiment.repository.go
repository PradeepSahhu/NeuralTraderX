package repository

import (
	"github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/models"
	"gorm.io/gorm"
)

type StockSentimentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *StockSentimentRepository {
	return &StockSentimentRepository{
		db: db,
	}
}

func (r *StockSentimentRepository) GetAllStockSentiment() ([]models.StockSentiment, error) {

	var stockSentiment []models.StockSentiment
	result := r.db.Find(&stockSentiment)

	if result.Error != nil {
		return nil, result.Error
	}

	return stockSentiment, nil
}

func (r *StockSentimentRepository) GetAllStockSentimentToday() ([]models.StockSentiment, error) {

	var stockSentiment []models.StockSentiment
	result := r.db.Where("DATE(updated_on) = CURRENT_DATE").Find(&stockSentiment)

	if result.Error != nil {
		return nil, result.Error
	}

	return stockSentiment, nil

}
