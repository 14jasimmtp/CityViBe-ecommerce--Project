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

func AddAddress(Address models.Address, UserId uint) (models.AddressRes, error) {
	var AddressRes models.AddressRes
	query := initialisers.DB.Raw(`INSERT INTO addresses(user_id,name,house_name,street,city,state,pin) VALUES (?,?,?,?,?,?,?)`, UserId, Address.Name, Address.HouseName, Address.Street, Address.City, Address.State, Address.Pin).Scan(&AddressRes)
	if query.Error != nil {
		return models.AddressRes{}, query.Error
	}
	return AddressRes, nil
}

func EditAddress(Address models.Address, UserId int) (models.AddressRes, error) {
	var AddressRes models.AddressRes
	query := initialisers.DB.Exec(`UPDATE addresses SET name = ?,house_name = ?,street = ?,city = ?,state = ?,pin=?`, Address.Name, Address.HouseName, Address.Street, Address.State, Address.Pin).Scan(&AddressRes)
	if query.Error != nil {
		return models.AddressRes{}, query.Error
	}
	return AddressRes, nil
}

func ViewAddress(id uint) ([]models.AddressRes, error) {
	var Address []models.AddressRes
	query := initialisers.DB.Raw(`SELECT * FROM addresses WHERE user_id = ?`, id).Scan(&Address)
	if query.Error != nil {
		return []models.AddressRes{}, query.Error
	}

	if query.RowsAffected < 1 {
		return []models.AddressRes{}, errors.New("no address found. add new address")
	}

	return Address, nil
}

func DeleteAddress(AddressId int, UserId int) error {
	query := initialisers.DB.Exec(`DELETE FROM addresses WHERE id = ? && user_id = ?`, AddressId, UserId)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
