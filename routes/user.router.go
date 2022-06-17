package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nr-wazid/fiber-api/database"
	"github.com/nr-wazid/fiber-api/models"
)

type UserSerializer struct {
	// Serializer
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) UserSerializer {
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("not valid")
	}
	return nil
}

func CerateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// GET ALL USERS
func GetUsers(c *fiber.Ctx) error {
	// 1. Make a slice
	users := []models.User{}
	// 2. Find data from db
	database.Database.Db.Find(&users)
	// 3. Single user slice
	responseUsers := []UserSerializer{}
	// 4. Loop through users
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

// GET ONE USERS
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("id is an integer")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	userData := CreateResponseUser(user)
	return c.Status(200).JSON(userData)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	// 1. Make a slice
	users := []models.User{}
	// 2. Find data from db
	database.Database.Db.Find(&users)
	// 3. Single user slice
	responseUsers := []UserSerializer{}
	// 4. Loop through users
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}
