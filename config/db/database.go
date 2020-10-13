package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//GetDB is
func GetDB() (*gorm.DB, error) {

	dsn := "user=postgres password=2020$ dbname=ecommerce port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
