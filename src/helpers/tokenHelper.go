package helpers

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Email    string
	FullName string
	Role     string
	Uid      string
	jwt.StandardClaims
}

func GenerateTokens(email, fullName, role, uid string) (signedToken, signedRefreshToken string, err error) {
	err = nil

	claims := &SignedDetails{
		Email:    email,
		FullName: fullName,
		Role:     role,
		Uid:      uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(170)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		fmt.Println("error occurred while creating token")
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		fmt.Println("error occurred while creating refresh token")
		return
	}

	return token, refreshToken, err
}
