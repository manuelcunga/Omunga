package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	DSN := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		panic("failed to connect database")

	}

	fmt.Println("Connection Opened to Database")
	DB.Logger.LogMode(logger.Silent)
	DB.AutoMigrate(entities.User{})

	log.Println("Connected to database...")

	return DB
}
