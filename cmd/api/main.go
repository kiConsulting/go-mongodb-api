package main

import (
	"log"
	"net/http"

	"go-mongodb-api/internal/config"
	"go-mongodb-api/internal/db"
	"go-mongodb-api/internal/handlers"
)

func main() {
	cfg := config.Load()

	client, err := db.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	} 
	defer client.Disconnect(nil)
	
	userHandler := handlers.NewUserHandler(client)

	http.HandleFunc("/users", userHandler.GetUsers)
	http.HandleFunc("/users/{id}/{age}", userHandler.GetUerIDAndAge)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}