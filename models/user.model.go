package models

type User struct {
	UserID        uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Name          string    `gorm:"type:varchar(200);default:null" json:"name"`
	Email         string    `gorm:"type:varchar(50);unique;default:null" json:"email"`
	Password      string    `gorm:"type:varchar(255);not null" json:"password"`
	KTPNumber     string    `gorm:"type:varchar(16);unique;default:null" json:"ktp_number"`
	BirthDate     string    `gorm:"type:varchar(20);default:null" json:"birth_date"`
	PhoneNumber   string    `gorm:"type:varchar(15);not null" json:"phone_number"`
	Address       string    `gorm:"type:varchar(255);default:null" json:"address"`
	MonthlyIncome float64   `gorm:"type:decimal(10,2);default:null" json:"monthly_income"`
	AccountNumber string    `gorm:"type:varchar(20);unique;default:null" json:"account_number"`
	Status        string    `gorm:"type:varchar(20);default:null" json:"status"`
	Loans         []Loan    `gorm:"constraint:OnDelete:CASCADE;" json:"loans"`
	Payments      []Payment `gorm:"constraint:OnDelete:CASCADE;" json:"payments"`
}

type Admin struct {
	AdminID  uint   `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	Username string `gorm:"type:varchar(50);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}
