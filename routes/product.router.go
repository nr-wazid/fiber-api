package routes

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nr-wazid/fiber-api/database"
	"github.com/nr-wazid/fiber-api/models"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("not valid")
	}
	return nil
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	fmt.Println(products)

	database.Database.Db.Find(&products)
	fmt.Println(products)

	responseProducts := []Product{}

	for _, product := range products {
		responseProduct := CreateResponseProduct(product)

		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

// GET ONE PRODUCT
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	// 1. Make a slice
	products := []models.Product{}
	// 2. Find data from db
	database.Database.Db.Find(&products)
	// 3. Single products slice
	responseProducts := []Product{}
	// 4. Loop through products
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}
