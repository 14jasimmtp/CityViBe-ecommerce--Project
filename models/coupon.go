package models

type Coupon struct {
	Coupon            string  `json:"coupon" validate:"required"`
	DiscoutPercentage float64 `json:"DiscountRate" validate:"required"`
	UsageLimit        int     `json:"usage_limit" validate:"required"`
}

type CheckoutCoupon struct {
	Coupon string `json:"coupon"`
	Wallet bool   `json:"wallet"`
}

type Couponlist struct {
	Coupon             string  `json:"coupon" validate:"required"`
	DiscountPercentage float64 `json:"Discount_percentage" validate:"required"`
	
}

type CouponStatus struct {
	CouponID uint `json:"coupon_id"`
}