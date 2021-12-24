package routers

import (
	"labqid/app/controllers"
	"labqid/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Category(app *fiber.App) {
	router := app.Group("/categories")

	router.Get(
		"/",
		controllers.FetchAllCategories,
	)
	router.Post(
		"/",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.CreateCategory,
	)
}
