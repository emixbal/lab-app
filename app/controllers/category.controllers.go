package controllers

import (
	"labqid/app/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FetchAllCategories(c *fiber.Ctx) error {
	result, _ := models.FethAllCategories()
	return c.Status(result.Status).JSON(result)
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	category.Name = c.FormValue("name")
	category.Description = c.FormValue("description")

	if category.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}

	result, _ := models.CreateACategory(&category)
	return c.Status(result.Status).JSON(result)
}
