package handlers

import (
	"fmt"
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AdminGetUsers(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	var users []models.User
	if err := databases.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": users})
}
func AdminGetUserByID(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	var user models.User
	id := c.Params("id")
	if err := databases.DB.Preload("Loans").First(&user, "user_id", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

func AdminGetLoanStatus(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	var loans []models.Loan
	if err := databases.DB.Find(&loans, "status = 'lunas'").Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed get data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "get data loan",
		"data":    loans,
	})
}

func AdminGetLoanCicilan(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	var loans []models.Loan
	if err := databases.DB.Find(&loans, "status = 'cicilan'").Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed get data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "get data loan",
		"data":    loans,
	})
}

func AdminGetLoans(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	var loans []models.Loan
	if err := databases.DB.Find(&loans).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed get data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "get data loan",
		"data":    loans,
	})
}

func AdminGetLoanByID(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	loan_id := c.Params("id")
	var loan models.Loan
	if err := databases.DB.Preload("Payments").First(&loan, "loan_id", loan_id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed get data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "get data loan",
		"data":    loan,
	})
}

func AdminUpdateStatusLoan(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	loan_id := c.Params("id")
	id, err := strconv.Atoi(loan_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	ID := uint(id)
	type inputStatus struct {
		Status string `json:"status"`
	}
	var input inputStatus
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "failed get data"})
	}

	var loan models.Loan
	if err := databases.DB.First(&loan, "loan_id = ?", ID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "loan not found"})
	}
	term := loan.LoanTerm
	t := time.Now()
	bulan := [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	aprovalDate := fmt.Sprintf("%d %s %d", t.Day(), bulan[t.Month()-1], t.Year())
	disbursementDate := t.AddDate(0, term, 0)
	disbursementDateFormatted := fmt.Sprintf("%d %s %d", disbursementDate.Day(), bulan[disbursementDate.Month()-1], disbursementDate.Year())

	if err := databases.DB.Model(&loan).Updates(map[string]interface{}{"status": input.Status, "approval_date": aprovalDate, "disbursement_date": disbursementDateFormatted}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed update data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "update data loan",
		"data":    loan,
	})
}

func GetPayments(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	var payments []models.Payment
	if err := databases.DB.Find(&payments).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "payments not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "get data payment",
		"data":    payments,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	_, ok := c.Locals("admin").(models.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	user_id := c.Params("id")
	id, err := strconv.Atoi(user_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}

	var user models.User
	if err := databases.DB.First(&user, "user_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}

	if err := databases.DB.Delete(&models.Loan{}, "user_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to delete loans"})
	}
	if err := databases.DB.Delete(&models.Payment{}, "user_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to delete payments"})
	}

	if err := databases.DB.Delete(&user, "user_id = ? ", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to delete user"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "user deleted successfully",
	})
}
