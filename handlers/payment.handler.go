package handlers

import (
	"fmt"
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoanPayment(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	id := c.Params("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	Id := uint(ID)

	type inputPy struct {
		PaymentAmount float64 `json:"payment_amount"`
	}
	var input inputPy
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON"})
	}
	var loan models.Loan
	if err := databases.DB.First(&loan, "loan_id", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "cannot get loan"})
	}

	finalLoan := loan.LoanAmount - input.PaymentAmount

	status := "cicilan"
	if finalLoan <= 0 {
		status = "lunas"
		finalLoan = 0
	}

	t := time.Now()
	bulan := [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}
	tanggal := fmt.Sprintf("%d %s %d", t.Day(), bulan[t.Month()-1], t.Year())

	payment := models.Payment{
		LoanID:        Id,
		UserID:        user.UserID,
		PaymentAmount: input.PaymentAmount,
		PaymentDate:   tanggal,
		PaymentStatus: status,
	}

	if err := databases.DB.Model(&loan).Where("loan_id", Id).Updates(map[string]interface{}{"loan_amount": finalLoan, "status": status}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "cannot update loan amount"})
	}

	if err := databases.DB.Create(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "cannot create payment"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "201",
		"message": "success create payment",
		"data":    payment,
	})
}
