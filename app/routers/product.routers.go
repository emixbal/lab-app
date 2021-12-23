package routers

import (
	"labqid/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Product(app *fiber.App) {
	router := app.Group("/products")

	router.Get("/", middlewares.IsAuthenticated, func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "ok",
		})
	})
}
