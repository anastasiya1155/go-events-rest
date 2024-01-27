package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "27go9g32g@#@14cq@@}"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 2),
		},
	)

	return token.SignedString([]byte(secret))
}
