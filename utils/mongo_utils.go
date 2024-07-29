package utils

import (
	"pripgo/database"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection returns a MongoDB collection from the database
func GetCollection(collectionName string) *mongo.Collection {
	return database.DB.Collection(collectionName)
}
