package main

import (
	"log"
	"net/http"

	"github.com/jaibhavaya/dashboard-go/pkg/database"
	"github.com/jaibhavaya/dashboard-go/pkg/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	db, err := database.InitializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer db.Close()
	userRepo := models.NewUserRepository(db)

	r := setupRoutes(userRepo)

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
