package models

type Payment struct {
	PaymentID     uint    `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	UserID        uint    `gorm:"constraint:OnDelete:CASCADE;" json:"user_id"`
	LoanID        uint    `gorm:"constraint:OnDelete:CASCADE;" json:"loan_id"`
	PaymentDate   string  `json:"payment_date"`
	PaymentAmount float64 `json:"payment_amount"`
	PaymentStatus string  `json:"payment_status"`
}
