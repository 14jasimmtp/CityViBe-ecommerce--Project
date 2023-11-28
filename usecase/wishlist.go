package usecase

import (
	"errors"

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

func ViewUserWishlist(Token string) ([]models.UpdateProduct, error) {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return []models.UpdateProduct{}, err
	}
	WishedProducts, err := repository.GetWishlistProducts(userID)
	if err != nil {
		return []models.UpdateProduct{}, err
	}
	return WishedProducts, nil
}

func RemoveProductFromWishlist(pid, Token string) error {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}

	err = repository.CheckExistInWishlist(userID, pid)
	if err == nil {
		return errors.New(`no product found in wishlist with this id`)
	}
	err = repository.RemoveProductFromWishlist(pid, userID)
	if err != nil {
		return err
	}
	return nil
}
