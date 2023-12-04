package models

type Coupon struct {
	Coupon            string  `json:"coupon"`
	DiscoutPercentage float64 `json:"DiscountRate"`
	UsageLimit        int     `json:"usage_limit"`
}

type CheckoutCoupon struct {
	Coupon string `json:"coupon"`
	Wallet bool   `json:"wallet"`
}
