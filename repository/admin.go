package repository

import (
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

func GetAllOrderDetailsBrief() ([]models.CombinedOrderDetails, error) {

	var orderDatails []models.CombinedOrderDetails
	err := initialisers.DB.Raw("SELECT orders.order_id,orders.final_price,orders.shipment_status,orders.payment_status,users.firstname,users.email,users.phone,addresses.house_name,addresses.street,addresses.city,addresses.state,addresses.pin FROM orders INNER JOIN users ON orders.user_id = users.id INNER JOIN addresses ON orders.address_id = addresses.id ").Scan(&orderDatails).Error
	if err != nil {
		return []models.CombinedOrderDetails{}, nil
	}
	return orderDatails, nil

}
