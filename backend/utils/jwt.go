package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const JWTSECRET string = "rentalsSecret123"

func CreateJwtfromID(id primitive.ObjectID) (*string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString([]byte(JWTSECRET))
	if err != nil {
		return nil, err
	}
	return &tokenstr, nil
}

func CreateCookieFromTokenString(tokenstr *string) *fiber.Cookie {
	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: *tokenstr,
	}
	return &cookie
}

func CreateCookieFromID(id primitive.ObjectID) (*fiber.Cookie, error) {
	tokenStr, err := CreateJwtfromID(id)
	if err != nil {
		return nil, err
	}
	cookie := CreateCookieFromTokenString(tokenStr)
	return cookie, nil
}
