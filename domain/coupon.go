package domain

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	Coupon string  `gorm:"UNIQUE;NOT NULL"`
	Value  float64 `gorm:"NOT NULL"`
}
