package models

type Admin struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	TokenString string `json:"token"`
}
