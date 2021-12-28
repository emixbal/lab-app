package routers

import (
	"labqid/app/controllers"
	"labqid/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Chart(app *fiber.App) {
	router := app.Group("/charts")

	router.Post("/", middlewares.IsAuthenticated, controllers.NewChart)
	router.Get("/mychart", middlewares.IsAuthenticated, controllers.NewChart)
}
