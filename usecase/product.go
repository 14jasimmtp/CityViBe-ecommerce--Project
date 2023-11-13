package usecase

import (
	"main.go/domain"
	"main.go/models"
	"main.go/repository"
)

func AddProduct(product models.Product) (domain.Product, error) {
	ProductResponse, err := repository.AddProduct(product)
	if err != nil {
		return domain.Product{}, err
	}
	return ProductResponse, nil
}


