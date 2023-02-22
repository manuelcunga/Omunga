package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	DSN := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	// dbPath := os.Getenv("./")
	// DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		panic("failed to connect database")

	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(entities.User{})

	log.Println("Connected to database...")

	return DB
}
