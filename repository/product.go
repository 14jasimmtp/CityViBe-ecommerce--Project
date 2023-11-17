package repository

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func AddProduct(product models.AddProduct) (models.AddProduct, error) {
	var dproduct domain.Product
	result := initialisers.DB.Raw("INSERT INTO products(name,description,category_id,size,stock,price,color) values(?,?,?,?,?,?,?)", product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price, product.Color).Scan(&dproduct)
	if result.Error != nil {
		return models.AddProduct{}, result.Error
	}
	var Product models.AddProduct
	fmt.Println(Product)
	copier.Copy(&Product, &dproduct)
	return Product, nil
}

func EditProductDetails(id string, product models.AddProduct) (models.AddProduct, error) {
	var updatedProduct domain.Product

	result := initialisers.DB.Raw("UPDATE products SET name=?,description=?,category_id=?,size=?,stock=?,price=?,color=? WHERE id=?", product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price, product.Color, id).Scan(&updatedProduct)
	if result.Error != nil {
		return models.AddProduct{}, result.Error
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
	query := initialisers.DB.Raw(`SELECT name,description,categories.category,size,color,price FROM products INNER JOIN categories ON categories.id = products.category_id WHERE deleted = false`).Scan(&products)
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
