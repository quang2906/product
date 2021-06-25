package main

import (
	"example.com/product/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Use(cors.New())

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	productRouter := app.Group("/api/product")
	userRouter := app.Group("/api/user")
	categoryRouter := app.Group("/api/category")
	reviewRouter := app.Group("api/review")
	routes.ConfigProductRouter(&productRouter)   //http://localhost:3000/api/product
	routes.ConfigUserRouter(&userRouter)         //http://localhost:3000/api/user
	routes.ConfigCategoryRouter(&categoryRouter) //http://localhost:3000/api/category
	routes.ConfigReviewRouter(&reviewRouter)     //http://localhost:3000/api/review

	app.Listen(":3000")
}
