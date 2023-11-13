package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func AdminLogin(admin models.Admin) (models.Admin, error) {
	AdminDetails, err := repository.AdminLogin(admin)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Admin doesn't exist")
		return models.Admin{}, errors.New("admin not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(AdminDetails.Password), []byte(admin.Password)) != nil {
		fmt.Println("wrong password")
		return models.Admin{}, errors.New("wrong password")
	}

	tokenString, err := utils.AdminTokenGenerate(AdminDetails, "admin")
	if err != nil {
		fmt.Println("error generating token")
		return models.Admin{}, errors.New("error generating token")
	}

	return models.Admin{
		Firstname:   AdminDetails.Firstname,
		Lastname:    admin.Lastname,
		TokenString: tokenString,
	}, nil

}

func GetAllUsers() ([]models.UserDetailsResponse, error) {
	Users, err := repository.GetAllUsers()
	if err != nil {
		return []models.UserDetailsResponse{}, err
	}
	return Users, nil
}

func BlockUser(idStr string) error {
	id, _ := strconv.Atoi(idStr)
	user, err := repository.GetUserById(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}
	err = repository.BlockUserByID(user)
	if err != nil {
		return err
	}
	return nil

}

func UnBlockUser(idStr string) error {
	id, _ := strconv.Atoi(idStr)
	user, err := repository.GetUserById(id)
	if err != nil {
		return err
	}
	if !user.Blocked {
		return errors.New("already unblocked")
	} else {
		user.Blocked = false
	}
	err = repository.UnBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil

}
