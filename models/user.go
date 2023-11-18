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
	ID        uint   `json:"ID"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Blocked   bool   `json:"blocked"`
}

type UserToken struct {
	User  UserDetailsResponse
	Token string
}

type UserLoginDetails struct {
	Phone    string `json:"phone" binding:"required" validate:"min=10,max=10"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
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
