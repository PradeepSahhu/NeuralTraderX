package repository

import (
	"github.com/PradeepSahhu/NeuralTraderX/internal/models"
	"gorm.io/gorm"
)

type TempRepository struct {
	db *gorm.DB
}

func NewTempRepository(db *gorm.DB) *TempRepository {
	return &TempRepository{
		db: db,
	}
}

func (r *TempRepository) CreateTemp(firstName, lastName string) error {
	temp := models.Temp{
		FirstName: firstName,
		LastName:  lastName,
	}

	return r.db.Omit("CreatedAt").Create(&temp).Error
}
