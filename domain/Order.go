package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          int           `json:"user_id" gorm:"not null"`
	User            User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       int           `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	PaymentStatus   string        `json:"payment_status" gorm:"default:'not paid'"`
	TotalPrice      float64       `json:"total_price"`
	FinalPrice      float64       `json:"final_price"`
	Approval        bool          `json:"approval" gorm:"default:false"`
}

type OrderItem struct {
	ID          uint    `json:"id" gorm:"primaryKey;not null"`
	OrderID     uint    `json:"order_id"`
	Order       Order   `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	ProductID   uint    `json:"product_id"`
	Products    Product `json:"-" gorm:"foreignkey:ProductID"`
	UserID      int     `json:"user_id"`
	User        User    `json:"-" gorm:"foreignkey:UserID"`
	OrderStatus string  `json:"order_status" gorm:"default:'pending'"`
	Quantity    float64 `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	FinalPrice  float64 `json:"final_price"`
}
