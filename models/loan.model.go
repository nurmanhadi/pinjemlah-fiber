package models

type Loan struct {
	LoanID           uint      `gorm:"primaryKey;autoIncrement" json:"loan_id"`
	UserID           uint      `gorm:"constraint:OnDelete:CASCADE;" json:"user_id"`
	LoanType         string    `gorm:"type:varchar(10);not null" json:"loan_type"`
	LoanAmount       float64   `gorm:"type:decimal(10,2);not null" json:"loan_amount"`
	InterestRate     float64   `gorm:"type:decimal(10,2);not null" json:"interest_rate"`
	LoanTerm         int       `gorm:"not null" json:"loan_term"`
	Status           string    `gorm:"type:varchar(15);not null" json:"status"`
	ApplicationDate  string    `gorm:"type:varchar(50);not null" json:"application_date"`
	ApprovalDate     string    `gorm:"type:varchar(50);default:null" json:"approval_date"`
	DisbursementDate string    `gorm:"type:varchar(50);default:null" json:"disbursement_date"`
	LoanSlug         string    `gorm:"type:varchar(50);not null" json:"loan_slug"`
	Payments         []Payment `json:"payments"`
}
