package repository

import (
	"errors"
	"fmt"

	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func AddProduct(product models.AddProduct) (models.UpdateProduct, error) {
	var dproduct models.UpdateProduct
	var p domain.Product
	result := initialisers.DB.Raw("INSERT INTO products(name,description,category_id,size_id,stock,price,color) values(?,?,?,?,?,?,?)", product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price, product.Color).Scan(&p)
	fmt.Println(p)
	if result.Error != nil {
		return models.UpdateProduct{}, result.Error
	}
	query := initialisers.DB.Raw(`SELECT products.id,name,description,categories.category,sizes.size,stock,price,color FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE name = ?`, product.Name).Scan(&dproduct)
	if query.Error != nil {
		return models.UpdateProduct{}, query.Error
	}
	fmt.Println(dproduct)
	return dproduct, nil
}

func EditProductDetails(id string, product models.AddProduct) (models.UpdateProduct, error) {
	var updatedProduct models.UpdateProduct

	result := initialisers.DB.Raw("UPDATE products SET name=?,description=?,category_id=?,size_id=?,stock=?,price=?,color=? WHERE id=?", product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price, product.Color, id).Scan(&updatedProduct)
	if result.Error != nil {
		return models.UpdateProduct{}, result.Error
	}
	query := initialisers.DB.Raw(`SELECT products.id,name,description,categories.category,sizes.size,stock,color,price FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE products.id = ?`, id).Scan(&updatedProduct)
	if query.Error != nil {
		return models.UpdateProduct{}, query.Error

	}
	return updatedProduct, nil
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
	query := initialisers.DB.Raw(`SELECT name,description,categories.category,sizes.size,stock,color,price FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE deleted = false`).Scan(&products)
	if query.Error != nil {
		return []models.Product{}, query.Error
	}
	return products, nil
}

func SeeAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := initialisers.DB.Raw("SELECT * FROM products ").Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
