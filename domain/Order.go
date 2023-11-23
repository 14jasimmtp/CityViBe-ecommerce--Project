package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        int     `json:"user_id" gorm:"not null"`
	User          User    `json:"-" gorm:"foreignkey:UserID"`
	AddressID     int     `json:"address_id" gorm:"not null"`
	Address       Address `json:"-" gorm:"foreignkey:AddressID"`
	CartID        int     `json:"cart_id" gorm:"not null"`
	Cart          Cart    `json:"-" gorm:"foreignkey:CartID"`
	OrderStatus   string  `json:"order_status" gorm:"default:'pending'"`
	PaymentStatus string  `json:"payment_status" gorm:"default:'not paid'"`
	FinalPrice    float64 `json:"final_price"`
	Approval      bool    `json:"approval" gorm:"default:false"`
}
