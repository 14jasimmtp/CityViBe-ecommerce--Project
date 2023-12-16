package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
	"main.go/domain"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func CheckOut(Token string) (interface{}, error) {
	userId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.CheckOutInfo{}, err
	}

	AllUserAddress, err := repository.ViewAddress(userId)
	if err != nil {
		return models.CheckOutInfo{}, err
	}

	AllCartProducts, err := repository.DisplayCart(userId)
	if err != nil {
		return models.CheckOutInfo{}, err
	}

	TotalAmount, err := repository.CartTotalAmount(userId)
	if err != nil {
		return models.CheckOutInfo{}, err
	}

	if AllCartProducts[0].FinalPrice == 0 {
		return models.CheckOutInfo{
			Address:     AllUserAddress,
			Cart:        AllCartProducts,
			TotalAmount: TotalAmount,
		}, nil
	} else {

		finalPrice, err := repository.CartFinalPrice(userId)
		if err != nil {
			return models.CheckOutInfo{}, err
		}
		for i := 0; i < len(AllCartProducts); i++ {
			AllCartProducts[i].Price = AllCartProducts[i].FinalPrice
		}
		return models.CheckOutInfoDiscount{
			Address:        AllUserAddress,
			Cart:           AllCartProducts,
			TotalAmount:    TotalAmount,
			DiscountAmount: finalPrice,
		}, nil
	}
}

func ExecutePurchase(Token string, OrderInput models.CheckOut) (models.OrderSuccessResponse, error) {
	var TotalAmount float64
	var method string
	userId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	addressExist := repository.CheckAddressExist(userId, OrderInput.AddressID)
	if !addressExist {
		return models.OrderSuccessResponse{}, errors.New(`address doesn't exist`)
	}

	paymentExist := repository.CheckPaymentMethodExist(OrderInput.PaymentID)
	if !paymentExist {
		return models.OrderSuccessResponse{}, errors.New(`payment method doesn't exist`)
	}
	if OrderInput.PaymentID == 1 {
		method = "COD"
	} else {
		method = "Razorpay"
	}

	cartExist := repository.CheckCartExist(userId)
	if !cartExist {
		return models.OrderSuccessResponse{}, errors.New(`cart is empty`)
	}

	cartItems, err := repository.DisplayCart(userId)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}
	if cartItems[0].FinalPrice != 0 {

		for i := 0; i < len(cartItems); i++ {
			cartItems[i].Price = cartItems[i].FinalPrice
		}

		TotalAmount, err = repository.CartFinalPrice(userId)
		if err != nil {
			return models.OrderSuccessResponse{}, errors.New(`error while calculating total amount`)
		}
	} else {
		TotalAmount, err = repository.CartTotalAmount(userId)
		if err != nil {
			return models.OrderSuccessResponse{}, errors.New(`error while calculating total amount`)
		}
	}

	OrderID, err := repository.OrderFromCart(OrderInput.AddressID, OrderInput.PaymentID, userId, TotalAmount)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	if err := repository.AddOrderProducts(userId, OrderID, cartItems); err != nil {
		return models.OrderSuccessResponse{}, err
	}

	var orderItemDetails domain.OrderItem
	for _, c := range cartItems {
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = c.Quantity
		err := repository.UpdateCartAndStockAfterOrder(userId, int(orderItemDetails.ProductID), orderItemDetails.Quantity)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
	}
	return models.OrderSuccessResponse{
		OrderID:       OrderID,
		PaymentMethod: method,
		TotalAmount:   TotalAmount,
		PaymentStatus: "not paid",
	}, nil
}

func ExecutePurchaseWallet(Token string, OrderInput models.CheckOut) (models.OrderSuccessResponse, error) {
	var TotalAmount float64
	userId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	user, err := repository.GetUserById(int(userId))
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	addressExist := repository.CheckAddressExist(userId, OrderInput.AddressID)
	if !addressExist {
		return models.OrderSuccessResponse{}, errors.New(`address doesn't exist`)
	}

	paymentExist := repository.CheckPaymentMethodExist(OrderInput.PaymentID)
	if !paymentExist {
		return models.OrderSuccessResponse{}, errors.New(`payment method doesn't exist`)
	}

	cartExist := repository.CheckCartExist(userId)
	if !cartExist {
		return models.OrderSuccessResponse{}, errors.New(`cart is empty`)
	}

	cartItems, err := repository.DisplayCart(userId)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	if cartItems[1].FinalPrice != 0 {

		for i := 0; i < len(cartItems); i++ {
			cartItems[i].Price = cartItems[i].FinalPrice
		}

		TotalAmount, err = repository.CartFinalPrice(userId)
		if err != nil {
			return models.OrderSuccessResponse{}, errors.New(`error while calculating total amount`)
		}
	} else {
		TotalAmount, err = repository.CartTotalAmount(userId)
		if err != nil {
			return models.OrderSuccessResponse{}, errors.New(`error while calculating total amount`)
		}
	}

	if user.Wallet < TotalAmount {
		return models.OrderSuccessResponse{}, errors.New(`insufficient Balance in Wallet.Add amount to wallet to purchase`)
	}

	OrderID, err := repository.OrderFromCart(OrderInput.AddressID, OrderInput.PaymentID, userId, TotalAmount)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	if err := repository.AddOrderProducts(userId, OrderID, cartItems); err != nil {
		return models.OrderSuccessResponse{}, err
	}
	_, err = repository.UpdateShipmentAndPaymentByOrderID("processing", "paid", OrderID)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	user.Wallet -= TotalAmount

	err = repository.UpdateWallet(user.Wallet, userId)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	var orderItemDetails domain.OrderItem
	for _, c := range cartItems {
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = c.Quantity
		err := repository.UpdateCartAndStockAfterOrder(userId, int(orderItemDetails.ProductID), orderItemDetails.Quantity)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
	}
	return models.OrderSuccessResponse{
		OrderID:       OrderID,
		PaymentMethod: "Wallet",
		TotalAmount:   TotalAmount,
		PaymentStatus: "paid",
	}, nil
}

func ViewUserOrders(Token string) ([]models.ViewOrderDetails, error) {
	userId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return []models.ViewOrderDetails{}, err
	}

	OrderDetails, err := repository.GetOrderDetails(userId)
	if err != nil {
		return []models.ViewOrderDetails{}, err
	}
	return OrderDetails, nil
}

func CancelOrder(Token, orderId string, pid string) error {
	UserID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}
	err = repository.CheckOrder(orderId, UserID)
	if err != nil {
		return errors.New(`no orders found with this id`)
	}

	OrderDetails, err := repository.CancelOrderDetails(UserID, orderId, pid)
	if err != nil {
		return err
	}

	if OrderDetails.OrderStatus == "Delivered" {
		return errors.New(`the order is delivered .Can't Cancel`)
	}

	if OrderDetails.OrderStatus == "Cancelled" {
		return errors.New(`the order is already cancelled`)
	}

	if OrderDetails.PaymentStatus == "paid" {
		err := repository.ReturnAmountToWallet(UserID, orderId, pid)
		if err != nil {
			return err
		}

	}
	err = repository.UpdateOrderFinalPrice(OrderDetails.OrderID, OrderDetails.TotalPrice)
	if err != nil {
		return err
	}
	proid, _ := strconv.Atoi(pid)
	err = repository.UpdateStock(proid, OrderDetails.Quantity)
	if err != nil {
		return err
	}

	err = repository.CancelOrder(orderId, pid, UserID)
	if err != nil {
		return err
	}

	return nil

}

func CancelOrderByAdmin(userID, order_id, pid int) error {
	orderID := strconv.Itoa(order_id)
	Pid := strconv.Itoa(pid)
	err := repository.CheckOrder(orderID, uint(userID))
	fmt.Println(err)
	if err != nil {
		return errors.New(`no orders found with this id`)
	}
	err = repository.CheckSingleOrder(Pid, orderID, uint(userID))
	if err != nil {
		return err
	}
	OrderDetails, err := repository.CancelOrderDetails(uint(userID), orderID, Pid)
	if err != nil {
		return err
	}

	if OrderDetails.OrderStatus == "Delivered" {
		return errors.New(`the order is delivered .Can't Cancel`)
	}

	if OrderDetails.OrderStatus == "Cancelled" {
		return errors.New(`the order is already cancelled`)
	}

	if OrderDetails.PaymentStatus == "paid" {
		err := repository.ReturnAmountToWallet(uint(userID), orderID, Pid)
		if err != nil {
			return err
		}

	}
	err = repository.UpdateOrderFinalPrice(OrderDetails.OrderID, OrderDetails.TotalPrice)
	if err != nil {
		return err
	}

	err = repository.UpdateStock(pid, OrderDetails.Quantity)
	if err != nil {
		return err
	}

	err = repository.CancelOrder(orderID, Pid, uint(userID))
	if err != nil {
		return err
	}

	return nil
}

func ShipOrders(userID, orderId, pid int) error {
	orderID := strconv.Itoa(orderId)
	Pid := strconv.Itoa(pid)
	err := repository.CheckOrder(orderID, uint(userID))
	fmt.Println(err)
	if err != nil {
		return errors.New(`no orders found with this id`)
	}
	err = repository.CheckSingleOrder(Pid, orderID, uint(userID))
	if err != nil {
		return err
	}
	OrderStatus, err := repository.GetOrderStatus(orderID, Pid)
	fmt.Println(OrderStatus)
	if err != nil {
		return err
	}
	if OrderStatus == "Cancelled" {
		return errors.New("the order is cancelled,cannot ship it")
	}

	if OrderStatus == "Delivered" {
		return errors.New("the order is already delivered")
	}

	if OrderStatus == "Shipped" {
		return errors.New("the order is already Shipped")
	}

	if OrderStatus == "pending" || OrderStatus == "processing" {
		err := repository.ShipOrder(userID, orderId)
		if err != nil {
			return err
		}
		return nil
	}
	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return nil
}

func DeliverOrder(useriD, orderId, pid int) error {
	orderID := strconv.Itoa(orderId)
	Pid := strconv.Itoa(pid)
	err := repository.CheckOrder(orderID, uint(useriD))
	fmt.Println(err)
	if err != nil {
		return errors.New(`no orders found with this id`)
	}
	err = repository.CheckSingleOrder(Pid, orderID, uint(useriD))
	if err != nil {
		return err
	}
	OrderStatus, err := repository.GetOrderStatus(orderID, Pid)
	if err != nil {
		return err
	}
	if OrderStatus == "Cancelled" {
		return errors.New("the order is cancelled,cannot deliver it")
	}

	if OrderStatus == "Delivered" {
		return errors.New("the order is already delivered")
	}

	if OrderStatus == "pending" {
		return errors.New("the order is not shipped yet")
	}

	if OrderStatus == "returned" {
		return errors.New(`the order is returned already by the customer`)
	}

	if OrderStatus == "Shipped" {
		err := repository.DeliverOrder(useriD, orderID)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func ReturnOrder(Token, orderID, pid string) error {
	UserID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}

	err = repository.CheckOrder(orderID, uint(UserID))
	fmt.Println(err)
	if err != nil {
		return errors.New(`no orders found with this id`)
	}

	err = repository.CheckSingleOrder(pid, orderID, uint(UserID))
	if err != nil {
		return err
	}

	Order, err := repository.GetOrderStatus(orderID, pid)
	if err != nil {
		return err
	}

	if Order != "returned" {
		return errors.New(`order is not delivered .Can't return`)
	}

	if Order == "returned" {
		return errors.New(`order is already returned`)
	}

	err = repository.ReturnAmountToWallet(UserID, orderID, pid)
	if err != nil {
		return err
	}

	err = repository.ReturnOrder(UserID, orderID, pid)
	if err != nil {
		return err
	}

	return nil
}

func ExecuteSalesReportByPeriod(period string) (*gofpdf.Fpdf, error) {
	startdate, enddate := utils.CalcualtePeriodDate(period)

	orders, err := repository.GetByDate(startdate, enddate)
	if err != nil {
		return nil, errors.New("report fetching failed")
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Sales Report")
	pdf.Ln(10)
	pdf.Cell(0, 10, "Period:"+period)
	pdf.Ln(10)

	pdf.Cell(0, 10, "Total Sales: "+strconv.FormatFloat(orders.TotalSales, 'f', 2, 64))
	pdf.Ln(10)
	pdf.Cell(0, 10, "Total Orders: "+strconv.Itoa(int(orders.TotalOrders)))
	pdf.Ln(10)
	pdf.Cell(0, 10, "Average Order Price: "+strconv.FormatFloat(orders.AverageOrder, 'f', 2, 64))
	pdf.Ln(10)
	return pdf, nil
}

func ExecuteSalesReportByDate(startdate, enddate time.Time) (*gofpdf.Fpdf, error) {
	orders, err := repository.GetByDate(startdate, enddate)
	if err != nil {
		return nil, errors.New("report fetching failed")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Sales Report")
	pdf.Ln(10)
	pdf.Cell(0, 10, "From date: "+startdate.Format("02-01-2006"))
	pdf.Ln(10)
	pdf.Cell(0, 10, "To date: "+enddate.Format("02-01-2006"))
	pdf.Ln(10)

	pdf.Cell(0, 10, "Total Sales: "+strconv.FormatFloat(orders.TotalSales, 'f', 2, 64))
	pdf.Ln(10)
	pdf.Cell(0, 10, "Total Orders: "+strconv.Itoa(int(orders.TotalOrders)))
	pdf.Ln(10)
	pdf.Cell(0, 10, "Average Order Price: "+strconv.FormatFloat(orders.AverageOrder, 'f', 2, 64))
	pdf.Ln(10)
	return pdf, nil
}

func ExecuteSalesReportByPaymentMethod(startdate, enddate time.Time, paymentmethod string) (*gofpdf.Fpdf, error) {
	var payment string
	if paymentmethod == "1" {
		payment = "Cash On Delivery"
	} else if paymentmethod == "2" {
		payment = "Razorpay"
	} else if paymentmethod == "3" {
		payment = "Wallet"
	} else {
		return nil, errors.New(`payment method doesn't exist`)
	}
	orders, err := repository.GetByPaymentMethod(startdate, enddate, paymentmethod)
	if err != nil {
		return nil, errors.New("report fetching failed")
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40,10, "Sales Report")
	pdf.Ln(10)
	pdf.Cell(0, 10, "Payment Method: "+payment)
	pdf.Ln(10)
	pdf.Cell(0, 10, "From date: "+startdate.Format("02-01-2006"))
	pdf.Ln(10)
	pdf.Cell(0, 10, "To date: "+enddate.Format("02-01-2006"))
	pdf.Ln(10)

	pdf.Cell(0, 10, "Total Sales: "+strconv.FormatFloat(orders.TotalSales, 'f', 2, 64))
	pdf.Ln(10)
	pdf.Cell(0, 10, "Total Orders: "+strconv.Itoa(int(orders.TotalOrders)))
	pdf.Ln(10)
	pdf.Cell(0, 10, "Average Order Price: "+strconv.FormatFloat(orders.AverageOrder, 'f', 2, 64))
	pdf.Ln(10)

	return pdf, nil
}

func PrintInvoice(orderID int, Token string) (*gofpdf.Fpdf, error) {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return nil, err
	}

	orde, err := repository.GetOrderInvoice(orderID, int(userID))
	if err != nil {
		return nil, err
	}

	usr, err := repository.GetUserById(int(userID))
	if err != nil {
		return nil, err
	}

	usadres, err := repository.GetAddressFromOrders(orde.AddressID, int(userID))
	if err != nil {
		return nil, err
	}

	items, err := repository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Invoice")
	pdf.Ln(10)

	pdf.Cell(0, 10, "Customer Name: "+usr.Firstname)
	pdf.Ln(10)
	pdf.Cell(0, 10, "House Name: "+usadres.Housename)
	pdf.Ln(10)
	pdf.Cell(0, 10, "State: "+usadres.State)
	pdf.Ln(10)
	pdf.Cell(0, 10, "Phone: "+usadres.Phone)
	pdf.Ln(10)

	for _, item := range items {
		pdf.Cell(0, 10, "Item: "+item.Name)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Price: "+strconv.FormatFloat(item.Price, 'f', 2, 64))
		pdf.Ln(10)
		pdf.Cell(0, 10, "Quantity: "+strconv.Itoa(item.Stock))
		pdf.Ln(10)
	}
	pdf.Ln(10)
	pdf.Cell(0, 10, "Total Amount: "+strconv.FormatFloat(float64(orde.FinalPrice), 'f', 2, 64))
	pdf.Ln(20)
	pdf.Cell(40, 10, "CityVibe: Thanks for shopping!")
	return pdf, nil
}

func ApplyCoupon(coupon, Token string) error {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}

	err = repository.CheckCouponUsage(userID, coupon)
	if err != nil {
		return err
	}
	DiscountRate, err := repository.GetDiscountRate(coupon)
	if err != nil {
		return err
	}
	_, err = repository.UpdateCartAmount(userID, uint(DiscountRate))
	if err != nil {
		return err
	}
	err = repository.UpdateCouponUsage(userID, coupon)
	if err != nil {
		return err
	}
	err = repository.UpdateCouponCount(coupon)
	if err != nil {
		return err
	}

	return nil
}
