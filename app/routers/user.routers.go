package routers

import (
	"fiber-gorm/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	user := app.Group("/users")

	user.Get("/auth/login", controllers.UserLogin)
	user.Get("/auth/register", controllers.UserRegister)
}
