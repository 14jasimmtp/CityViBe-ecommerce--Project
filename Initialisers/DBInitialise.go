package initialisers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/domain"
)

var DB *gorm.DB

func DBInitialise() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error while connecting database")
	}
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.Admin{})
	DB.AutoMigrate(&domain.Product{})
	DB.AutoMigrate(&domain.Category{})
	DB.AutoMigrate(&domain.Size{})
	DB.AutoMigrate(&domain.Address{})
	DB.AutoMigrate(&domain.Cart{})
	DB.AutoMigrate(&domain.Order{})
	DB.AutoMigrate(&domain.OrderItem{})
	DB.AutoMigrate(&domain.Wishlist{})
	DB.AutoMigrate(&domain.Wallet{})
	DB.AutoMigrate(&domain.Coupon{})
	DB.AutoMigrate(&domain.PaymentMethod{})
	DB.AutoMigrate(&domain.RazorPay{})
}
