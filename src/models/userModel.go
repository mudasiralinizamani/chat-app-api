package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Full_name     *string            `json:"full_name"`
	Email         *string            `json:"email"`
	Password      *string            `json:"password"`
	Role          string             `json:"role"`
	Profile_pic   *string            `json:"profile_pic"`
	Friends       int                `json:"friends"`
	Dms           bool               `json:"dms"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
