package models

type Stocks struct {
	Id           uint   `gorm:"primarykey;autoIncrement"`
	Company_name string `gorm:"column:company_name"`
	Symbol       string `gorm:"column:symbol"`
}

func (s *Stocks) TableName() string {
	return "stocks"
}
