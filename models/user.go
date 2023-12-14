package models

type UserSignUpDetails struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname" binding:"required"`
	Email           string `json:"email" binding:"required" validate:"email"`
	Phone           string `json:"phone" binding:"required" validate:"min=10,max=10"`
	Password        string `json:"password" binding:"required" validate:"min=6,max=20"`
	ConfirmPassword string `json:"confirm password" binding:"required"`
}

type UserDetailsResponse struct {
	ID        uint    `json:"ID"`
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Blocked   bool    `json:"blocked"`
	Wallet    float64 `json:"wallet"`
}

type UserToken struct {
	User  UserDetailsResponse
	Token string
}

type UserLoginDetails struct {
	Phone    string `json:"phone" validate:"required,len=10,number"`
	Password string `json:"password"  validate:"required,min=6,max=20"`
}

type UserLoginResponse struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ClientToken struct {
	Email string
	Phone string
	ID    uint
	Role  string
}

type Phone struct {
	Phone string `json:"phone" validate:"required,max=10,min=10"`
}
type ForgotPassword struct {
	Phone       string `json:"phone"  validate:"required,max=10,min=10"`
	OTP         string `json:"otp" validate:"required"`
	NewPassword string `json:"new password"  validate:"required,min=6,max=20"`
}

type Address struct {
	Name      string `json:"name" validate:"required"`
	Housename string `json:"house_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type AddressRes struct {
	ID         int    `json:"id"`
	Name       string `json:"name" validate:"required"`
	House_name string `json:"house_name" validate:"required"`
	Phone      string `json:"phone"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state" validate:"required"`
	Pin        string `json:"pin" validate:"required"`
}

type UserProfile struct {
	Firstname string  `json:"firstname" validate:"required"`
	Lastname  string  `json:"lastname" validate:"required"`
	Email     string  `json:"email" validate:"required"`
	Phone     string  `json:"phone" validate:"required,min=10,max=10"`
	Wallet    float64 `json:"wallet"`
}

type ChangePassword struct {
	Oldpassword     string `json:"Current password" validate:"required"`
	Newpassword     string `json:"New password" validate:"required"`
	ConfirmPassword string `json:"Confirm new password" validate:"required"`
}

type Errors struct {
	Error string
}
