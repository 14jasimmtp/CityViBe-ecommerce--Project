package repository

import (
	"errors"

	initialisers "main.go/Initialisers"
	"main.go/models"
)

func GetPaymentDetails(orderID int) (models.Payment, error) {
	var Paymentdt models.Payment
	query := initialisers.DB.Raw(`SELECT users.firstname,orders.final_price,users.phone FROM orders INNER JOIN users ON orders.user_id=users.id WHERE orders.id = ? `, orderID).Scan(&Paymentdt)
	if query.Error != nil {
		return models.Payment{}, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return models.Payment{}, errors.New(`no orders foun with this id `)
	}
	return Paymentdt, nil
}

func PaymentAlreadyPaid(orderID int) (bool, error) {
	var paymentStatus string
	query := initialisers.DB.Raw(`SELECT payment_status from orders where id = ? `, orderID).Scan(&paymentStatus)
	if query.Error != nil {
		return false, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return false, errors.New(`no orders found with this id `)
	}
	if paymentStatus == `paid` {
		return true, nil
	}
	return false, nil
}

func PayMethod(orderID int) (int, error) {
	var id int
	query := initialisers.DB.Raw(`SELECT payment_method_id FROM orders WHERE id = ?`, orderID).Scan(&id)
	if query.Error != nil {
		return 0, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return 0, errors.New(`no orders foun with this id `)
	}
	return id, nil
}

func AddRazorPayDetails(orderID int, RazorID string) error {
	err := initialisers.DB.Exec("INSERT INTO razor_pays (order_id,razor_id) VALUES (?,?)", orderID, RazorID).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckPaymentStatus(orderID int) (string, error) {
	var paymentStatus string
	err := initialisers.DB.Raw(`SELECT payment_status FROM orders WHERE id = $1`, orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}
	return paymentStatus, nil
}
func UpdatePaymentDetails(orderID int, paymentID string) error {
	err := initialisers.DB.Exec("UPDATE razor_pays set payment_id = ? WHERE order_id= ?", paymentID, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateShipmentAndPaymentByOrderID(orderStatus string, paymentStatus string, orderID int) error {
	err := initialisers.DB.Exec("UPDATE orders SET payment_status = ?  WHERE id = ?", paymentStatus, orderID).Error
	if err != nil {
		return err
	}
	err=initialisers.DB.Exec("UPDATE order_items SET order_status = ?  WHERE order_id = ?", orderStatus, orderID).Error
	if err != nil{
		return errors.New(`something went wrong`)
	}
	return nil
}

