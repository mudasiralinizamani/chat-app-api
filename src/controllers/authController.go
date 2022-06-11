package controllers

import (
	"chat_me_api/src/data"
	"chat_me_api/src/dtos"
	"chat_me_api/src/helpers"
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
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ServerError", "error": "Something went wrong"})
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
