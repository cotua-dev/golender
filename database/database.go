package database

import (
	"log"
	"os"

	"github.com/cotua-dev/golender/models"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var dbInstance DBInstance

func InitDB() {
	// Attempt to load the environment variables
	err := godotenv.Load()

	// Check for any errors produced by loading the .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Grab the environment variable values
	databaseUserName := os.Getenv("DB_USERNAME")
	databaseName := os.Getenv("DB_NAME")
	databasePassword := os.Getenv("DB_PASSWORD")

	// Create the database connection string
	connectionString := "host=golender_db user=" + databaseUserName + " password=" + databasePassword + " dbname=" + databaseName + " port=5432"

	log.Printf("Connection String: %q", connectionString)

	// Attempt to connect to the Postgres database
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// Log an error and exit if an error is present
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	// Log connection confirmation and set log mode to informational
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Run the migrations for all models
	log.Println("running migrations")
	db.AutoMigrate(&models.User{
		Email:    "test@test.com",
		Password: "test",
	})

	// Set the database instance variable
	dbInstance = DBInstance{
		DB: db,
	}
}
