package models

type Cart struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	FinalPrice  float64 `json:"-"`
}

type CartResponse struct {
	TotalPrice float64
	Cart       []Cart
}
