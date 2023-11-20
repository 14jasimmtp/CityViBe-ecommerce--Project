package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func SignUp(User models.UserSignUpDetails) error {
	CheckEmail, err := repository.CheckUserExistsEmail(User.Email)
	if err != nil {
		fmt.Println("server error")
		return errors.New("server error")
	}
	if CheckEmail != nil {
		fmt.Println("user already exist")
		return errors.New("user already exist with this email")
	}

	CheckPhone, err := repository.CheckUserExistsByPhone(User.Phone)
	if err != nil {
		fmt.Println("server error")
		return errors.New("server error")
	}
	if CheckPhone != nil {
		fmt.Println("user already exist with this number")
		return errors.New("user already exist with this number")
	}

	if User.Password != User.ConfirmPassword {
		fmt.Println("passwords doesn't match")
		return errors.New("paswords doesn't match")
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), 10)
	if err != nil {
		fmt.Println("error while hashing ")
		return errors.New("server error occured(password hashing)")
	}
	User.Password = string(HashedPassword)

	sentOtp := utils.SendOtp(User.Phone)
	if sentOtp != nil {
		fmt.Println("error gen otp")
		return errors.New("error occured generating otp")
	}
	var Userdt domain.User
	err = copier.Copy(&Userdt, &User)
	if err != nil {
		return err
	}
	initialisers.DB.Create(&Userdt)
	return nil
}

func UserLogin(user models.UserLoginDetails) error {
	CheckPhone, err := repository.CheckUserExistsByPhone(user.Phone)
	if err != nil {
		return errors.New("error with server")
	}
	if CheckPhone == nil {
		return errors.New("phone number doesn't exist")
	}
	userdetails, err := repository.FindUserByPhone(user.Phone)
	fmt.Println(userdetails, user.Password)
	if err != nil {
		return err
	}

	if userdetails.Blocked {
		return errors.New("user is blocked")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userdetails.Password), []byte(user.Password))
	if err != nil {
		return errors.New("password not matching")
	}
	sentOtp := utils.SendOtp(user.Phone)
	if sentOtp != nil {
		fmt.Println("error gen otp")
		return errors.New("error occured generating otp")
	}

	return nil
}

func ForgotPassword(phone string) error {
	user, err := repository.FindUserByPhone(phone)
	if err != nil {
		return errors.New("user doesn't found with this number")
	}

	if user.Blocked {
		return errors.New("user is blocked")
	}
	err = utils.SendOtp(phone)
	if err != nil {
		return errors.New("error generating otp ")
	}

	return nil
}

func ResetForgottenPassword(Newpassword models.ForgotPassword) error {
	err := utils.CheckOtp(Newpassword.Phone, Newpassword.OTP)
	if err != nil {
		return err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(Newpassword.NewPassword), 10)
	if err != nil {
		return errors.New("error while hashing password")
	}
	Newpassword.NewPassword = string(hashed)

	err = repository.ChangePassword(Newpassword)
	if err != nil {
		return err
	}

	return nil
}
