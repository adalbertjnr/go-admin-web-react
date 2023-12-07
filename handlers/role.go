package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/souzagmu/go-admin-web-react/types"
)

func (g StoreDB) GetAllRoles(c *fiber.Ctx) error {
	role := make([]types.Role, 0)
	g.db.Find(&role)
	return c.JSON(role)
}

func (g StoreDB) GetRole(c *fiber.Ctx) error {

	var (
		id   = c.Params("id")
		role = types.Role{}
	)

	g.db.Where("id = ?", id).First(&role)
	return c.JSON(role)
}

func (g StoreDB) UpdateRole(c *fiber.Ctx) error {

	var (
		id        = c.Params("id")
		convId, _ = strconv.Atoi(id)
		role      = types.Role{Id: uint(convId)}
	)

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	g.db.Model(&role).Updates(role)
	return c.JSON(role)
}

func (g StoreDB) CreateRole(c *fiber.Ctx) error {
	role := types.Role{}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	g.db.Create(&role)

	return c.JSON(role)

}

func (g StoreDB) DeleteRole(c *fiber.Ctx) error {

	var (
		id        = c.Params("id")
		convId, _ = strconv.Atoi(id)
		role      = types.Role{Id: uint(convId)}
	)

	g.db.Delete(&role)
	msg := fmt.Sprintf("role %s deleted", role.Name)
	return c.JSON(fiber.Map{"message": msg})
}
