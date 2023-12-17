package handlers

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func (g StoreDB) Upload(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "err"})
	}

	files := form.File["image"]
	fileName := ""
	for _, file := range files {
		fileName = file.Filename
		filePath := filepath.Join("./uploads", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return c.JSON(map[string]string{"msg": "err"})
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + fileName,
	})
}
