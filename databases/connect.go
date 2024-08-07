package databases

import (
	"log"
	"pinjemlah-fiber/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "host=localhost user=postgres password=admin dbname=pinjemlah port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	} else {
		log.Println("Connected")
	}

	err = db.AutoMigrate(&models.User{}, &models.Loan{}, &models.Payment{}, &models.Admin{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
	DB = db
}
