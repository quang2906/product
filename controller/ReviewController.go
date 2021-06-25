package controller

import (
	"fmt"

	"example.com/product/model"
	repo "example.com/product/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}

func GetReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	review, err := repo.Reviews.FindReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(review)
}

func FindReviewById(c *fiber.Ctx) int {
	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(c.Status(400).SendString(err.Error()))
	}
	review, err := repo.Reviews.FindReviewById(int64(id))
	if err != nil {
		fmt.Println(c.Status(404).SendString(err.Error()))
	}
	return (int(review.Id))
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	productId := repo.Reviews.FindReviewById2(int64(id)).ProductId
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	repo.UpdateProductRating(repo.Products.FindProductById2(int64(productId)))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	reviewId := repo.Reviews.CreateNewReview(review)
	repo.UpdateProductRating(repo.Products.FindProductById2(review.ProductId))
	return c.SendString(fmt.Sprintf("New review is created successfully with id = %d", reviewId))

}

func UpdateReview(c *fiber.Ctx) error {
	updatedReview := new(model.Review)

	err := c.BodyParser(&updatedReview)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Reviews.UpdateReview(updatedReview)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	repo.UpdateProductRating(repo.Products.FindProductById2(updatedReview.ProductId))
	return c.SendString(fmt.Sprintf("Review with id = %d is successfully updated", updatedReview.Id))

}

func UpsertReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Reviews.Upsert(review)
	repo.UpdateProductRating(repo.Products.FindProductById2(review.ProductId))
	return c.SendString(fmt.Sprintf("Review with id = %d is successfully upserted", id))
}
