package models

type Admin struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6,max=20"`
	TokenString string `json:"token"`
}
