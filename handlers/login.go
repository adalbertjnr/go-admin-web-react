package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/souzagmu/go-admin-web-react/types"
	"golang.org/x/crypto/bcrypt"
)

func (g StoreDB) Login(c *fiber.Ctx) error {
	var login types.UserLoginParams

	if err := c.BodyParser(&login); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "incorrect login type"})
	}

	valid := types.IsEmailValid(login.Email)

	if !valid {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "not valid email"})
	}

	var user types.User

	g.db.Where("email = ?", login.Email).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.EncPassword), []byte(login.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "invalid password"})
	}

	token, err := createToken(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}

	cookie := &fiber.Cookie{
		Name:     "hunter-123-token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	c.Cookie(cookie)

	return c.JSON(fiber.Map{"message": "connected"})

}

func createToken(id uint) (string, error) {

	idString := strconv.Itoa(int(id))
	now := time.Now()
	expires := now.Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      idString,
		"expires": expires,
	})

	tokenString, err := token.SignedString([]byte(JWTTOKEN))
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("error signing the token")
	}

	return tokenString, nil
}
