package repository

import (
	"errors"

	initialisers "main.go/Initialisers"
	"main.go/models"
)

func AddToCart(pid int, userid uint) error {
	query := initialisers.DB.Exec(`INSERT INTO carts (user_id,product_id,quantity) VALUES (?,?,?)`, userid, pid, 1)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func DisplayCart(userid uint) ([]models.Cart, error) {

	var count int
	if err := initialisers.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ? ", userid).First(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var Cart []models.Cart

	if err := initialisers.DB.Raw("SELECT carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.price FROM carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userid).First(&Cart).Error; err != nil {
		return []models.Cart{}, err
	}

	return Cart, nil
}

func RemoveProductFromCart(pid int, userid uint) error {
	query := initialisers.DB.Exec(`DELETE FROM carts WHERE product_id = ? AND user_id = ?`, pid, userid)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func CheckProductExistInCart(userId uint, pid string) error {
	query := initialisers.DB.Raw(`SELECT * FROM carts WHERE user_id = ? AND product_id = ?`, userId, pid)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected < 1 {
		return errors.New(`product already exist in cart`)
	}

	return nil
}

func UpdateQuantity(userid uint, pid, quantity string) ([]models.Cart, error) {
	query := initialisers.DB.Raw(`UPDATE carts SET quantity = ? WHERE user_id = ? AND product_id = ?`, quantity, userid, pid)
	if query.Error != nil {
		return []models.Cart{}, query.Error
	}

	var Cart []models.Cart
	if err := initialisers.DB.Raw("SELECT carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.price FROM carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userid).First(&Cart).Error; err != nil {
		return []models.Cart{}, err
	}

	return Cart, nil
}
