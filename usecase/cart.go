package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func ViewCart(Token string) (models.CartResponse, error) {
	UserId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.CartResponse{}, err
	}

	Cart, err := repository.DisplayCart(UserId)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := repository.CartTotalAmount(UserId)
	if err != nil {
		return models.CartResponse{}, err
	}

	return models.CartResponse{
		TotalPrice: cartTotal,
		Cart:       Cart,
	}, nil

}

func AddToCart(pid, Token string) (models.CartResponse, error) {

	_, err := repository.GetSingleProduct(pid)
	if err != nil {
		return models.CartResponse{}, errors.New("product doesn't exist")
	}

	UserId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.CartResponse{}, err
	}

	ProId, err := strconv.Atoi(pid)
	if err != nil {
		return models.CartResponse{}, err
	}

	productPrize, err := repository.GetProductAmountFromID(pid)
	if err != nil {
		return models.CartResponse{}, err
	}
	true, err := repository.CheckProductExistInCart(UserId, pid)
	if err != nil {
		return models.CartResponse{}, err
	}
	fmt.Println(true)
	if true {
		TotalProductAmount, err := repository.TotalPrizeOfProductInCart(UserId, pid)
		if err != nil {
			return models.CartResponse{}, err
		}

		err = repository.UpdateCart(1, TotalProductAmount+productPrize, UserId, pid)
		if err != nil {
			return models.CartResponse{}, err
		}
	} else {
		err := repository.AddToCart(ProId, UserId, productPrize)
		if err != nil {
			return models.CartResponse{}, err
		}
	}

	CartDetails, err := repository.DisplayCart(UserId)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotalAmount, err := repository.CartTotalAmount(UserId)
	if err != nil {
		return models.CartResponse{}, err
	}

	return models.CartResponse{
		TotalPrice: cartTotalAmount,
		Cart:       CartDetails,
	}, nil
}

func RemoveProductsFromCart(pid, Token string) (models.CartResponse, error) {
	ProId, err := strconv.Atoi(pid)
	if err != nil {
		return models.CartResponse{}, err
	}

	UserId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.CartResponse{}, err
	}

	err = repository.RemoveProductFromCart(ProId, UserId)
	if err != nil {
		return models.CartResponse{}, err
	}

	updatedCart, err := repository.DisplayCart(UserId)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := repository.CartTotalAmount(UserId)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		TotalPrice: cartTotal,
		Cart:       updatedCart,
	}, nil
}

func UpdateQuantityFromCart(Token, pid, quantity string) ([]models.Cart, error) {
	UserId, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return []models.Cart{}, err
	}

	updatedCart, err := repository.UpdateQuantity(UserId, pid, quantity)
	if err != nil {
		return []models.Cart{}, err
	}

	return updatedCart, nil
}

func UpdateQuantityIncrease(Token, pid string) error {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}

	check, err := repository.CheckProductExistInCart(userID, pid)
	if err != nil {
		return err
	}
	if !check {
		return errors.New(`no products found in cart with this id`)
	}

	err = repository.UpdateQuantityAdd(userID, pid)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePriceAdd(Token, pid string) error {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}
	err = repository.UpdateTotalPrice(userID, pid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateQuantityDecrease(Token, pid string) error {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}

	check, err := repository.CheckProductExistInCart(userID, pid)
	if err != nil {
		return err
	}

	if !check {
		return errors.New(`no products found in cart with this id`)
	}

	quantity, err := repository.ProductQuantityCart(userID, pid)
	if err != nil {
		return err
	}

	if quantity == 1 {
		return errors.New(`quantity is 1 .can't reduce anymore`)
	}

	err = repository.UpdateQuantityless(userID, pid)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePriceDecrease(Token, pid string) error {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return err
	}
	err = repository.UpdateTotalPrice(userID, pid)
	if err != nil {
		return err
	}
	return nil
}

func EraseCart(Token string) (models.CartResponse, error) {
	userID, err := utils.ExtractUserIdFromToken(Token)
	if err != nil {
		return models.CartResponse{}, err
	}
	ok, err := repository.CartExist(userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("cart already empty")
	}
	if err := repository.EmptyCart(userID); err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := repository.CartTotalAmount(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		TotalPrice: cartTotal,
		Cart:       []models.Cart{},
	}

	return cartResponse, nil
}
