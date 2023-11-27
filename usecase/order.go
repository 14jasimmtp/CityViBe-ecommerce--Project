package usecase

import (
	"errors"
	"fmt"

	"main.go/domain"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func CheckOut(Token string) (models.CheckOutInfo, error) {
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

	return models.CheckOutInfo{
		Address:     AllUserAddress,
		Cart:        AllCartProducts,
		TotalAmount: TotalAmount,
	}, nil
}

func OrderFromCart(Token string, AddressId uint) (domain.Order, error) {
	userId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return domain.Order{}, err
	}

	addressExist := repository.CheckAddressExist(userId, AddressId)
	if !addressExist {
		return domain.Order{}, errors.New(`address doesn't exist`)
	}

	cartExist := repository.CheckCartExist(userId)
	if !cartExist {
		return domain.Order{}, errors.New(`cart is empty`)

	}

	TotalAmount, err := repository.CartTotalAmount(userId)
	if err != nil {
		return domain.Order{}, errors.New(`error while calculating total amount`)
	}

	cartItems, err := repository.DisplayCart(userId)
	if err != nil {
		return domain.Order{}, err
	}

	OrderID, err := repository.OrderFromCart(AddressId, userId, TotalAmount)
	if err != nil {
		return domain.Order{}, err
	}

	if err := repository.AddOrderProducts(OrderID, cartItems); err != nil {
		return domain.Order{}, err
	}

	body, err := repository.GetOrder(OrderID)
	if err != nil {
		return domain.Order{}, err
	}

	var orderItemDetails domain.OrderItem
	for _, c := range cartItems {
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = c.Quantity
		err := repository.UpdateCartAndStockAfterOrder(userId, int(orderItemDetails.ProductID), orderItemDetails.Quantity)
		if err != nil {
			return domain.Order{}, err
		}
	}
	return body, nil

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

func CancelOrder(Token, orderId string) error {
	UserID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}
	err = repository.CheckOrder(orderId)
	if err != nil {
		return errors.New(`no orders found with this id`)
	}

	orderDetails, err := repository.GetProductDetailsFromOrders(orderId)
	if err != nil {
		return err
	}

	OrderStatus, err := repository.GetOrderStatus(orderId)
	if err != nil {
		return err
	}

	if OrderStatus == "Delivered" {
		return errors.New(`the order is delivered .Can't Cancel`)
	}

	if OrderStatus == "Cancelled" {
		return errors.New(`the order is already cancelled`)
	}

	err = repository.CancelOrder(orderId, UserID)
	if err != nil {
		return err
	}

	err = repository.UpdateStock(orderDetails)
	if err != nil {
		return err
	}

	return nil

}

func CancelOrderByAdmin(order_id string) error {
	err := repository.CheckOrder(order_id)
	fmt.Println(err)
	if err != nil {
		return errors.New(`no orders found with this id`)
	}
	orderProduct, err := repository.GetProductDetailsFromOrders(order_id)
	if err != nil {
		return err
	}
	err = repository.CancelOrderByAdmin(order_id)
	if err != nil {
		return err
	}
	// update the quantity to products since the order is cancelled
	err = repository.UpdateStock(orderProduct)
	if err != nil {
		return err
	}
	return nil
}

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

	// err = repository.updateAmount(orderID)
	// if err !=nil{
	// 	return []models.OrderProducts{},err
	// }

	orderDetails, err := repository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return []models.OrderProducts{}, err
	}


	return orderDetails, nil

}
