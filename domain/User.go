package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint    `json:"id" gorm:"unique;not null"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Email      string  `json:"email" validate:"email"`
	Phone      string  `json:"phone"`
	Password   string  `json:"-" validate:"min=8,max=20"`
	Blocked    bool    `json:"blocked" gorm:"default:false"`
	Wallet     float64 `json:"wallet" gorm:"default:0"`
}

type Address struct {
	Id        int    `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	Phone     string `json:"phone" validate:"required min=10 max=10"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type Wallet struct {
	ID     uint    `json:"id" gorm:"unique;not null"`
	UserID uint    `json:"user_id"`
	User   User    `json:"-" gorm:"foreignkey:UserID"`
	Amount float64 `json:"Balance" gorm:"default:0;not null"`
}
