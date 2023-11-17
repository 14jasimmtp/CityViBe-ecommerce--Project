package usecase

import (
	"strconv"

	"main.go/domain"
	"main.go/models"
	"main.go/repository"
)

func AddProduct(product models.AddProduct) (domain.Product, error) {
	ProductResponse, err := repository.AddProduct(product)
	if err != nil {
		return domain.Product{}, err
	}
	return ProductResponse, nil
}

func GetAllProducts() ([]models.Product, error) {
	ProductDetails, err := repository.GetAllProducts()
	if err != nil {
		return []models.Product{}, err
	}
	return ProductDetails, nil
}

func EditProductDetails(id string, product models.AddProduct) (domain.Product, error) {
	UpdatedProduct, err := repository.EditProductDetails(id, product)
	if err != nil {
		return domain.Product{}, err
	}
	return UpdatedProduct, nil
}

func DeleteProduct(id string) error {
	idnum, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = repository.DeleteProduct(idnum)
	if err != nil {
		return err
	}

	return nil
}

func ShowProductsByCategory() ([]models.Product, error) {

	return []models.Product{}, nil
}



func SeeAllProducts() ([]domain.Product, error) {
	products, err := repository.SeeAllProducts()
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}
