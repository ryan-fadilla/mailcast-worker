package database

import (
	"log"
	"mailcast-worker/configuration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

// ConnectGORM initializes a GORM DB connection
func ConnectGORM(cfg configuration.Config) {
	// cfg := configuration.LoadConfig() // Load DB config from YAML

	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("🚀 Application running with GORM connected successfully!")
	DB = db
}
