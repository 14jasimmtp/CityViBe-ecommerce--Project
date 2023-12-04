package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"main.go/domain"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func CheckOut(Token, coupon string) (interface{}, error) {
	var FinalPrice float64
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
	if coupon != "" {
		err := repository.CheckCouponUsage(userId, coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		DiscountRate, err := repository.GetDiscountRate(coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		FinalPrice, err = repository.UpdateCartAmount(userId, uint(DiscountRate))
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		err = repository.UpdateCouponUsage(userId, coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		err = repository.UpdateCouponCount(coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}

		fmt.Println(FinalPrice)
		if FinalPrice == 0 {
			FinalPrice = TotalAmount
		}
		return models.CheckOutInfoDiscount{
			Address:        AllUserAddress,
			Cart:           AllCartProducts,
			TotalAmount:    TotalAmount,
			DiscountAmount: FinalPrice,
		}, nil
	}
	return models.CheckOutInfo{
		Address:     AllUserAddress,
		Cart:        AllCartProducts,
		TotalAmount: TotalAmount,
	}, nil
}

func ExecutePurchase(Token string, OrderInput models.CheckOut) (models.OrderSuccessResponse, error) {
	var FinalPrice float64
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

	TotalAmount, err := repository.CartTotalAmount(userId)
	if err != nil {
		return models.OrderSuccessResponse{}, errors.New(`error while calculating total amount`)
	}

	cartItems, err := repository.DisplayCart(userId)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	OrderID, err := repository.OrderFromCart(OrderInput.AddressID, OrderInput.PaymentID, userId, TotalAmount, FinalPrice)
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
	var FinalPrice float64
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

	TotalAmount, err := repository.CartTotalAmount(userId)
	if err != nil {
		return models.OrderSuccessResponse{}, errors.New(`error while calculating total amount`)
	}

	if OrderInput.Coupon != "" {
		err := repository.CheckCouponUsage(userId, OrderInput.Coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		DiscountRate, err := repository.GetDiscountRate(OrderInput.Coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		FinalPrice, err = repository.UpdateCartAmount(userId, uint(DiscountRate))
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		err = repository.UpdateCouponUsage(userId, OrderInput.Coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
		err = repository.UpdateCouponCount(OrderInput.Coupon)
		if err != nil {
			return models.OrderSuccessResponse{}, err
		}
	}
	if FinalPrice == 0 {
		FinalPrice = TotalAmount
	}
	if user.Wallet < FinalPrice {
		return models.OrderSuccessResponse{}, errors.New(`insufficient Balance in Wallet.Add amount to wallet to purchase`)
	}
	cartItems, err := repository.DisplayCart(userId)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	OrderID, err := repository.OrderFromCart(OrderInput.AddressID, OrderInput.PaymentID, userId, TotalAmount, FinalPrice)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	user.Wallet -= FinalPrice

	err = repository.UpdateWallet(user.Wallet, userId)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	if err := repository.AddOrderProducts(userId, OrderID, cartItems); err != nil {
		return models.OrderSuccessResponse{}, err
	}
	err = repository.UpdateShipmentAndPaymentByOrderID("pending", "paid", OrderID)
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
		TotalAmount:   FinalPrice,
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

// func CancelOrderByAdmin(order_id string) error {
// 	err := repository.CheckOrder(order_id,)
// 	fmt.Println(err)
// 	if err != nil {
// 		return errors.New(`no orders found with this id`)
// 	}
// 	orderProduct, err := repository.GetProductDetailsFromOrders(order_id)
// 	if err != nil {
// 		return err
// 	}
// 	err = repository.CancelOrderByAdmin(order_id)
// 	if err != nil {
// 		return err
// 	}
// 	// update the quantity to products since the order is cancelled
// 	err = repository.UpdateStock(orderProduct)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func ShipOrders(userID, orderId, pid string) error {
	OrderStatus, err := repository.GetOrderStatus(orderId, pid)
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

func DeliverOrder(useriD, orderId, pid string) error {

	OrderStatus, err := repository.GetOrderStatus(orderId, pid)
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
		err := repository.DeliverOrder(useriD, orderId)
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

	Order, err := repository.GetOrderStatus(orderID, pid)
	if err != nil {
		return err
	}

	if Order == "pending" || Order == "processing" || Order == "Cancelled" || Order == "Shipped" {
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
