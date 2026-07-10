package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RegisterPostgresSql() *gorm.DB {

	dsn := "host=localhost user=pradeep password=password dbname=stocks port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")

	return db

}
