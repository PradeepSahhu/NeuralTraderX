package repository

type Temp struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"column:firstname"`
	LastName  string `gorm:"column:lastname"`
	CreatedAt string `gorm:"column:created_at"`
}
