package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

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
func VerifyPayment(c *gin.Context) {
	var Verify models.PaymentVerify
	if c.ShouldBindJSON(&Verify) != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"enter fields correctly"})
		return
	}
	Error,err:=utils.Validation(Verify)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":Error})
		return
	}
	err = usecase.SavePaymentDetails(Verify.OrderID, Verify.PaymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated payment details successfully"})
}
