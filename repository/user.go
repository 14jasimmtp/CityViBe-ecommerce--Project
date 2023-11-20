package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func CheckUserExistsEmail(email string) (*domain.User, error) {
	var user domain.User
	result := initialisers.DB.Where(&domain.User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil

}

func CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	result := initialisers.DB.Where(&domain.User{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func SignUpUser(user models.UserSignUpDetails) (*models.UserDetailsResponse, error) {
	var User models.UserDetailsResponse

	result := initialisers.DB.Raw("INSERT INTO users(firstname,lastname,email,phone,password) VALUES(?,?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Phone, user.Password).Scan(&User)
	if result.Error != nil {
		return nil, result.Error
	}
	return &User, nil
}

func FindUserByPhone(phone string) (*domain.User, error) {
	var user domain.User
	result := initialisers.DB.Raw("SELECT * FROM users WHERE phone = ?", phone).Scan(&user)
	if result.Error != nil {
		return &domain.User{}, result.Error
	}
	return &user, nil
}

func GetUserById(id int) (models.UserDetailsResponse, error) {
	var user models.UserDetailsResponse

	result := initialisers.DB.Raw("SELECT * FROM users WHERE id = ? ", id).Scan(&user)
	if result.Error != nil {
		fmt.Println("error fetching user")
		return models.UserDetailsResponse{}, result.Error
	}
	return user, nil
}

func ChangePassword(ResetUser models.ForgotPassword) error {
	query := initialisers.DB.Exec(`UPDATE users SET password = ? WHERE phone = ?`, ResetUser.NewPassword, ResetUser.Phone)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
