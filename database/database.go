package database

import (
	"log"
	"os"

	"github.com/nr-wazid/fiber-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "host=localhost user=postgres password=11110111 dbname=bookdb port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected To Database Successfully")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migration")

	// TODO: Add Migration
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
