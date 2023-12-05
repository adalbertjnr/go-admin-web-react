package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/souzagmu/go-admin-web-react/types"
)

func (g StoreDB) User(c *fiber.Ctx) error {

	var user types.User

	email := c.Locals("email").(string)

	g.db.Where("email = ?", email).First(&user)

	return c.JSON(user)
}
