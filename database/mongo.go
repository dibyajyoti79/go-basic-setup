package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"pripgo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Create a new MongoDB client with the URI from the environment variable
	clientOptions := options.Client().ApplyURI(config.GetEnv("MONGODB_URI", "mongodb://localhost:27017"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	// Ping MongoDB to verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// Set the global DB variable
	DB = client.Database(config.GetEnv("DB_NAME", "pripgoDev"))
	fmt.Println("Database connected!")
}
