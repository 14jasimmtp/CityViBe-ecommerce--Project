package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
)

func ExecuteRazorPayPayment(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"error from orderID"})
		return
	}
	paymentMethodID, err := usecase.PaymentMethodID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if paymentMethodID == 2 {
		payment, _ := usecase.PaymentAlreadyPaid(orderID)
		if payment {
			c.HTML(http.StatusOK, "pay.html", nil)
			return

		}
		orderDetail,Razor_id,err := usecase.MakePaymentRazorPay(orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"final_price": orderDetail.Final_price,
			"razor_id":    Razor_id,
			"user_name":   orderDetail.Username,
			"total":       orderDetail.Total_price,
		})
	}
	c.HTML(http.StatusNotFound, "notfound.html", nil)
}
func VerifyPayment(c *gin.Context) {
	var RazorpayDetails models.PaymentVerify

	if c.ShouldBindJSON(&RazorpayDetails) != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"enter details correctly"})
		return
	}
	
	order,err:=usecase.PaymentVerification(RazorpayDetails)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"successfully updated payment details","order":order})

}
