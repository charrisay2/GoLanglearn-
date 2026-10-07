package services

import (
	"context"
	configs "sliceServer/Configs"
	"sliceServer/models"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetCollection() *mongo.Collection {
	return configs.DB.Collection("users")
}
func GetAllUsers() []models.User {
	var users []models.User
	cursor, _ := GetCollection().Find(context.TODO(), bson.M{})
	cursor.All(context.TODO(), &users)
	return users
}

func Create(newUser models.User) models.User {
	newUser.ID = bson.NewObjectID()
	GetCollection().InsertOne(context.TODO(), newUser)
	return newUser
}

func GetUserById(id string) models.User {
	var user models.User
	objID, _ := bson.ObjectIDFromHex(id)
	GetCollection().FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	return user
}

func UserExistByID(id string) bool {
	user := GetUserById(id)
	if user.ID.IsZero() {
		return false
	}
	return true
}
