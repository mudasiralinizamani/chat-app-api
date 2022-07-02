package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func CheckIsPasswordValid(hashedPassword, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))

	var isValid bool = true
	var msg string = ""

	if err != nil {
		msg = "Password is incorrect"
		isValid = false
	}

	return isValid, msg
}
