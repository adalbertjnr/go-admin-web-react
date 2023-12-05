package handlers

import "github.com/gofiber/fiber/v2"

func (g StoreDB) Logout(c *fiber.Ctx) error {
	c.ClearCookie("hunter-123-token")
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "successful logout"})
}
