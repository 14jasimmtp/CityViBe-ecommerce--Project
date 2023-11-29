package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
)

func MakeCoupon(c *gin.Context) {
	var Coupon models.Coupon

	if c.ShouldBindJSON(&Coupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Constraints correctly"})
		return
	}

	CouponDetails, err := usecase.CreateCoupon(Coupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Coupon created successfully", "coupon": CouponDetails})
}

func DisableCoupon(c *gin.Context) {
	var Coupon struct {
		CouponID uint `json:"coupon_id"`
	}
	if c.ShouldBindJSON(&Coupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details Correctly"})
		return
	}
	err := usecase.DisableCoupon(Coupon.CouponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon disabled successfully"})

}

func EnableCoupon(c *gin.Context) {
	var Coupon struct {
		CouponID uint `json:"coupon_id"`
	}
	if c.ShouldBindJSON(&Coupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details Correctly"})
		return
	}
	err := usecase.EnableCoupon(Coupon.CouponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon enabled successfully"})

}

func ViewCouponsAdmin(c *gin.Context) {
	Coupons, err := usecase.GetCouponsForAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Coupons", "Coupons": Coupons})
}

func UpdateCoupon(c *gin.Context) {
	var UpdateCoupon models.Coupon
	if c.ShouldBindJSON(&UpdateCoupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details correctly"})
		return
	}
	coupon_id := c.Query("couponID")

	Coupon, err := usecase.UpdateCoupon(UpdateCoupon, coupon_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon updated successfully", "coupon": Coupon})
}
