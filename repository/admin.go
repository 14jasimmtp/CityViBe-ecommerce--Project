package repository

import (
	"errors"
	"fmt"
	"time"

	initialisers "main.go/Initialisers"
	"main.go/models"
	"main.go/utils"
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

func DashBoardUserDetails() (models.DashBoardUser, error) {
	var userDetails models.DashBoardUser
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	err = initialisers.DB.Raw("SELECT COUNT(*) FROM users WHERE blocked=true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	return userDetails, nil
}

func DashBoardProductDetails() (models.DashBoardProduct, error) {
	var productDetails models.DashBoardProduct
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	err = initialisers.DB.Raw("SELECT COUNT(*) FROM products WHERE stock<=0").Scan(&productDetails.OutofStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	return productDetails, nil
}

func TotalRevenue() (models.DashboardRevenue, error) {
	var revenueDetails models.DashboardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := initialisers.DB.Raw("SELECT COALESCE(SUM(total_price),0) FROM orders WHERE payment_status = 'paid' AND created_at >=? AND created_at <=?", startTime, endTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = utils.CalcualtePeriodDate("monthly")
	err = initialisers.DB.Raw("SELECT COALESCE (SUM(total_price),0) FROM orders WHERE payment_status = 'paid' AND created_at >=? AND created_at <=?", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = utils.CalcualtePeriodDate("yearly")
	err = initialisers.DB.Raw("SELECT COALESCE (SUM(total_price),0) FROM orders WHERE payment_status = 'paid' AND created_at >=? AND created_at <=?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	return revenueDetails, nil
}

func AmountDetails() (models.DashboardAmount, error) {
	var amountDetails models.DashboardAmount
	err := initialisers.DB.Raw("SELECT COALESCE (SUM(total_price),0) FROM orders WHERE payment_status = 'paid' ").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}
	err = initialisers.DB.Raw("SELECT COALESCE(SUM(total_price),0) FROM orders WHERE payment_status = 'not paid' ").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}
	return amountDetails, nil
}
