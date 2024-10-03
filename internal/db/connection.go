package db

import (
	"exp/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbData := "host=localhost user=admin dbname=GE password=admin sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbData), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate([]domain.User{}, []domain.Post{})
	return db, nil
}
