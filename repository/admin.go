package repository

import (
	"errors"
	"fmt"

	initialisers "main.go/Initialisers"
	"main.go/models"
)

func AdminLogin(adminDetails models.Admin) (models.Admin, error) {
	var details models.Admin
	if err := initialisers.DB.Raw("SELECT * FROM admins WHERE email=?", adminDetails.Email).Scan(&details).Error; err != nil {
		return models.Admin{}, err
	}
	return details, nil
}

func GetAllUsers() ([]models.UserDetailsResponse, error) {
	var users []models.UserDetailsResponse
	result := initialisers.DB.Raw("SELECT id,email,firstname,lastname,phone,blocked FROM users").Scan(&users)
	if result.Error != nil {
		fmt.Println("data fetching error")
		return []models.UserDetailsResponse{}, result.Error
	}
	return users, nil
}

func BlockUserByID(user models.UserDetailsResponse) error {
	result := initialisers.DB.Exec("UPDATE users SET blocked = ? WHERE id = ?", user.Blocked, user.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UnBlockUserByID(user models.UserDetailsResponse) error {
	result := initialisers.DB.Exec("UPDATE users SET blocked = ? WHERE id = ?", user.Blocked, user.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllOrderDetailsBrief() ([]models.ViewAdminOrderDetails, error) {

	var orderDatails []models.AdminOrderDetails
	query := initialisers.DB.Raw("SELECT orders.user_id,orders.id, total_price as final_price, payment_methods.payment_mode AS payment_method, payment_status FROM orders INNER JOIN payment_methods ON orders.payment_method_id=payment_methods.id  ORDER BY orders.id DESC").Scan(&orderDatails)
	if query.Error != nil {
		return []models.ViewAdminOrderDetails{}, errors.New(`something went wrong`)
	}
	var fullOrderDetails []models.ViewAdminOrderDetails
	for _, ok := range orderDatails {
		var OrderProductDetails []models.OrderProductDetails
		initialisers.DB.Raw("SELECT order_items.product_id,products.name AS product_name,order_items.order_status,order_items.quantity,order_items.total_price FROM order_items INNER JOIN products ON order_items.product_id = products.id WHERE order_items.order_id = ? ORDER BY order_id DESC", ok.Id).Scan(&OrderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.ViewAdminOrderDetails{OrderDetails: ok, OrderProductDetails: OrderProductDetails})
	}
	return fullOrderDetails, nil

}

func GetSingleOrderDetails(orderID string) ([]models.OrderProductDetails, error) {
	var Order []models.OrderProductDetails
	query := initialisers.DB.Raw(`SELECT product_id,products.name AS product_name,order_status,quantity,Total_price FROM order_items INNER JOIN products ON product_id=products.id WHERE order_id = ?`, orderID).Scan(&Order)
	if query.Error != nil {
		return []models.OrderProductDetails{}, query.Error
	}
	return Order, nil
}
