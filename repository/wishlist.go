package repository

import (
	"errors"

	initialisers "main.go/Initialisers"
	"main.go/models"
)

func CheckExistInWishlist(userID uint, pid string) error {
	var product models.Product
	query := initialisers.DB.Raw(`SELECT * FROM wishlists WHERE user_id = ? AND product_id = ?`, userID, pid).Scan(&product)
	if query.Error != nil {
		return errors.New(`something got wrong`)
	}
	if query.RowsAffected > 0 {
		return errors.New(`product already exist in wishlist`)
	}
	return nil
}

func AddProductToWishlist(pid string, userID uint) error {
	query := initialisers.DB.Raw(`INSERT INTO wishlists(user_id,product_id) VALUES(?,?)`, userID, pid)
	if query.Error != nil {
		return errors.New(`something got wrong`)
	}
	return nil
}

func GetWishlistProducts(userID uint) ([]models.Product, error) {
	var Products []models.Product
	query := initialisers.DB.Raw(
		`SELECT products.name,products.description,categories.category,sizes.size,products.stock,products.price,products.color
		 FROM products 
		 INNER JOIN wishlists ON products.id=wishlists.product_id
		 INNER JOIN categories ON products.category_id=categories.id
		 INNER JOIN sizes ON products.size_id=sizes.id
		 WHERE wishlists.user_id = ?`, userID,
	).Scan(&Products)
	if query.Error != nil {
		return []models.Product{}, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return []models.Product{}, errors.New(`no products in wishlist`)
	}
	return Products, nil
}
