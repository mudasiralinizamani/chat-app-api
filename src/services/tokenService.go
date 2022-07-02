package services

import (
	"chat_me_api/src/data"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateTokens(token, refreshToken, userId string) {
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", token})
	updateObj = append(updateObj, bson.E{"refresh_token", refreshToken})

	UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", UpdatedAt})

	upsert := false
	filter := bson.M{"user_id": userId}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := data.UserCollection.UpdateOne(ctx, filter, bson.D{
		{"$set", updateObj},
	}, &opt)

	if err != nil {
		return
	}
}
