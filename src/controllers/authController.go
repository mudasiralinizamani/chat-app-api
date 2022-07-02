package controllers

import (
	"chat_me_api/src/data"
	"chat_me_api/src/dtos"
	"chat_me_api/src/helpers"
	"chat_me_api/src/models"
	"chat_me_api/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signupDto dtos.SignupDto

		if err := c.BindJSON(&signupDto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ServerError", "error": err.Error()})
			return
		}

		err := validate.Struct(signupDto)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "ValidationError", "error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		count, err := data.UserCollection.CountDocuments(ctx, bson.M{"email": signupDto.Email})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ServerError", "error": "something went wrong while counting user emails"})
			return
		} else if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"code": "EmailAlreadyExist", "error": "email address already exists"})
			return
		}

		hashedPassword := helpers.HashPassword(*signupDto.Password)

		insertNumber, err := services.CreateUser(signupDto, hashedPassword)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ServerError", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"succeeded": true, "insertNumber": insertNumber.InsertedID})
	}
}

func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto dtos.SigninDto

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ServerError", "error": err.Error()})
			return
		}

		validationErr := validate.Struct(dto)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "ValidationError", "error": validationErr.Error()})
			return
		}

		var user models.User

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := data.UserCollection.FindOne(ctx, bson.M{"email": dto.Email}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "EmailNotFound", "error": "email address does not exist"})
			return
		}

		isValid, msg := helpers.CheckIsPasswordValid(*user.Password, *dto.Password)

		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"code": "IncorrectPassword", "error": msg})
			return
		}

		token, refreshToken, err := helpers.GenerateTokens(*user.Email, *user.Full_name, user.Role, user.User_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ServerError", "error": " Error occurred while creating token and refresh token"})
			return
		}

		services.UpdateTokens(token, refreshToken, user.User_id)

		err = data.UserCollection.FindOne(ctx, bson.M{"user_id": user.User_id}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ServerError", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
