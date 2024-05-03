package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"jdnielss.dev/cats-social-app/model"
)

var JWT_KEY = os.Getenv("JWT_KEY")

func GenerateJWT(payload model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name": payload.Name,
			"exp":  time.Now().Add(time.Hour * 8).Unix(),
		})

	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
