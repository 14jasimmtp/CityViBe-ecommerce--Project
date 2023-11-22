package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint    `json:"user_id" gorm:"uniquekey; not null"`
	Users     User    `json:"-" gorm:"foreignkey:UserID"`
	ProductID uint    `json:"product_id"`
	Products  Product `json:"-" gorm:"foreignkey:ProductID"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
}

// type CartItem struct {
// 	ID        uint    `json:"-"`
// 	UserID    uint    `json:"user_id" gorm:"uniquekey; not null"`
// 	Users     User    `json:"-" gorm:"foreignkey:UserID"`
// 	ProductID uint    `json:"product_id"`
// 	Products  Product `json:"-" gorm:"foreignkey:ProductID"`
// 	Quantity  float64 `json:"quantity"`
// 	Price     float64 `json:"price"`
// }

// type CartII struct {
// 	ID uint `json:"cartId"`

// 	CartItems []CartItem `json:"CartItems" gorm:"foreignkey:"`
// }
