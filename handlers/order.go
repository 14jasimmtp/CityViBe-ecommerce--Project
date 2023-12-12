package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

func OrderFromCart(c *gin.Context) {
	var OrderInput models.CheckOut

	if c.ShouldBindJSON(&OrderInput) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Enter constraints correctly"})
		return
	}
	Error, err := utils.Validation(OrderInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if OrderInput.PaymentID == 1 || OrderInput.PaymentID == 2 {
		OrderDetails, err := usecase.ExecutePurchase(Token, OrderInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})
		}

	} else if OrderInput.PaymentID == 3 {
		OrderDetails, err := usecase.ExecutePurchaseWallet(Token, OrderInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter a valid payment method"})
		return
	}
}

func ViewCheckOut(c *gin.Context) {
	var coupon models.CheckoutCoupon
	if c.ShouldBindJSON(&coupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter coupon correctly"})
	}
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.CheckOut(Token, coupon.Coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CheckOut Page loaded successfully", "order Details": OrderDetails})
}

func ViewOrders(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.ViewUserOrders(Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "orders", "order Details": OrderDetails})

}

func CancelOrder(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Query("order_id")
	product_id := c.Query("product_id")

	err = usecase.CancelOrder(Token, orderId, product_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order cancelled successfully"})

}

func CancelOrderByAdmin(c *gin.Context) {
	var cancel models.AdminOrder
	if c.ShouldBindJSON(&cancel) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Constraints correctly"})
		return
	}

	Error, err := utils.Validation(cancel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}
	err = usecase.CancelOrderByAdmin(cancel.UserID, cancel.OrderID, cancel.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

func ShipOrderByAdmin(c *gin.Context) {
	var Ship models.AdminOrder
	if c.ShouldBindJSON(&Ship) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter constraints correctly"})
		return
	}

	Error, err := utils.Validation(Ship)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
	}

	err = usecase.ShipOrders(Ship.UserID, Ship.OrderID, Ship.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Shipped successfully"})
}

func DeliverOrderByAdmin(c *gin.Context) {
	var Deliver models.AdminOrder

	if c.ShouldBindJSON(&Deliver) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter constraints correctly"})
		return
	}

	Error, err := utils.Validation(Deliver)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}

	err = usecase.DeliverOrder(Deliver.UserID, Deliver.OrderID, Deliver.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order delivered successfully"})

}

func ReturnOrder(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Query("order_id")
	productId := c.Query("product_id")

	err = usecase.ReturnOrder(Token, orderId, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order returned successfully.Amount will be credited to wallet."})
}

func SalesReportByDate(c *gin.Context) {
	startDateStr := c.PostForm("start")
	endDateStr := c.PostForm("end")
	startDate, err := time.Parse("2-1-2006", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endDate, err := time.Parse("2-1-2006", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	report, err := usecase.ExecuteSalesReportByDate(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"report": report})
}

func SalesReportByPeriod(c *gin.Context) {
	period := c.PostForm("period")

	report, err := usecase.ExecuteSalesReportByPeriod(period)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"report": report})
}

func SalesReportByPayment(c *gin.Context) {
	startDateStr := c.PostForm("start")
	endDateStr := c.PostForm("end")
	paymentmethod := c.PostForm("paymentmethod")
	startDate, err := time.Parse("2-1-2006", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	endDate, err := time.Parse("2-1-2006", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	report, err := usecase.ExecuteSalesReportByPaymentMethod(startDate, endDate, paymentmethod)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"report": report})
}

func PrintInvoice(c *gin.Context) {
	Token,err:=c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	strorderId := c.Query("orderid")
	orderid, err := strconv.Atoi(strorderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pdf, err := usecase.PrintInvoice(orderid,Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment;filename=Invoice.pdf")
	c.Header("Content_Type", "application/pdf")

	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pdfFilePath := "Invoice/invoice.pdf"

	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.File(pdfFilePath)

}
