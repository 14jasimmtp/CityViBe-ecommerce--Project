package usecase

import (
	"errors"
	"strconv"

	"main.go/domain"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func CheckOut(Token, coupon string) (models.CheckOutInfo, error) {
	var DiscountAmount float64
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
		DiscountRate, err := repository.GetDiscountRate(coupon)
		if err != nil {
			return models.CheckOutInfo{}, err
		}
		DiscountAmount = TotalAmount - (TotalAmount * (DiscountRate / 100))

		return models.CheckOutInfo{
			Address:        AllUserAddress,
			Cart:           AllCartProducts,
			TotalAmount:    TotalAmount,
			DiscountAmount: DiscountAmount,
		}, nil
	}
	return models.CheckOutInfo{
		Address:     AllUserAddress,
		Cart:        AllCartProducts,
		TotalAmount: TotalAmount,
	}, nil
}

func OrderFromCart(Token string, AddressId uint, PaymentID uint) (models.OrderSuccessResponse, error) {
	userId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	addressExist := repository.CheckAddressExist(userId, AddressId)
	if !addressExist {
		return models.OrderSuccessResponse{}, errors.New(`address doesn't exist`)
	}

	paymentExist := repository.CheckPaymentMethodExist(PaymentID)
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

	cartItems, err := repository.DisplayCart(userId)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	OrderID, err := repository.OrderFromCart(AddressId, PaymentID, userId, TotalAmount)
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}

	if err := repository.AddOrderProducts(userId,OrderID, cartItems); err != nil {
		return models.OrderSuccessResponse{}, err
	}

	body, err := repository.GetOrder(OrderID)
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
		PaymentStatus: body.PaymentStatus,
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

	err = repository.CancelOrder(orderId,pid, UserID)
	if err != nil {
		return err
	}

	return nil

}

// func CancelOrderByAdmin(order_id string) error {
// 	err := repository.CheckOrder(order_id)
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

func ShipOrders(orderId string) error {
	OrderStatus, err := repository.GetOrderStatus(orderId)
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

	if OrderStatus == "pending" {
		err := repository.ShipOrder(orderId)
		if err != nil {
			return err
		}
		return nil
	}
	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return nil
}

func DeliverOrder(orderId string) error {

	OrderStatus, err := repository.GetOrderStatus(orderId)
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

	if OrderStatus == "Shipped" {
		err := repository.DeliverOrder(orderId)
		if err != nil {
			return err
		}
		return nil
	}
	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return nil
}

func CancelSingleProduct(pid, Token, orderID string) ([]models.OrderProducts, error) {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return []models.OrderProducts{}, err
	}

	err = repository.CheckSingleOrder(pid, orderID, userID)
	if err != nil {
		return []models.OrderProducts{}, errors.New(`no orders found with this id`)
	}

	OrderStatus, err := repository.GetOrderStatus(orderID)
	if err != nil {
		return []models.OrderProducts{}, err
	}

	if OrderStatus == "Delivered" {
		return []models.OrderProducts{}, errors.New(`the order is delivered .Can't Cancel`)
	}

	if OrderStatus == "Cancelled" {
		return []models.OrderProducts{}, errors.New(`the order is already cancelled`)
	}

	err = repository.CancelSingleOrder(pid, orderID, userID)
	if err != nil {
		return []models.OrderProducts{}, err
	}

	err = repository.UpdateSingleStock(pid)
	if err != nil {
		return []models.OrderProducts{}, err
	}

	err = repository.UpdateAmount(orderID, userID)
	if err != nil {
		return []models.OrderProducts{}, err
	}

	orderDetails, err := repository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return []models.OrderProducts{}, err
	}

	return orderDetails, nil

}

func ReturnOrder() {

}
