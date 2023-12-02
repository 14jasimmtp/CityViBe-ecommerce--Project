package repository

import (
	"errors"

	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func OrderFromCart(addressid uint,paymentid, userid uint, price float64) (int, error) {
	var id int
	query := `
    INSERT INTO orders (created_at , user_id , address_id ,payment_method_id, final_price)
    VALUES (NOW(),?, ?, ?,?)
    RETURNING id`
	initialisers.DB.Raw(query, userid, addressid,paymentid, price).Scan(&id)
	return id, nil
}

func AddAmountToOrder(Amount float64, orderID uint) error {
	err := initialisers.DB.Exec("UPDATE orders SET final_price = ? WHERE id = ?", Amount, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func AddOrderProducts(order_id int, cart []models.Cart) error {
	query := `
    INSERT INTO order_items (order_id,product_id,quantity,total_price)
    VALUES (?, ?, ?, ?) `
	for _, v := range cart {
		var productID int
		if err := initialisers.DB.Raw("SELECT id FROM products WHERE name = $1", v.ProductName).Scan(&productID).Error; err != nil {
			return err
		}
		if err := initialisers.DB.Exec(query, order_id, productID, v.Quantity, v.Price).Error; err != nil {
			return err
		}
	}
	return nil

}

func CheckPaymentMethodExist(paymentid uint)bool{
	query:=initialisers.DB.Raw(`SELECT * FROM payment_methods WHERE id = ?`,paymentid)
	return query.RowsAffected < 1
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
	initialisers.DB.Raw("SELECT id, final_price, order_status,payment_method, payment_status FROM orders WHERE user_id = ? ", userID).Scan(&orderDatails)
	var fullOrderDetails []models.ViewOrderDetails
	for _, ok := range orderDatails {
		var OrderProductDetails []models.OrderProductDetails
		initialisers.DB.Raw("SELECT order_items.product_id,products.name AS product_name,order_items.quantity,order_items.total_price FROM order_items INNER JOIN products ON order_items.product_id = products.id WHERE order_items.order_id = ?", ok.Id).Scan(&OrderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.ViewOrderDetails{OrderDetails: ok, OrderProductDetails: OrderProductDetails})
	}
	return fullOrderDetails, nil

}

func CheckOrder(orderid string) error {
	var count int
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM orders WHERE id = ?", orderid).Scan(&count).Error
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
	if err := initialisers.DB.Raw("SELECT product_id,quantity FROM order_items WHERE id = ?", order_id).Scan(&OrderProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}
	return OrderProductDetails, nil
}

func GetOrderStatus(orderId string) (string, error) {
	var status string
	err := initialisers.DB.Raw("SELECT order_status FROM orders WHERE id= ?", orderId).Scan(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

func CancelOrder(order_id string, userID uint) error {
	status := "Cancelled"
	err := initialisers.DB.Exec("UPDATE orders SET order_status = ? ,payment_status = 'refunded', approval='false' WHERE id = ? AND user_id = ?", status, order_id, userID).Error
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

func UpdateSingleStock(pid string) error {
	var quantity int
	if err := initialisers.DB.Raw("SELECT stock FROM products WHERE id = ?", pid).Scan(&quantity).Error; err != nil {
		return err
	}
	quantity = quantity + 1
	if err := initialisers.DB.Exec("UPDATE products SET stock  = ? WHERE id = ?", quantity, pid).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCartAndStockAfterOrder(userID uint, productID int, quantity float64) error {
	err := initialisers.DB.Exec("DELETE FROM carts WHERE user_id = ? and product_id = ?", userID, productID).Error
	if err != nil {
		return err
	}

	err = initialisers.DB.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", quantity, productID).Error
	if err != nil {
		return err
	}

	return nil
}

func CheckSingleOrder(pid, orderId string, userId uint) error {
	var count int
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM order_items WHERE product_id = ? AND order_id =? AND user_id = ?", pid, orderId, userId).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New(`no orders found`)
	}
	return nil
}

func CancelSingleOrder(pid, orderId string, userId uint) error {
	err := initialisers.DB.Exec("DELETE FROM order_items WHERE product_id = ? AND order_id = ? AND user_id = ? ", pid, orderId, userId).Error
	if err != nil {
		return err
	}
	return nil
}

func CancelOrderByAdmin(orderID string) error {
	status := "Cancelled"
	err := initialisers.DB.Exec("UPDATE orders SET order_status = ? ,payment_status = 'refunded', approval='false' WHERE id = ? ", status, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func ShipOrder(orderId string) error {
	err := initialisers.DB.Exec("UPDATE orders SET order_status = 'Shipped' , approval = 'true' WHERE id = ?", orderId).Error
	if err != nil {
		return err
	}
	return nil
}

func DeliverOrder(orderId string) error {
	status := "Delivered"
	err := initialisers.DB.Exec("UPDATE orders SET order_status = ? ,payment_status = 'paid', approval='false' WHERE id = ? ", status, orderId).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAmount(oid string, userID uint) error {
	var Amount float64
	query := initialisers.DB.Raw(`SELECT SUM(total_price) FROM order_items WHERE order_id = ? AND user_id = ?`, oid, userID).Scan(&Amount)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	query = initialisers.DB.Exec(`UPDATE FROM orders SET final_price = final_price - ? WHERE id = ?`, Amount, oid)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}
