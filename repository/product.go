package repository

import (
	"errors"

	"github.com/jinzhu/copier"
	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func AddProduct(product models.Product) (domain.Product, error) {
	var Product domain.Product
	result := initialisers.DB.Raw("INSERT INTO products(name,description,category_id,size,stock,price) values(?,?,?,?,?,?)", product.Name, product.Description, product.CategoryId, product.Size, product.Stock, product.Price).Scan(&Product)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	return Product, nil
}

func EditProductDetails(id string, product models.Product) (models.Product, error) {
	var updatedProduct domain.Product

	result := initialisers.DB.Raw("UPDATE products SET id=?,name=?,description=?,category_id=?,size=?,stock=?,price=?,color=? WHERE id=?", product.ID, product.Name, product.Description, product.CategoryId, product.Size, product.Stock, product.Price, product.Color, id).Scan(&updatedProduct)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	copier.Copy(&product, &updatedProduct)
	return product, nil
}

func DeleteProduct(id int) error {
	query := initialisers.DB.Exec(`UPDATE products SET deleted = true WHERE id = ?`, id)
	if query.Error != nil {
		return errors.New("no product found to delete")
	}
	return nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	query := initialisers.DB.Raw(`SELECT * FROM products`).Scan(&products)
	if query.Error != nil {
		return []models.Product{}, query.Error
	}
	return products, nil
}

func SeeAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := initialisers.DB.Raw("SELECT * FROM products WHERE deleted = false").Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
