package services

import (
	"context"
	"pripgo/models"
	"pripgo/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserService(id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, utils.NewApiError(400, "Invalid User ID", nil)
	}
	var user models.User
	collection := utils.GetCollection("users")
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.NewApiError(404, "User not found", nil)
		}
		return nil, utils.NewApiError(500, "Internal Server Error", []string{err.Error()})
	}
	return &user, nil
}
