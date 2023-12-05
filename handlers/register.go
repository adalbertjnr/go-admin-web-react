package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/souzagmu/go-admin-web-react/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StoreDB struct {
	db *gorm.DB
}

func NewStoreDB(db *gorm.DB) *StoreDB {
	return &StoreDB{db: db}
}

func (g StoreDB) Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	fmt.Println(data)

	if data["pass_confirm"] != data["password"] || data["pass_confirm"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "passwords dont match"})
	}

	userParams := types.NewUser{
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
		Password:  data["password"],
	}

	user, err := types.NewUserFn(userParams)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "error"})
	}

	g.db.Create(&user)

	return c.JSON(user)
}

func (g StoreDB) Login(c *fiber.Ctx) error {
	var login types.UserLoginParams

	if err := c.BodyParser(&login); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "incorrect login type"})
	}

	valid := types.IsEmailValid(login.Email)

	if !valid {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "not valid email"})
	}

	var user types.User

	g.db.Where("email = ?", login.Email).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{"message": "not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.EncPassword), []byte(login.Password)); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "invalid password"})
	}

	token, err := createToken(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}

	cookie := &fiber.Cookie{
		Name:     "jwt",
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

	tokenString, err := token.SignedString([]byte("S5AD456SA45C"))
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("error signing the token")
	}

	return tokenString, nil
}
