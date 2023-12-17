package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthenticate(c *fiber.Ctx) error {

	fmt.Println("jwt authing...")

	token := c.Cookies("hunter-123-token")

	claims, err := parseToken(token)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "unauthorized"})
	}

	if float64(time.Now().Unix()) > claims["expires"].(float64) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "token expired"})
	}

	c.Locals("email", claims["email"])
	c.Locals("id", claims["id"])

	return c.Next()

}

func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTTOKEN")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
