package db

import (
	"gorm.io/gorm"
)

//Product is
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//Migrate is
func Migrate() error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{Price: 100})

	return nil
}
