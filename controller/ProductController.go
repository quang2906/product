package controller

import (
	"fmt"

	"example.com/product/model"
	repo "example.com/product/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	return c.JSON(repo.Products.GetAllProducts())
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product, err := repo.Products.FindProductById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(product)
}

func FindProductById(c *fiber.Ctx) int {
	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(c.Status(400).SendString(err.Error()))
	}
	product, err := repo.Products.FindProductById(int64(id))
	if err != nil {
		fmt.Println(c.Status(404).SendString(err.Error()))
	}
	return (int(product.Id))
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Products.DeleteProductById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	err := c.BodyParser(&product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	productId := repo.Products.CreateNewProduct(product)
	return c.SendString(fmt.Sprintf("New product is created successfully with id = %d", productId))

}

func UpdateProduct(c *fiber.Ctx) error {
	updatedProduct := new(model.Product)

	err := c.BodyParser(&updatedProduct)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Products.UpdateProduct(updatedProduct)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated", updatedProduct.Id))

}

func UpsertProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	err := c.BodyParser(&product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Products.Upsert(product)
	return c.SendString(fmt.Sprintf("Product with id = %d is successfully upserted", id))
}
