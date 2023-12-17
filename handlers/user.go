package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/souzagmu/go-admin-web-react/types"
)

func (g StoreDB) GetAllUsers(c *fiber.Ctx) error {

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	var (
		limit  = 5
		total  int64
		offset = (page - 1) * 5
		users  = []types.User{}
	)

	g.db.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	g.db.Model(types.User{}).Count(&total)

	lastPageInt := (int(total) / limit)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPageInt,
		},
	})
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
