package postgres

import (
	"go-self-payroll-system/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Position{})
	if err != nil {
		log.Fatal(err)
	}

	return db

}
