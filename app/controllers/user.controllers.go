package controllers

import (
	"fiber-gorm/app/helpers"
	"fiber-gorm/app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UserRegister(c *fiber.Ctx) error {
	var user models.User

	txtPassword := c.FormValue("password")

	hashPassword, err := helpers.GeneratePassword(txtPassword)
	if err != nil {
		fmt.Println(err)
	}

	user.Email = c.FormValue("email")
	user.Name = c.FormValue("name")
	user.Password = hashPassword

	if user.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "email is required"})
	}
	if user.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}

	result, _ := models.UserRegister(&user)
	return c.Status(result.Status).JSON(result)

}

func UserLogin(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(map[string]string{"message": "User Logged in"})
}

func UserRefreshToken(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(map[string]string{"message": "User Refresh Token"})
}
