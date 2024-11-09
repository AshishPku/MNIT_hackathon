package middleware

import (
	"easyjobBackend/utils"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
)

func JwtAuth(c *fiber.Ctx) error {
	tokenStr := c.Cookies("jwt")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unauthorized")
		}
		return []byte(utils.JWTSECRET), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("unauthorized")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["id"] != nil {
			c.Locals("id", claims["id"])
			//id is company id
		}
		c.Next()
		return nil
	}
	return fmt.Errorf("unauthorized")

}
