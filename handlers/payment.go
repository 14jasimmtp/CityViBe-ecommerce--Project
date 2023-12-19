package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

// ExecuteRazorPayPayment godoc
// @Summary Execute RazorPay payment for a given order
// @Description Execute RazorPay payment for a specified order. Returns necessary details for the payment process.
// @Tags Payments
// @Accept json
// @Produce json
// @Param order_id query int true "Order ID"
// @Success 200 {object} string "final_price": "Final Price", "razor_id": "RazorPay ID", "user_name": "User Name", "total": "Total Price"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 404 {object} string "error": "Not Found"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /payment/razorpay [Get]
func ExecuteRazorPayPayment(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error from orderID"})
		return
	}
	paymentMethodID, err := usecase.PaymentMethodID(orderID)
	if err != nil {

		c.HTML(http.StatusNotFound, "notfound.html", nil)
		return
	}
	if paymentMethodID == 2 {
		payment, _ := usecase.PaymentAlreadyPaid(orderID)
		if payment {
			c.HTML(http.StatusOK, "pay.html", nil)
			return

		}
		orderDetail, Razor_id, err := usecase.MakePaymentRazorPay(orderID)
		if err != nil {
			c.HTML(http.StatusNotFound, "notfound.html", nil)
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"final_price": orderDetail.Final_price,
			"razor_id":    Razor_id,
			"user_name":   orderDetail.Username,
			"total":       orderDetail.Total_price,
		})
		return
	}
	c.HTML(http.StatusNotFound, "notfound.html", nil)
}

// VerifyPayment godoc
// @Summary Verify payment for a given order
// @Description Verify payment for a specified order using payment details. Returns updated order details after verification.
// @Tags Payments
// @Accept json
// @Produce json
// @Param orderId query int true "Order ID"
// @Param input body models.PaymentVerify true "Payment Verification Details"
// @Success 200 {object} string "message": "Updated payment details successfully", "Order Details": models.Order
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /payment/verify [post]
func VerifyPayment(c *gin.Context) {
	var Verify models.PaymentVerify
	if c.ShouldBindJSON(&Verify) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter fields correctly"})
		return
	}
	Error, err := utils.Validation(Verify)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}
	order := c.Query("orderId")
	order_id, err := strconv.Atoi(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong"})
	}
	Order, err := usecase.VerifyPayment(Verify, order_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated payment details successfully", "Order Details": Order})
}
