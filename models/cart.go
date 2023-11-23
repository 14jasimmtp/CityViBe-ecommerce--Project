package models


type Cart struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    float64 `json:"quantity"`
	Price  float64 `json:"price"`
}

type CartResponse struct {
	TotalPrice float64
	Cart       []Cart
}
