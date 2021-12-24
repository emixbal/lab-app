package routers

import (
	"labqid/app/controllers"
	"labqid/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Category(app *fiber.App) {
	router := app.Group("/categories")

	router.Get("/", middlewares.IsAuthenticated, middlewares.IsAdmin, controllers.FetchAllCategories) // contoh menggunakan middleware
	router.Post("/", middlewares.IsAdmin, controllers.CreateCategory)
}
