package main

import (
	"log"
	"mailcast-worker/configuration"
	"mailcast-worker/database"
	"mailcast-worker/workers"
)

func main() {

	log.Println("🚀 Application running  ... ")

	// Load configuration
	cfg := configuration.LoadConfig()

	// Initialize database connection
	database.ConnectGORM(cfg)

	workers.WorkersServe()
	// workers.WorkersServeTest()
}
