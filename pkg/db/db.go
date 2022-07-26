package db

import (
	"log"

	"github.com/simple-crud-go-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// Composition: POSTGRES_USER : POSTGRES_PASSWORD @localhost:5432  POSTGRES_DB
	dbURL := "postgres://postgres:root@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	//Create table Books
	db.AutoMigrate(&models.Book{})

	return db
}
