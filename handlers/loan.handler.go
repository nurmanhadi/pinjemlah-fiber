package handlers

import (
	"fmt"
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddLoan(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	type inputLoan struct {
		LoanAmount float64 `json:"loan_amount"`
		LoanTerm   int     `json:"loan_term"`
	}

	var input inputLoan
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON"})
	}

	rate := 3.0
	tipe := "konsumsi"

	t := time.Now()

	bulan := [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	tanggal := fmt.Sprintf("%d %s %d", t.Day(), bulan[t.Month()-1], t.Year())
	persen := 0.05
	finalAmount := input.LoanAmount + (input.LoanAmount * persen)
	slug := strconv.FormatFloat(finalAmount, 'f', 2, 64)

	loan := models.Loan{
		UserID:          user.UserID,
		LoanAmount:      finalAmount,
		LoanTerm:        input.LoanTerm,
		InterestRate:    rate,
		LoanType:        tipe,
		ApplicationDate: tanggal,
		LoanSlug:        slug,
		Status:          "pending",
	}
	if err := databases.DB.Create(&loan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "registrasi gagal"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "201",
		"message": "registrasi berhasil",
		"data":    loan,
	})
}

func GetLoans(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	var loans []models.Loan
	if err := databases.DB.Where("user_id = ?", user.UserID).Find(&loans).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed get data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "get data loan",
		"data":    loans,
	})
}

func GetLoanByID(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	id := c.Params("id")

	var loan models.Loan
	if err := databases.DB.Preload("Payments").First(&loan, "loan_id = ?", id).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed get data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "get data loan",
		"data":    loan,
	})
}
