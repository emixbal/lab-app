package routers

import (
	"labqid/app/controllers"
	"labqid/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Category(app *fiber.App) {
	router := app.Group("/categories")

	router.Get("/", middlewares.ExampleMiddleware, controllers.FetchAllCategories) // contoh menggunakan middleware
	router.Post("/", controllers.CreateCategory)
}
