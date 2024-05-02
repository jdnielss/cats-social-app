package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"jdnielss.dev/cats-social-app/model"
)

var JWT_KEY = os.Getenv("JWT_KEY")

func GenerateJWT(payload model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(8 * time.Hour)
	claims["authorized"] = true
	claims["user"] = payload.Email

	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
