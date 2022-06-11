package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FullName     *string            `json:"full_name"`
	Email        *string            `json:"email"`
	Password     *string            `json:"password"`
	Role         *string            `json:"role"`
	ProfilePic   *string            `json:"profile_pic"`
	Friends      *int               `json:"friends"`
	Dms          *bool              `json:"dms"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       string             `json:"user_id"`
}
