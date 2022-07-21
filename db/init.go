package db

import (
	"fmt"
	"log"

	"github.com/awesome-sphere/as-payment-consumer/db/models"
	"github.com/awesome-sphere/as-payment-consumer/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dbUser := utils.GetenvOr("POSTGRES_USER", "pkinwza")
	dbPassword := utils.GetenvOr("POSTGRES_PASSWORD", "securepassword")
	dbHost := utils.GetenvOr("POSTGRES_HOST", "localhost")
	dbPort := utils.GetenvOr("POSTGRES_PORT", "5434")
	dbName := utils.GetenvOr("POSTGRES_DB", "as-payment")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&models.Order{}, &models.OrderSeats{})

	DB = db
}
