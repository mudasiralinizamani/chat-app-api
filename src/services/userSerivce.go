package services

import (
	"chat_me_api/src/data"
	"chat_me_api/src/dtos"
	"chat_me_api/src/helpers"
	"chat_me_api/src/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(dto dtos.SignupDto, hashedPassword string) (*mongo.InsertOneResult, error) {
	var user models.User

	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Full_name = dto.FullName
	user.Email = dto.Email
	user.Role = "USER"
	user.Profile_pic = dto.ProfilePic
	user.Password = &hashedPassword
	user.Dms = true
	user.Friends = 0

	token, refreshToken, err := helpers.GenerateTokens(*user.Email, *user.Full_name, user.Role, user.User_id)

	if err != nil {
		return &mongo.InsertOneResult{}, errors.New("error occurred while generating token")
	}

	user.Token = &token
	user.Refresh_token = &refreshToken

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertionNumber, insertErr := data.UserCollection.InsertOne(ctx, user)

	if insertErr != nil {
		return &mongo.InsertOneResult{}, errors.New("error occurred while inserting the user")
	}

	return insertionNumber, nil
}
