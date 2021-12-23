package routers

import (
	"labqid/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	user := app.Group("/users")

	user.Post("/auth/login", controllers.UserLogin)
	user.Post("/auth/register", controllers.UserRegister)
}
