package repository

import (
	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func AddProduct(product models.Product) (domain.Product, error) {
	var Product domain.Product
	result := initialisers.DB.Exec("INSERT INTO products(name,description,category_id,size,stock,price) values(?,?,?,?,?,?)", product.Name, product.Description, product.Category, product.Size, product.Stock, product.Price).Scan(&Product)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	return Product, nil
}

func EditProductDetails(id string, product domain.Product) (domain.Product, error) {
	var updatedProduct domain.Product

	result := initialisers.DB.Exec("UPDATE products SET stock=stock+$1 WHERE id = $2", product.Stock, id).Scan(&updatedProduct)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	return updatedProduct, nil
}
