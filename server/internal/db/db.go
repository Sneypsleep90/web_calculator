package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"web_calculator/internal/calculationService"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatalf("error: could not connecr to data base %v", err)
	}
	if err := db.AutoMigrate(&calculationService.Calculation{}); err != nil {

		log.Fatalf("could not migrate: %v", err)
	}

	return db, nil

}
