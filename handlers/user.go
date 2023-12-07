package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/souzagmu/go-admin-web-react/types"
)

func (g StoreDB) GetAllUsers(c *fiber.Ctx) error {
	users := make([]types.User, 0)
	g.db.Find(&users)
	return c.JSON(users)
}

func (g StoreDB) GetUser(c *fiber.Ctx) error {

	var (
		id   = c.Params("id")
		user = types.User{}
	)

	g.db.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func (g StoreDB) UpdateUser(c *fiber.Ctx) error {

	var (
		id        = c.Params("id")
		convId, _ = strconv.Atoi(id)
		user      = types.User{Id: uint(convId)}
	)

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	g.db.Model(&user).Updates(user)
	return c.JSON(user)
}

func (g StoreDB) CreateUser(c *fiber.Ctx) error {
	user := types.NewUser{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPass("pass123")

	if err := user.ValidateLen(); len(err) > 0 {
		return c.JSON(err)
	}

	newUser, err := types.NewUserFn(user)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "error generating the new user"})
	}

	g.db.Create(&newUser)

	return c.JSON(newUser)

}

func (g StoreDB) DeleteUser(c *fiber.Ctx) error {

	var (
		id        = c.Params("id")
		convId, _ = strconv.Atoi(id)
		user      = types.User{Id: uint(convId)}
	)

	g.db.Delete(&user)
	msg := fmt.Sprintf("user %s deleted", user.FirstName)
	return c.JSON(fiber.Map{"message": msg})
}
