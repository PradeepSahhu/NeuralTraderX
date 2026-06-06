package repository

import (
	"github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/models"
	"gorm.io/gorm"
)

type StocksRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *StocksRepository {
	return &StocksRepository{
		db: db,
	}
}

func (r *StocksRepository) GetAllStocks() ([]models.Stocks, error) {

	var stock []models.Stocks
	result := r.db.Find(&stock)

	if result.Error != nil {
		return nil, result.Error
	}

	return stock, nil

}
