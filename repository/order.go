package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func OrderFromCart(addressid uint, paymentid, userid uint, price, finalprice float64) (int, error) {
	var id int
	query := `
    INSERT INTO orders (created_at , user_id , address_id ,payment_method_id,total_price,final_price)
    VALUES (NOW(),?, ?, ?,?,?)
    RETURNING id`
	initialisers.DB.Raw(query, userid, addressid, paymentid, price, finalprice).Scan(&id)
	return id, nil
}

func AddAmountToOrder(Amount float64, orderID uint) error {
	err := initialisers.DB.Exec("UPDATE orders SET final_price = ? WHERE id = ?", Amount, orderID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func AddOrderProducts(userID uint, orderid int, cart []models.Cart) error {
	query := `
    INSERT INTO order_items (order_id,product_id,user_id,quantity,total_price)
    VALUES (?, ?, ?, ?, ?) `
	for _, v := range cart {
		var productID int
		if err := initialisers.DB.Raw("SELECT id FROM products WHERE name = $1", v.ProductName).Scan(&productID).Error; err != nil {
			return errors.New(`something went wrong`)
		}
		if err := initialisers.DB.Exec(query, orderid, productID, userID, v.Quantity, v.Price).Error; err != nil {
			return errors.New(`something went wrong`)
		}
	}
	return nil

}

func CheckPaymentMethodExist(paymentid uint) bool {
	query := initialisers.DB.Raw(`SELECT * FROM payment_methods WHERE id = ?`, paymentid)
	return query.RowsAffected < 1
}
func GetOrder(orderID int) (domain.Order, error) {
	var order domain.Order
	err := initialisers.DB.Raw("SELECT * FROM orders WHERE id = ?", orderID).Scan(&order).Error
	if err != nil {
		return domain.Order{}, errors.New(`something went wrong`)
	}
	return order, nil
}

func GetOrderDetails(userID uint) ([]models.ViewOrderDetails, error) {
	var orderDatails []models.OrderDetails
	query := initialisers.DB.Raw("SELECT orders.id, total_price as final_price, payment_methods.payment_mode AS payment_method, payment_status FROM orders INNER JOIN payment_methods ON orders.payment_method_id=payment_methods.id WHERE user_id = ? ORDER BY orders.id DESC", userID).Scan(&orderDatails)
	if query.Error != nil {
		return []models.ViewOrderDetails{}, errors.New(`something went wrong`)
	}
	var fullOrderDetails []models.ViewOrderDetails
	for _, ok := range orderDatails {
		var OrderProductDetails []models.OrderProductDetails
		initialisers.DB.Raw("SELECT order_items.product_id,products.name AS product_name,order_items.order_status,order_items.quantity,order_items.total_price FROM order_items INNER JOIN products ON order_items.product_id = products.id WHERE order_items.order_id = ? ORDER BY order_id DESC", ok.Id).Scan(&OrderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.ViewOrderDetails{OrderDetails: ok, OrderProductDetails: OrderProductDetails})
	}
	return fullOrderDetails, nil

}

func CheckOrder(orderid string, userID uint) error {
	var count int
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM order_items WHERE order_id = ? AND user_id = ?", orderid, userID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New(`no orders found`)
	}
	return nil
}

func GetProductDetailsFromOrders(orderid int) ([]models.Product, error) {
	var OrderProductDetails []models.Product
	if err := initialisers.DB.Raw("SELECT product_id,products.name,products.description,order_items.quantity as stock,order_items.total_price as price FROM order_items INNER JOIN products ON order_items.product_id = products.id WHERE order_items.order_id = ?", orderid).Scan(&OrderProductDetails).Error; err != nil {
		return []models.Product{}, err
	}
	return OrderProductDetails, nil
}

func GetOrderStatus(orderId, pid string) (string, error) {
	var status struct {
		OrderStatus string `json:"order_status"`
	}
	err := initialisers.DB.Raw("SELECT order_status FROM order_items WHERE order_id = ? AND product_id = ? ", orderId, pid).Scan(&status.OrderStatus).Error
	if err != nil {
		return "", errors.New(`something went wrong`)
	}
	return status.OrderStatus, nil
}

func CancelOrder(orderid, pid string, userID uint) error {
	status := "Cancelled"
	err := initialisers.DB.Exec("UPDATE order_items SET order_status = ?  WHERE order_id = ? AND product_id = ? AND user_id = ?", status, orderid, pid, userID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func UpdateStock(pid int, quantity int) error {
	query := initialisers.DB.Exec(`UPDATE products SET stock = stock + $1 WHERE id = $2`, quantity, pid)
	if query.Error != nil {
		return errors.New(`something went wrong`)
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
		return errors.New(`something went wrong`)
	}

	err = initialisers.DB.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", quantity, productID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}

	return nil
}

func CheckSingleOrder(pid, orderId string, userId uint) error {
	var count int
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM order_items WHERE product_id = ? AND order_id =? AND user_id = ?", pid, orderId, userId).Scan(&count).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	if count < 1 {
		return errors.New(`no product found with this id`)
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
		return errors.New(`something went wrong`)
	}
	return nil
}

func ShipOrder(userID, orderId int) error {
	err := initialisers.DB.Exec("UPDATE order_items SET order_status = 'Shipped'  WHERE order_id = ? AND user_id = ?", orderId, userID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func DeliverOrder(userID int, orderId string) error {
	status := "Delivered"
	err := initialisers.DB.Exec("UPDATE order_items SET order_status = ? WHERE order_id = ? AND user_id = ?", status, orderId, userID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	err = initialisers.DB.Exec("UPDATE orders SET payment_status = 'paid' WHERE id = ? ", orderId).Error
	if err != nil {
		return errors.New(`something went wrong`)
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

// func ReturnAmountToWallet(userID uint, orderID, pid string) error {
// 	var amount float64
// 	query := initialisers.DB.Raw(`SELECT total_price FROM order_items WHERE product_id = ? AND order_id = ? AND user_id = ?`, pid, orderID, userID).Scan(&amount)
// 	if query.Error != nil {
// 		return errors.New(`something went wrong`)
// 	}

// 	query = initialisers.DB.Exec(`UPDATE users SET wallet = wallet + $1 WHERE id = $2`, amount, userID)
// 	if query.Error != nil {
// 		return errors.New(`something went wrong`)
// 	}
// 	if query.RowsAffected == 0 {
// 		return errors.New(`no orders found with this id`)
// 	}
// 	return nil
// }

func ReturnAmountToWallet(userID uint, orderID, pid string) error {
	tx := initialisers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var amount float64
	query := tx.Raw(`SELECT total_price FROM order_items WHERE product_id = ? AND order_id = ? AND user_id = ?`, pid, orderID, userID).Scan(&amount)
	if query.Error != nil {
		tx.Rollback()
		return errors.New(`something went wrong`)
	}
	query = tx.Exec(`UPDATE users SET wallet = wallet + $1 WHERE id = $2`, amount, userID)
	if query.Error != nil {

		tx.Rollback()
		return errors.New(`something went wrong`)
	}

	if query.RowsAffected == 0 {

		tx.Rollback()
		return errors.New(`no orders found with this id`)
	}

	tx.Commit()

	return nil
}

func CancelOrderDetails(userID uint, orderID, pid string) (models.CancelDetails, error) {
	var Details models.CancelDetails
	query := initialisers.DB.Raw(`SELECT order_status,quantity,orders.payment_status,order_items.total_price,order_id FROM order_items INNER JOIN orders ON orders.id =order_items.order_id WHERE order_items.order_id = ? AND order_items.user_id = ? AND order_items.product_id = ?`, orderID, userID, pid).Scan(&Details)
	if query.Error != nil {
		return models.CancelDetails{}, errors.New(`something went wrong`)
	}
	return Details, nil
}

func UpdateOrderFinalPrice(orderID int, amount float64) error {
	query := initialisers.DB.Exec(`UPDATE orders SET final_price = final_price - $1 WHERE id = $2`, amount, orderID)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func UpdateCartAmount(userID, discount uint) (float64, error) {
	var finalprice float64
	perc := 1 - (float64(discount) / 100)
	query := initialisers.DB.Exec(`UPDATE carts SET final_price = price * ? WHERE user_id = ?`, perc, userID)
	if query.Error != nil {
		return 0.0, errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return 0.0, errors.New(`something went wrong`)
	}
	query = initialisers.DB.Raw(`SELECT SUM(final_price) FROM carts WHERE user_id = ?`, userID).Scan(&finalprice)
	if query.Error != nil {
		return 0.0, errors.New(`something went wrong`)
	}
	return finalprice, nil
}

func ReturnOrder(userID uint, orderID, pid string) error {
	query := initialisers.DB.Exec(`UPDATE order_items SET order_status = 'returned' WHERE user_id = ? AND product_id = ? AND order_id = ?`, userID, pid, orderID)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return errors.New(`no order with this id found to return`)
	}
	return nil
}

func GetOrderInvoice(orderID, UserID int) (domain.Order, error) {
	var order domain.Order
	query := initialisers.DB.Raw(`SELECT * FROM orders WHERE id = ? AND user_id = ?`, orderID, UserID).Scan(&order)
	if query.Error != nil {
		return domain.Order{}, errors.New(`something went wrong`)
	}

	if query.RowsAffected == 0 {
		return domain.Order{}, errors.New(`no orders found`)
	}
	return order, nil
}

func GetByDate(startdate, enddate time.Time) (*models.SalesReport, error) {
	var err error
	var order []domain.Order
	var report models.SalesReport
	enddate = enddate.Add(+24 * time.Hour)

	if err := initialisers.DB.Model(&order).Where("created_at BETWEEN ? AND ? ", startdate, enddate).Select("SUM(total_price) as total_sales").Scan(&report).Error; err != nil {
		return nil, err
	}
	if err := initialisers.DB.Model(&order).Where("created_at BETWEEN ? AND ? ", startdate, enddate).Count(&report.TotalOrders).Error; err != nil {
		return nil, err
	}
	if err := initialisers.DB.Model(&order).Where("created_at BETWEEN ? AND ? ", startdate, enddate).Select("AVG(total_price) as average_order").Scan(&report).Error; err != nil {
		return nil, err
	}

	formattedValue := fmt.Sprintf("%.2f", report.AverageOrder)
	fmt.Println(formattedValue)

	report.AverageOrder, err = strconv.ParseFloat(formattedValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &report, nil

}
func GetByPaymentMethod(startdate, enddate time.Time, paymentmethod string) (*models.SalesReport, error) {
	var err error
	var order []domain.Order
	enddate = enddate.Add(+24 * time.Hour)
	var report models.SalesReport

	if err := initialisers.DB.Model(&order).Where("created_at BETWEEN ? AND ? AND payment_method_id=?", startdate, enddate, paymentmethod).Select("SUM(total_price) as total_sales").Scan(&report).Error; err != nil {
		return nil, err
	}
	if err := initialisers.DB.Model(&order).Where("created_at BETWEEN ? AND ? AND payment_method_id=?", startdate, enddate, paymentmethod).Count(&report.TotalOrders).Error; err != nil {
		return nil, err
	}
	if err := initialisers.DB.Model(&order).Where("created_at BETWEEN ? AND ? AND payment_method_id=?", startdate, enddate, paymentmethod).Select("AVG(total_price) as average_order").Scan(&report).Error; err != nil {
		return nil, err
	}

	formattedValue := fmt.Sprintf("%.2f", report.AverageOrder)
	fmt.Println(formattedValue)

	report.AverageOrder, err = strconv.ParseFloat(formattedValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &report, nil
}

func GetAddressFromOrders(address_id, userID int) (models.Address, error) {
	var Address models.Address
	query := initialisers.DB.Raw(`SELECT name,house_name as housename,phone,state FROM addresses WHERE user_id = ? AND id = ? `, userID, address_id).Scan(&Address)
	if query.Error != nil {
		return models.Address{}, errors.New(`something went wrong`)
	}

	if query.RowsAffected == 0 {
		return models.Address{}, errors.New(`no orders found`)
	}

	return Address, nil

}

func DashBoardOrder() (models.DashboardOrder, error) {
	var orderDetail models.DashboardOrder
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM order_items WHERE order_status= 'Delivered'").Scan(&orderDetail.DeliveredOrderProducts).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}
	err = initialisers.DB.Raw("SELECT COUNT(*) FROM order_items WHERE order_status='pending' OR order_status = 'processing'").Scan(&orderDetail.PendingOrderProducts).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}
	err = initialisers.DB.Raw("select count(*) from order_items where order_status = 'Cancelled' OR order_status = 'returned'").Scan(&orderDetail.CancelledOrderProducts).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = initialisers.DB.Raw("select count(*) from order_items").Scan(&orderDetail.TotalOrderItems).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = initialisers.DB.Raw("select COALESCE(SUM(quantity),0) from order_items").Scan(&orderDetail.TotalOrderQuantity).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}
	return orderDetail, nil
}

func CheckVerifiedPayment(orderID int) (bool, error) {
	var payment string
	query := initialisers.DB.Raw(`SELECT payment_status FROM orders WHERE id = ?`, orderID).Scan(&payment)
	if query.Error != nil {
		return false, errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return false, errors.New(`no orders found`)
	}
	if payment == "paid" {
		return true, nil
	}

	return false, nil
}
