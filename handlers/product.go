package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/souzagmu/go-admin-web-react/types"
)

func (g StoreDB) GetAllProducts(c *fiber.Ctx) error {

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	var (
		limit    = 5
		total    int64
		offset   = (page - 1) * 5
		products = []types.Product{}
	)

	g.db.Offset(offset).Limit(limit).Find(&products)
	g.db.Model(types.Product{}).Count(&total)

	lastPageInt := (int(total) / limit)

	return c.JSON(fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPageInt,
		},
	})
}

func (g StoreDB) GetProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	product := types.Product{Id: uint(convId)}

	g.db.Find(&product)
	return c.JSON(product)
}

func (g StoreDB) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	product := types.Product{Id: uint(convId)}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	g.db.Model(&product).Updates(product)
	return c.JSON(product)
}

func (g StoreDB) CreateProduct(c *fiber.Ctx) error {
	product := types.Product{}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	if responseValidate := product.ValidateProduct(); len(responseValidate) > 0 {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON(responseValidate)
	}

	g.db.Create(&product)

	return c.JSON(product)

}

func (g StoreDB) DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	product := types.Product{Id: uint(convId)}

	g.db.Delete(&product)

	return c.JSON(product)
}
