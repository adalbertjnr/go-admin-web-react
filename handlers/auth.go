package handlers

import (
	"fmt"
	"os"
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

	token, err := createToken(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}

	cookie := &fiber.Cookie{
		Name:     "hunter-123-token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(cookie)

	return c.JSON(fiber.Map{"message": "connected"})

}

func createToken(user types.User) (string, error) {

	var (
		idString = strconv.Itoa(int(user.Id))
		now      = time.Now()
		expires  = now.Add(time.Hour * 24).Unix()
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      idString,
		"expires": expires,
		"email":   user.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTTOKEN")))
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("error signing the token")
	}

	return tokenString, nil
}

func (g StoreDB) Logout(c *fiber.Ctx) error {
	c.ClearCookie("hunter-123-token")
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "successful logout"})
}

func (g StoreDB) UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "error"})
	}

	id := c.Locals("id")
	idInt, err := strconv.Atoi(id.(string))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "conversion error"})
	}

	user := types.User{
		Id:        uint(idInt),
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
	}

	g.db.Model(&user).Updates(&user)

	return c.JSON(user)
}

func (g StoreDB) UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "error"})
	}

	if data["pass_confirm"] != data["password"] || data["pass_confirm"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "passwords dont match"})
	}

	newHashedPass, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "bcrypt error"})
	}

	data["encPassword"] = string(newHashedPass)

	id := c.Locals("id")
	idInt, err := strconv.Atoi(id.(string))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "conversion error"})
	}

	user := types.User{
		Id: uint(idInt),
	}

	g.db.Model(&user).Update("enc_password", data["encPassword"])

	return c.JSON(user)
}
