package usecase

import (
	"errors"
	"strconv"

	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func ViewCart(Token string) ([]models.Cart, error) {
	UserId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return []models.Cart{}, err
	}

	Cart, err := repository.DisplayCart(UserId)
	if err != nil {
		return []models.Cart{}, err
	}

	return Cart, nil

}

func AddToCart(pid, Token string) error {

	_, err := repository.GetSingleProduct(pid)
	if err != nil {
		return errors.New("no products exist with this id")
	}

	ProId, err := strconv.Atoi(pid)
	if err != nil {
		return err
	}

	err = repository.CheckStock(ProId)
	if err != nil {
		return errors.New("product out of stock")
	}

	UserId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}

	err = repository.AddToCart(ProId, UserId)
	if err != nil {
		return err
	}

	return nil
}

func RemoveProductsFromCart(pid, Token string) error {
	ProId,err:=strconv.Atoi(pid)
	if err != nil{
		return err
	}

	UserId,err:=utils.ExtractUserIdFromToken(Token)
	if err != nil{
		return err
	}

	err=repository.RemoveProductFromCart(ProId,UserId)
	if err != nil{
		return err
	}

	return nil
}
