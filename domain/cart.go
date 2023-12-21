package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint    `json:"user_id" gorm:"uniquekey; not null"`
	Users      User    `json:"-" gorm:"foreignkey:UserID"`
	ProductID  uint    `json:"product_id"`
	Products   Product `json:"-" gorm:"foreignkey:ProductID"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
	FinalPrice float64 `json:"final_price" gorm:"default:0"`
}
