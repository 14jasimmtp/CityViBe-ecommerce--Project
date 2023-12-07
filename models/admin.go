package models

type Admin struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6,max=20"`
	TokenString string `json:"token"`
}

type AdminOrder struct {
	UserID    int `json:"user_id" validate:"required,number"`
	OrderID   int `json:"order_id" validate:"required,number"`
	ProductID int `json:"product_id" validate:"required,number"`
}

type SalesReport struct {
	TotalSales int
	
}
