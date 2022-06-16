package controllers

import (
	"fmt"
	"go-auth/database"
	"go-auth/models"
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

func Populate(c *fiber.Ctx) error {
	for i := 0; i < 5; i++ {
		database.DB.Create(&models.Product{
			Title:       faker.Word(),
			Description: faker.Paragraph(),
			Image:       fmt.Sprintf("https://loremflickr.com/320/240?%s", faker.UUIDDigit()),
			Price:       rand.Intn(90) + 10,
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func All(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Find(&products)
	return c.JSON(products)
}

func Search(c *fiber.Ctx) error {
	var products []models.Product

	sql := "SELECT * FROM products"

	if s := c.Query("s"); s != "" {
		sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
	}

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY price %s", sql, sort)
	}

	database.DB.Raw(sql).Scan(&products)
	return c.JSON(products)

}
