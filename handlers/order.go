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

// OrderFromCart godoc
// @Summary Place an order from the user's cart
// @Description Place an order using the provided checkout details.
// @Tags User Order
// @Accept json
// @Produce json
// @Param OrderInput body models.CheckOut true "Details for the order checkout"
// @Success 200 {object} string "message": "Ordered products successfully", "order Details": OrderDetails
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders [post]
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

// ViewCheckOut godoc
// @Summary View the checkout page
// @Description Retrieve details for the user's checkout page.
// @Tags User Order
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "CheckOut Page loaded successfully", "order Details": OrderDetails
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /checkout [get]
func ViewCheckOut(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.CheckOut(Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CheckOut Page loaded successfully", "order Details": OrderDetails})
}

// ViewOrders godoc
// @Summary View user orders
// @Description Retrieve details of orders for the authenticated user.
// @Tags User Order
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "Orders", "order Details": OrderDetails
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders [get]
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

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel an order for the authenticated user based on the provided order and product IDs.
// @Tags User Order
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID to be cancelled"
// @Param product_id query string true "Product ID in the order to be cancelled"
// @Success 200 {object} string "message": "Order cancelled successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders/cancel [put]
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

// CancelOrderByAdmin godoc
// @Summary Cancel an order by admin
// @Description Cancel an order by admin based on the provided user ID, order ID, and product ID.
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Param cancel body models.AdminOrder true "Details for cancelling an order by admin"
// @Success 200 {object} string "message": "Order cancelled successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/orders/cancel [post]
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

// ShipOrderByAdmin godoc
// @Summary Ship an order by admin
// @Description Ship an order by admin based on the provided user ID, order ID, and product ID.
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Param ship body models.AdminOrder true "Details for shipping an order by admin"
// @Success 200 {object} string "message": "Order shipped successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders/ship [post]
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

// DeliverOrderByAdmin godoc
// @Summary Deliver an order by admin
// @Description Deliver an order by admin based on the provided user ID, order ID, and product ID.
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Param deliver body models.AdminOrder true "Details for delivering an order by admin"
// @Success 200 {object} string "message": "Order delivered successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders/deliver [post]
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

// ReturnOrder godoc
// @Summary Return an order
// @Description Return an order for the authenticated user based on the provided order and product IDs.
// @Tags User Order
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID to be returned"
// @Param product_id query string true "Product ID in the order to be returned"
// @Success 200 {object} string "message": "Order returned successfully. Amount will be credited to wallet."
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders/return [put]
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

// SalesReportByDate godoc
// @Summary Generate and download sales report for a specific date range
// @Description Generate and download sales report in PDF format for a specific date range.
// @Tags Admin
// @Accept json
// @Produce json
// @Param start formData string true "Start date (format: DD-MM-YYYY)"
// @Param end formData string true "End date (format: DD-MM-YYYY)"
// @Success 200 {file} application/pdf
// @Failure 400 {object} string "error": "Bad Request"
// @Router /admin/salesreportbydate [post]
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

	c.Header("Content-Disposition", "attachment;filename=Invoice.pdf")
	c.Header("Content_Type", "application/pdf")

	err = report.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pdfFilePath := "Invoice/salesreport.pdf"

	err = report.OutputFileAndClose(pdfFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.File(pdfFilePath)
	// c.JSON(http.StatusOK, gin.H{"report": report})
}

// SalesReportByPeriod godoc
// @Summary Generate and download sales report for a specific period
// @Description Generate and download sales report in PDF format for a specific period.
// @Tags Admin
// @Accept json
// @Produce json
// @Param period formData string true "Period (e.g., 'last week', 'last month')"
// @Success 200 {file} application/pdf
// @Failure 400 {object} string "error": "Bad Request"
// @Router /admin/salesreportbyperiod [post]
func SalesReportByPeriod(c *gin.Context) {
	period := c.PostForm("period")

	report, err := usecase.ExecuteSalesReportByPeriod(period)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Header("Content-Disposition", "attachment;filename=Invoice.pdf")
	c.Header("Content_Type", "application/pdf")

	err = report.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pdfFilePath := "Invoice/salesreport.pdf"

	err = report.OutputFileAndClose(pdfFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.File(pdfFilePath)
	// c.JSON(http.StatusOK, gin.H{"report": report})
}

// SalesReportByPayment godoc
// @Summary Generate and download sales report for a specific payment method and date range
// @Description Generate and download sales report in PDF format for a specific payment method and date range.
// @Tags Admin
// @Accept json
// @Produce json
// @Param start formData string true "Start date (format: DD-MM-YYYY)"
// @Param end formData string true "End date (format: DD-MM-YYYY)"
// @Param paymentmethod formData string true "Payment method (e.g., 'credit card', 'wallet')"
// @Success 200 {file} application/pdf
// @Failure 400 {object} string "error": "Bad Request"
// @Router /admin/salesreportbypayment [post]
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

	c.Header("Content-Disposition", "attachment;filename=Invoice.pdf")
	c.Header("Content_Type", "application/pdf")

	err = report.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pdfFilePath := "Invoice/salesreport.pdf"

	err = report.OutputFileAndClose(pdfFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.File(pdfFilePath)
	// c.JSON(http.StatusOK, gin.H{"report": report})
}

// PrintInvoice godoc
// @Summary Generate and download invoice for a specific order
// @Description Generate and download invoice in PDF format for a specific order.
// @Tags User Order
// @Accept json
// @Produce json
// @Param orderid query int true "Order ID"
// @Success 200 {file} application/pdf
// @Failure 400 {object} string "error": "Bad Request"
// @Router /Invoice [get]
func PrintInvoice(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
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
	pdf, err := usecase.PrintInvoice(orderid, Token)
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

// ApplyCoupon godoc
// @Summary Apply a coupon to the user's account
// @Description Apply a coupon code to the user's account for potential discounts or benefits.
// @Tags User Order
// @Accept json
// @Produce json
// @Param coupon formData string true "Coupon Code"
// @Security ApiKeyAuth
// @Success 200 {object} string "message": "coupon applied successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /applycoupon [post]
func ApplyCoupon(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login to apply coupon"})
		return
	}
	coupon := c.PostForm("coupon")
	err = usecase.ApplyCoupon(coupon, Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon applied successfully"})
}

// @Summary Generate Excel Sales Report
// @Description Generate a stylish Excel sales report based on the provided start and end dates.
// @Tags Admin
// @Accept json
// @Produce json
// @Param StartDate query string true "Start date (format: dd-mm-yyyy)"
// @Param EndDate query string true "End date (format: dd-mm-yyyy)"
// @Success 200 {file} binary "Excel sales report"
// @Failure 400 {object}  string "error":"Bad Request"
// @Failure 500 {object}  string "error":"Internal Server Error"
// @Router /admin/salesreport/excel [post]
func SalesReportXL(c *gin.Context) {
	strStartDate := c.Query("StartDate")
	strEndDate := c.Query("EndDate")

	startDate, err := time.Parse("2-1-2006", strStartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endDate, err := time.Parse("2-1-2006", strEndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	report, err := usecase.SalesReportXL(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=salesReport.xlsx")
	c.Header("Expires", "0")

	err = report.Write(c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to write Excel file")
		return
	}

}
