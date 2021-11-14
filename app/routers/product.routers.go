package routers

import (
	"fiber-gorm/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Product(app *fiber.App) {
	user := app.Group("/products")

	user.Get("/", middlewares.IsAuthenticated, func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "ok",
		})
	})
}
