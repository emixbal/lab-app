package routers

import (
	"labqid/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	router := app.Group("/users")

	router.Post("/auth/login", controllers.UserLogin)
	router.Post("/auth/register", controllers.UserRegister)
}
