package routers

import (
	"labqid/app/controllers"
	"labqid/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Product(app *fiber.App) {
	router := app.Group("/products")

	router.Get(
		"/",
		controllers.FetchAllProducts,
	)

	router.Post(
		"/",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.CreateProduct,
	)

	router.Patch(
		"/active_inactive/:product_id",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.ActiveInActiveProduct,
	)
}
