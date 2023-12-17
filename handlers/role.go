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

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	role := types.Role{Id: uint(id)}

	g.db.Preload("Permissions").Find(&role)
	return c.JSON(role)
}

func (g StoreDB) UpdateRole(c *fiber.Ctx) error {

	var (
		id      = c.Params("id")
		dtoRole fiber.Map
	)

	convId, err := strconv.Atoi(id)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	if err := c.BodyParser(&dtoRole); err != nil {
		return err
	}

	var (
		listFromJSON = dtoRole["permissions"].([]interface{})
		permissions  = make([]types.Permission, len(listFromJSON))
	)

	for i, permissonId := range listFromJSON {
		switch v := permissonId.(type) {
		case float64:
			permIdInt := uint(v)
			permissions[i] = types.Permission{
				Id: permIdInt,
			}
		case string:
			permIdInt, _ := strconv.Atoi(v)
			permissions[i] = types.Permission{
				Id: uint(permIdInt),
			}
		}
	}

	g.db.Table("role_permissions").Where("role_id", id).Delete(struct{}{})

	role := types.Role{
		Id:          uint(convId),
		Name:        dtoRole["name"].(string),
		Permissions: permissions,
	}

	g.db.Model(&role).Updates(role)

	return c.JSON(role)
}

func (g StoreDB) CreateRole(c *fiber.Ctx) error {
	var dtoRole map[string]interface{}

	if err := c.BodyParser(&dtoRole); err != nil {
		return err
	}

	var (
		listFromJSON = dtoRole["permissions"].([]interface{})
		permissions  = make([]types.Permission, len(listFromJSON))
	)

	for i, permissonId := range listFromJSON {
		switch v := permissonId.(type) {
		case float64:
			permIdInt := uint(v)
			permissions[i] = types.Permission{
				Id: permIdInt,
			}
		case string:
			permIdInt, _ := strconv.Atoi(v)
			permissions[i] = types.Permission{
				Id: uint(permIdInt),
			}
		}
	}

	role := types.Role{
		Name:        dtoRole["name"].(string),
		Permissions: permissions,
	}

	g.db.Create(&role)

	return c.JSON(fiber.Map{"message": role})

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
