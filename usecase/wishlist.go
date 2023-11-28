package usecase

import (
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func AddProductToWishlist(pid string, Token string) error {

	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}
	_, err = repository.GetSingleProduct(pid)
	if err != nil {
		return err
	}

	err = repository.CheckExistInWishlist(userID, pid)
	if err != nil {
		return err
	}
	err = repository.AddProductToWishlist(pid, userID)
	if err != nil {
		return err
	}
	return nil
}

func ViewUserWishlist(Token string) ([]models.Product, error) {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return []models.Product{}, err
	}
	WishedProducts, err := repository.GetWishlistProducts(userID)
	if err != nil {
		return []models.Product{}, err
	}
	return WishedProducts, nil
}
