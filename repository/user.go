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

func UpdateAddress(userid uint, aid string, Address models.Address) (models.AddressRes, error) {
	var AddressRes models.AddressRes
	query := initialisers.DB.Exec(`UPDATE addresses SET name = ?,phone = ?,house_name = ?,street = ?,city = ?,state = ?,pin=? WHERE id = ? AND user_id = ?`, Address.Name, Address.Phone, Address.HouseName, Address.Street, Address.State, Address.Pin, aid, userid).Scan(&AddressRes)
	if query.Error != nil {
		return models.AddressRes{}, query.Error
	}
	if query.RowsAffected < 1 {
		return models.AddressRes{}, errors.New(`no address found to update with this id`)
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

func RemoveAddress(Userid uint, aid string) error {
	query := initialisers.DB.Exec(`DELETE FROM addresses WHERE id = ? && user_id = ?`, aid, Userid)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func UserProfile(userid uint) (models.UserProfile, error) {
	var User models.UserProfile
	query := initialisers.DB.Raw(`SELECT * FROM users WHERE id = ?`, userid).Scan(&User)
	if query.Error != nil {
		return models.UserProfile{}, query.Error
	}

	if query.RowsAffected < 1 {
		return models.UserProfile{}, errors.New(`no user profile found`)
	}

	return User, nil
}

func UpdateUserProfile(userid uint, user models.UserProfile) (models.UserProfile, error) {
	var UpdatedUser models.UserProfile
	query := initialisers.DB.Exec(`UPDATE users SET firstname = ?,lastname = ?,email = ?,phone = ? WHERE id = ?`, user.FirstName, user.LastName, user.Email, user.Phone, userid).Scan(&UpdatedUser)
	if query.Error != nil {
		return models.UserProfile{}, query.Error
	}

	return UpdatedUser, nil
}

func CheckAddressExist(userid uint,address string) bool{
	var count int
	if err := initialisers.DB.Raw("SELECT COUNT(*) FROM addresses WHERE id = ? AND user_id = ?", address, userid).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// func UpdateUserEmail(email string, userID int) error {
// 	err := initialisers.DB.Exec("UPDATE users SET email= ? WHERE id = ?", email, userID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdateUserPhone(phone string, userID int) error {
// 	if err := initialisers.DB.Exec("UPDATE users SET phone = ? WHERE id = ?", phone, userID).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdateFirstName(name string, userID int) error {

// 	err := initialisers.DB.Exec("update users set firstname = ? where id = ?", name, userID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
// func UpdateLastName(name string, userID int) error {

// 	err := initialisers.DB.Exec("update users set lastname = ? where id = ?", name, userID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
// func CheckAddressAvailabilityWithAddressID(addressID, userID int) bool {
// 	var count int
// 	if err := initialisers.DB.Raw("SELECT COUNT(*) FROM addresses WHERE id = ? AND user_id = ?", addressID, userID).Scan(&count).Error; err != nil {
// 		return false
// 	}
// 	return count > 0
// }
// func UpdateName(name string, addressID int) error {
// 	err := initialisers.DB.Exec("UPDATE addresses SET name= ? WHERE id = ?", name, addressID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdateHouseName(HouseName string, addressID int) error {
// 	err := initialisers.DB.Exec("UPDATE addresses SET house_name= ? WHERE id = ?", HouseName, addressID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdateStreet(street string, addressID int) error {
// 	err := initialisers.DB.Exec("UPDATE addresses SET street= ? WHERE id = ?", street, addressID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdateCity(city string, addressID int) error {
// 	err := initialisers.DB.Exec("UPDATE addresses SET city= ? WHERE id = ?", city, addressID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdateState(state string, addressID int) error {
// 	err := initialisers.DB.Exec("UPDATE addresses SET state= ? WHERE id = ?", state, addressID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdatePin(pin string, addressID int) error {
// 	err := initialisers.DB.Exec("UPDATE addresses SET pin= ? WHERE id = ?", pin, addressID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func AddressDetails(addressID int) (models.AddressInfoResponse, error) {
// 	var addressDetails models.AddressInfoResponse
// 	err := initialisers.DB.Raw("SELECT a.id, a.name, a.house_name, a.street, a.city, a.state, a.pin FROM addresses a WHERE a.id = ?", addressID).Row().Scan(&addressDetails.ID, &addressDetails.Name, &addressDetails.HouseName, &addressDetails.Street, &addressDetails.City, &addressDetails.State, &addressDetails.Pin)
// 	if err != nil {
// 		return models.AddressInfoResponse{}, err
// 	}
// 	return addressDetails, nil
// }

// func ChangePassword(id int, password string) error {
// 	err := initialisers.DB.Exec("UPDATE users SET password = $1 WHERE id = $2", password, id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func GetPassword(id int) (string, error) {
// 	var userPassword string
// 	err := initialisers.DB.Raw("SELECT password FROM users WHERE id = ?", id).Scan(&userPassword).Error
// 	if err != nil {
// 		return "", err
// 	}
// 	return userPassword, nil
// }
// func UpdateQuantityAdd(id, prdt_id int) error {
// 	err := initialisers.DB.Exec("UPDATE Carts SET quantity = quantity + 1 WHERE user_id=$1 AND product_id = $2 ", id, prdt_id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateTotalPrice(id, product_id int) error {
// 	err := initialisers.DB.Exec("UPDATE carts SET total_price = carts.quantity * products.price FROM products  WHERE carts.product_id = products.id AND carts.user_id = $1 AND carts.product_id = $2", id, product_id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateQuantityless(id, prdt_id int) error {
// 	err := initialisers.DB.Exec("UPDATE Carts SET quantity = quantity - 1 WHERE user_id=$1 AND product_id = $2 ", id, prdt_id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func FindUserByMobileNumber(phone string) bool {

// 	var count int
// 	if err := initialisers.DB.Raw("SELECT count(*) FROM users WHERE phone = ?", phone).Scan(&count).Error; err != nil {
// 		return false
// 	}

// 	return count > 0

// }
// func FindIdFromPhone(phone string) (int, error) {
// 	var id int
// 	if err := initialisers.DB.Raw("SELECT id FROM users WHERE phone=?", phone).Scan(&id).Error; err != nil {
// 		return id, err
// 	}
// 	return id, nil
// }
// func AddressExistInUserProfile(addressID, userID int) (bool, error) {
// 	var count int
// 	err := initialisers.DB.Raw("SELECT COUNT (*) FROM addresses WHERE user_id = $1 AND id = $2", userID, addressID).Scan(&count).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }
// func RemoveFromUserProfile(userID, addressID int) error {
// 	err := initialisers.DB.Exec("DELETE FROM addresses WHERE user_id = ? AND  id= ?", userID, addressID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
