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

func (r *TempRepository) CreateTemp(firstName, lastName, createdAt string) error {
	temp := models.Temp{
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: createdAt,
	}

	return r.db.Create(temp).Error
}
