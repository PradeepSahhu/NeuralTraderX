package models

type Temp struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"column:firstname"`
	LastName  string `gorm:"column:lastname"`
	CreatedAt string `gorm:"column:created_at"`
}
