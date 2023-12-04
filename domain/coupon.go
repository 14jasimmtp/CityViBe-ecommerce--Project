package domain

import (
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Coupon             string  `json:"coupon" gorm:"unique;not null"`
	DiscountPercentage float64 `json:"discount_rate" gorm:"not null"`
	UsageLimit         int     `json:"usage_limit"`
	Active             bool    `json:"active" gorm:"default:true"`
}

type UsedCoupon struct {
	UserId int    `json:"userid"`
	Coupon string `json:"coupon"`
}
