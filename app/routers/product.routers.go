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

	router.Get(
		"/:product_id",
		controllers.ProductDetail,
	)

	router.Post(
		"/",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.CreateProduct,
	)

	router.Post(
		"/upload_image/:product_id",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.UploadImage,
	)

	router.Patch(
		"/",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.UpdateProduct,
	)

	router.Patch(
		"/active_inactive/:product_id",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.ActiveInActiveProduct,
	)

	router.Patch(
		"/remove_image/:product_id",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.RemoveImage,
	)

	router.Patch(
		"/:product_id",
		middlewares.IsAuthenticated,
		middlewares.IsAdmin,
		controllers.UpdateProduct,
	)
}
