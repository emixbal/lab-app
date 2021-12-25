package controllers

import (
	"fmt"
	"labqid/app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FetchAllProducts(c *fiber.Ctx) error {
	result, _ := models.FethAllProducts()
	return c.Status(result.Status).JSON(result)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")
	price, _ := strconv.Atoi(c.FormValue("price"))
	product.Price = int16(price)

	file, err_file := c.FormFile("image")

	if err_file != nil {
		log.Println(err_file)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "Something went wrong!"})
	}

	// 👷 Save file inside uploads folder under current working directory:
	c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	// 👷 Save file using a relative path:
	c.SaveFile(file, fmt.Sprintf("./tmp/uploads_relative/%s", file.Filename))

	if product.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}

	result, _ := models.CreateAProduct(&product)
	return c.Status(result.Status).JSON(result)
}
