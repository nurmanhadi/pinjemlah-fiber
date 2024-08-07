package handlers

import (
	"fmt"
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf/v2"
)

func AdminReportUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := databases.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "gagal get users"})
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Logo
	pdf.Image("assets/logo.png", 10, 10, 30, 0, false, "", 0, "")

	// Tulisan di tengah sejajar dengan logo
	pdf.SetFont("Times", "B", 16)
	pdf.SetXY(50, 10)
	pdf.CellFormat(0, 10, "PINJEMLAH", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Alamat
	pdf.SetFont("Times", "", 12)
	pdf.SetXY(50, 20)
	pdf.MultiCell(0, 10, "Jl. Sahid, No. 8, Rt. 02/01, Kec. Ps. Minggu, Kota Jakarta Selatan Daerah Khusus, 12510", "", "C", false)
	pdf.Ln(2)

	// Garis bawah alamat
	pdf.SetDrawColor(0, 0, 0)
	pdf.Line(10, 40, 200, 40)
	pdf.Ln(10)

	pdf.SetFont("Times", "B", 14)
	pdf.CellFormat(0, 10, "Laporan Users", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Tabel User
	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(10, 10, "No", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Nama", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Email", "1", 0, "C", false, 0, "")
	pdf.CellFormat(63, 10, "Alamat", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Telepon", "1", 1, "C", false, 0, "")

	pdf.SetFont("Times", "", 12)
	for i, user := range users {
		pdf.CellFormat(10, 10, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, user.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(40, 10, user.Email, "1", 0, "L", false, 0, "")
		pdf.CellFormat(63, 10, user.Address, "1", 0, "L", false, 0, "")
		pdf.CellFormat(40, 10, user.PhoneNumber, "1", 1, "L", false, 0, "")
	}
	pdf.Ln(5)

	// Total Users di kiri
	pdf.SetFont("Times", "", 12)
	pdf.CellFormat(0, 10, fmt.Sprintf("Total Users: %d", len(users)), "", 1, "L", false, 0, "")

	// Tanggal dan Manager
	pdf.SetFont("Times", "", 12)
	t := time.Now()
	indonesianMonths := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	dateString := fmt.Sprintf("Jakarta, %d %s %d", t.Day(), indonesianMonths[t.Month()-1], t.Year())

	pdf.SetY(-90)
	pdf.SetX(-70)
	pdf.CellFormat(0, 10, dateString, "", 1, "R", false, 0, "")
	pdf.CellFormat(0, 10, "Manager", "", 1, "R", false, 0, "")
	pdf.Ln(20)
	pdf.CellFormat(0, 10, "_______________________", "", 1, "R", false, 0, "")
	pdf.CellFormat(0, 10, "Muhammad Nurman Hadi", "", 1, "R", false, 0, "")

	// Output PDF
	err := pdf.OutputFileAndClose("user_report.pdf")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating PDF",
			"error":   err.Error(),
		})
	}

	return c.SendFile("user_report.pdf")
}
