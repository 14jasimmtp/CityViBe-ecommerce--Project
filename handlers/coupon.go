package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

// MakeCoupon godoc
// @Summary Create a new coupon
// @Description Create a new coupon using the provided details.
// @Tags Admin Coupon Management
// @Accept json
// @Produce json
// @Param coupon body models.Coupon true "Details of the coupon to be created"
// @Success 200 {object} string "message": "Coupon created successfully", "coupon": CouponDetails
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/coupon [post]
func MakeCoupon(c *gin.Context) {
	var Coupon models.Coupon

	if c.ShouldBindJSON(&Coupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Constraints correctly"})
		return
	}
	data, err := utils.Validation(Coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": data})
		return
	}

	CouponDetails, err := usecase.CreateCoupon(Coupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Coupon created successfully", "coupon": CouponDetails})
}

// DisableCoupon godoc
// @Summary Disable a coupon
// @Description Disable a coupon based on the provided coupon ID.
// @Tags Admin Coupon Management
// @Accept json
// @Produce json
// @Param coupon body models.CouponStatus true "Coupon ID to be disabled"
// @Success 200 {object} string "message": "Coupon disabled successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/coupon/disable [put]
func DisableCoupon(c *gin.Context) {
	var Coupon models.CouponStatus
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

// EnableCoupon godoc
// @Summary Enable a coupon
// @Description Enable a coupon based on the provided coupon ID.
// @Tags Admin Coupon Management
// @Accept json
// @Produce json
// @Param coupon body models.CouponStatus true "Coupon ID to be enabled"
// @Success 200 {object} string "message": "Coupon enabled successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/coupon/enable [put]
func EnableCoupon(c *gin.Context) {
	var Coupon models.CouponStatus
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

// ViewCouponsAdmin godoc
// @Summary View all coupons for admin
// @Description Retrieve details of all coupons for admin.
// @Tags Admin Coupon Management
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "All Coupons", "Coupons": Coupons
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/coupon [get]
func ViewCouponsAdmin(c *gin.Context) {
	Coupons, err := usecase.GetCouponsForAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Coupons", "Coupons": Coupons})
}

// func ViewCouponsUser(c *gin.Context)

// UpdateCoupon godoc
// @Summary Update a coupon
// @Description Update a coupon based on the provided details and coupon ID.
// @Tags Admin Coupon Management
// @Accept json
// @Produce json
// @Param couponID query string true "Coupon ID to be updated"
// @Param updateCoupon body models.Coupon true "Details of the coupon to be updated"
// @Success 200 {object} string "message": "Coupon updated successfully", "coupon": Coupon
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/coupon/update [put]
func UpdateCoupon(c *gin.Context) {
	var UpdateCoupon models.Coupon
	if c.ShouldBindJSON(&UpdateCoupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details correctly"})
		return
	}
	Error, err := utils.Validation(UpdateCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
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

// ViewCouponsUser godoc
// @Summary View coupons for user
// @Description Retrieve details of coupons for the authenticated user.
// @Tags User Profile
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "Coupons", "Coupons": coupons
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /coupons [get]
func ViewCouponsUser(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "error in token .relogin again."})
		return
	}
	coupons, err := usecase.ViewCouponsUser(Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coupons", "Coupons": coupons})
}

// @Summary Remove Coupon
// @Description Removes a coupon associated with the provided authorization token.
// @Tags User Order
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header with bearer token"
// @Success 200 {object} string "message": "coupon removed successfully"
// @Failure 401 {object} string "error": "error in token .relogin again."
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /removecoupon [post]
func RemoveCoupon(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "error in token .relogin again."})
		return
	}

	if err := usecase.RemoveCoupon(Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon removed successfully"})

}
