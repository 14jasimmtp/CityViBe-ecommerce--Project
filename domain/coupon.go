package domain

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	Coupon             string  `json:"coupon" gorm:"unique;not null"`
	DiscountPercentage float64 `json:"discount_rate" gorm:"not null"`
	Valid              bool    `json:"valid" gorm:"default:true"`
}
