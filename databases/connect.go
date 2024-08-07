package databases

import (
	"log"
	"pinjemlah-fiber/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	const dsnProduction = "postgres://koyeb-adm:lu1LcrHRd4xN@ep-spring-block-a1syh4zb.ap-southeast-1.pg.koyeb.app/koyebdb"
	const dsnLocal = "host=localhost user=postgres password=admin dbname=pinjemlah port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsnProduction), &gorm.Config{})
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
