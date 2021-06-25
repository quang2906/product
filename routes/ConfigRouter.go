package routes

import (
	"example.com/product/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigProductRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllProduct)

	(*router).Get("/:id", controller.GetProductById)

	(*router).Delete("/:id", controller.DeleteProductById)

	(*router).Post("", controller.CreateProduct)

	(*router).Patch("", controller.UpdateProduct)

	(*router).Put("", controller.UpsertProduct)

}

func ConfigUserRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllUser)

	(*router).Get("/:id", controller.GetUserById)

	(*router).Delete("/:id", controller.DeleteUserById)

	(*router).Post("", controller.CreateUser)

	(*router).Patch("", controller.UpdateUser)

	(*router).Put("", controller.UpsertUser)

}

func ConfigCategoryRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllCategory)

	(*router).Get("/:id", controller.GetCategoryById)

	(*router).Delete("/:id", controller.DeleteCategoryById)

	(*router).Post("", controller.CreateCategory)

	(*router).Patch("", controller.UpdateCategory)

	(*router).Put("", controller.UpsertCategory)

}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllReview)

	(*router).Get("/:id", controller.GetReviewById)

	(*router).Delete("/:id", controller.DeleteReviewById)

	(*router).Post("", controller.CreateReview)

	(*router).Patch("", controller.UpdateReview)

	(*router).Put("", controller.UpsertReview)

}
