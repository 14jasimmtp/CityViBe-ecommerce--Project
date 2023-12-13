package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	"main.go/domain"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func AddProduct(product models.AddProduct, image *multipart.FileHeader) (models.UpdateProduct, error) {
	sess := utils.CreateSession()
	// fmt.Println("sess", sess)

	ImageURL, err := utils.UploadImageToS3(image, sess)
	if err != nil {
		fmt.Println("err:", err)
		return models.UpdateProduct{}, err
	}
	fmt.Println("err:", err)
	product.ImageURL = ImageURL
	fmt.Println("image,", ImageURL)

	ProductResponse, err := repository.AddProduct(product)
	if err != nil {
		return models.UpdateProduct{}, err
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

func EditProductDetails(id string, product models.AddProduct) (models.UpdateProduct, error) {
	UpdatedProduct, err := repository.EditProductDetails(id, product)
	if err != nil {
		return models.UpdateProduct{}, err
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

func ShowProductsByCategory() ([]domain.Product, error) {

	return []domain.Product{}, nil
}

func SeeAllProducts() ([]domain.Product, error) {
	products, err := repository.SeeAllProducts()
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func GetSingleProduct(id string) (models.Product, error) {
	product, err := repository.GetSingleProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func FilterProductCategoryWise(category string) ([]models.Product, error) {
	products, err := repository.FilterProductCategoryWise(category)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func SearchProduct(search string) ([]models.Product, error) {
	products, err := repository.SearchProduct(search)
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func FilterProducts(category, size string, minPrice, maxPrice float64) ([]models.UpdateProduct, error) {
	filteredProducts, err := repository.FilterProducts(category, size, minPrice, maxPrice)
	if err != nil {
		return []models.UpdateProduct{}, err
	}
	return filteredProducts, nil
}

func ExecuteAddOffer(offer *models.Offer) error {
	err := repository.CreateOffer(offer)
	if err != nil {
		return errors.New("error creating offer")
	} else {
		return nil
	}
}
