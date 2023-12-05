package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/souzagmu/go-admin-web-react/types"
	"gorm.io/gorm"
)

const JWTTOKEN = "a5sd4s56a4c654"

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
