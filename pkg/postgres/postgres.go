package postgres

import (
	"crud/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPQ() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=test password=test123 dbname=testdb port=5431 sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Usermodel{})
	return db

}
