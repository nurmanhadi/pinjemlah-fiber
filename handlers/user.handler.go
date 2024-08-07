package handlers

import (
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func UserProfile(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	if err := databases.DB.First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserUpdateProfile(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	type inputUser struct {
		Name          string  `json:"name"`
		Email         string  `json:"email"`
		BirthDate     string  `json:"birth_date"`
		Address       string  `json:"address"`
		Ktp           string  `json:"ktp_number"`
		Income        float64 `json:"monthly_income"`
		AccountNumber string  `json:"account_number"`
	}
	var input inputUser
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON"})
	}
	var userr models.User
	if err := databases.DB.First(&userr, "user_id", user.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}

	if err := databases.DB.Model(&userr).Updates(map[string]interface{}{
		"name":           input.Name,
		"email":          input.Email,
		"birth_date":     input.BirthDate,
		"address":        input.Address,
		"ktp_number":     input.Ktp,
		"monthly_income": input.Income,
		"account_number": input.AccountNumber,
	}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot update user profile"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update profile",
		"data":    userr,
	})
}
