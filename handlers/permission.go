package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/souzagmu/go-admin-web-react/types"
)

func (g StoreDB) GetAllPermissions(c *fiber.Ctx) error {
	permissions := make([]types.Permission, 0)
	g.db.Find(&permissions)
	return c.JSON(permissions)
}
