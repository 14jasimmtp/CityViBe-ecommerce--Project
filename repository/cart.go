package repository

import (
	"fmt"

	initialisers "main.go/Initialisers"
	"main.go/models"
)

func AddToCart(pid int, userid uint, productAmount float64) error {
	query := initialisers.DB.Exec(`INSERT INTO carts (user_id,product_id,quantity,price) VALUES (?,?,?,?)`, userid, pid, 1, productAmount)
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

func CheckProductExistInCart(userId uint, pid string) (bool, error) {
	var count int
	query := initialisers.DB.Raw(`SELECT COUNT(*) FROM carts WHERE user_id = ? AND product_id = ?`, userId, pid).Scan(&count)
	if query.Error != nil {
		return false, query.Error
	}
	fmt.Println(count)

	if count > 0 {
		return true, nil
	}

	return false, nil
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

func CartTotalAmount(userid uint) (float64, error) {
	var Amount float64
	err := initialisers.DB.Raw(`SELECT SUM(price) FROM carts WHERE user_id = ?`, userid).Scan(&Amount).Error

	if err != nil {
		return 0.0, nil
	}
	return Amount, nil
}

func CheckCartExist(userID uint) bool {
	var count int
	if err := initialisers.DB.Raw("SELECT COUNT(*) FROM carts WHERE  user_id = ?", userID).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func UpdateCart(quantity int, price float64, userID uint, product_id string) error {

	if err := initialisers.DB.Exec("update carts set quantity = quantity + $1, price = $2 where user_id = $3 and product_id = $4", quantity, price, userID, product_id).Error; err != nil {
		return err
	}

	return nil

}

func TotalPrizeOfProductInCart(userID uint, productID string) (float64, error) {

	var totalPrice float64
	if err := initialisers.DB.Raw("select sum(price) as total_price from carts where user_id = ? and product_id = ?", userID, productID).Scan(&totalPrice).Error; err != nil {
		return 0.0, err
	}
	return totalPrice, nil
}

func UpdateQuantityAdd(id uint, prdt_id string) error {
	err := initialisers.DB.Exec("UPDATE Carts SET quantity = quantity + 1 WHERE user_id=$1 AND product_id = $2 ", id, prdt_id).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTotalPrice(id uint, product_id string) error {
	err := initialisers.DB.Exec("UPDATE carts SET price = carts.quantity * products.price FROM products  WHERE carts.product_id = products.id AND carts.user_id = $1 AND carts.product_id = $2", id, product_id).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateQuantityless(id uint, prdt_id string) error {
	err := initialisers.DB.Exec("UPDATE Carts SET quantity = quantity - 1 WHERE user_id=$1 AND product_id = $2 ", id, prdt_id).Error
	if err != nil {
		return err
	}
	return nil
}

func CartExist(userID uint) (bool, error) {
	var count int
	if err := initialisers.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ? ", userID).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil

}

func EmptyCart(userID uint) error {

	if err := initialisers.DB.Exec("DELETE FROM carts WHERE user_id = ? ", userID).Error; err != nil {
		return err
	}

	return nil

}
