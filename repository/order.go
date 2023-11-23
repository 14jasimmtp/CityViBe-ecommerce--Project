package repository

import (
	"errors"

	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func OrderFromCart(cartid, addressid string, userid uint) (domain.Order, error) {
	var Order domain.Order
	err := initialisers.DB.Exec("INSERT INTO orders (created_at, user_id, address_id, cart_id) SELECT NOW(), c.user_id, a.id, c.id FROM carts c JOIN addresses a ON c.user_id = a.user_id WHERE a.id = ? AND c.id = ? AND c.user_id = ? ", cartid, addressid, userid).Scan(&Order).Error
	if err != nil {
		return domain.Order{}, err
	}

	return Order, nil
}

func AddAmountToOrder(Amount float64, orderID uint) error {
	err := initialisers.DB.Exec("UPDATE orders SET final_price = ? WHERE id = ?", Amount, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrder(orderID int) (domain.Order, error) {
	var order domain.Order
	err := initialisers.DB.Raw("SELECT * FROM orders WHERE id = ?", orderID).Scan(&order).Error
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func GetOrderDetails(userID uint) ([]models.ViewOrderDetails, error) {
	var orderDatails []models.OrderDetails
	initialisers.DB.Raw("SELECT order_id, final_price, order_status, payment_status FROM orders WHERE user_id = ? ", userID).Scan(&orderDatails)
	var fullOrderDetails []models.ViewOrderDetails
	for _, ok := range orderDatails {
		var OrderProductDetails []models.OrderProductDetails
		initialisers.DB.Raw("select order_items.product_id,products.name as product_name,order_items.quantity,order_items.total_price from order_items inner join products on order_items.product_id = products.id where order_items.order_id = ?", ok.OrderId).Scan(&OrderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.ViewOrderDetails{OrderDetails: ok, OrderProductDetails: OrderProductDetails})
	}
	return fullOrderDetails, nil

}

func CheckOrder(orderid string) error {
	var count int
	err := initialisers.DB.Raw("SELECT * FROM orders WHERE order_id = ?", orderid).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New(`no orders found`)
	}
	return nil
}

func GetProductDetailsFromOrders(order_id string) ([]models.OrderProducts, error) {
	var OrderProductDetails []models.OrderProducts
	if err := initialisers.DB.Raw("SELECT product_id,quantity FROM order_items WHERE order_id = ?", order_id).Scan(&OrderProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}
	return OrderProductDetails, nil
}

func GetOrderStatus(orderId string) (string, error) {
	var status string
	err := initialisers.DB.Raw("SELECT order_status FROM orders WHERE order_id= ?", orderId).Scan(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

func CancelOrder(order_id string,userID uint) error {
	status := "cancelled"
	err := initialisers.DB.Exec("UPDATE orders SET order_status = ? ,payment_status = refunded, approval='false' WHERE order_id = ? AND user_id = ?", status, order_id,userID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateStock(orderProducts []models.OrderProducts) error {
	for _, ok := range orderProducts {
		var quantity int
		if err := initialisers.DB.Raw("SELECT stock FROM products WHERE id = ?", ok.ProductId).Scan(&quantity).Error; err != nil {
			return err
		}
		ok.Stock += quantity
		if err := initialisers.DB.Exec("UPDATE products SET stock  = ? WHERE id = ?", ok.Stock, ok.ProductId).Error; err != nil {
			return err
		}
	}
	return nil
}

func CancelOrderByAdmin(orderID string) error{
	status := "cancelled"
	err := initialisers.DB.Exec("UPDATE orders SET order_status = ? ,payment_status = refunded, approval='false' WHERE order_id = ? ", status, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func ShipOrder(orderId string) error{
	err := initialisers.DB.Exec("UPDATE orders SET order_status = 'Shipped' , approval = 'true' WHERE order_id = ?", orderId).Error
	if err != nil {
		return err
	}
	return nil
}