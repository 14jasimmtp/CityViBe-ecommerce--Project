package models

type Coupon struct {
	Coupon            string  `json:"coupon"`
	DiscoutPercentage float64 `json:"DiscountRate"`
}
