package routers

import (
	"labqid/app/controllers"
	"labqid/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Book(app *fiber.App) {
	router := app.Group("/books")

	router.Get("/", middlewares.ExampleMiddleware, controllers.FetchAllBooks) // contoh menggunakan middleware
	router.Post("/", controllers.CreateBook)
}
