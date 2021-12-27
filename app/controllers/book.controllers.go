package controllers

import (
	"fmt"
	"labqid/app/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UploadBook(c *fiber.Ctx) error {

	// Get first file from form field "document":
	file, err := c.FormFile("book_images")

	// Check for errors:
	if err == nil {
		c.SaveFile(file, fmt.Sprintf("./tmp/uploads_relative/%s", file.Filename))
	} else {
		fmt.Println(err)
	}
	return c.Status(200).SendString("")

}

func FetchAllBooks(c *fiber.Ctx) error {
	result, _ := models.FethAllBooks()
	return c.Status(result.Status).JSON(result)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book

	book.Author = c.FormValue("author")
	book.Name = c.FormValue("name")
	book.NoISBN = c.FormValue("no_isbn")

	if book.Author == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "author is required"})
	}
	if book.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}
	if book.NoISBN == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "no_isbn is required"})
	}

	result, _ := models.CreateABook(&book)
	return c.Status(result.Status).JSON(result)
}
