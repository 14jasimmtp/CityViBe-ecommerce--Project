package models

type CheckOutInfo struct {
	Address     []AddressRes `json:"address"`
	Cart        []Cart       `json:"cart"`
	TotalAmount float64      `json:"Total Amount"`
}

type OrderResponse struct{

}

type ViewOrderDetails struct{
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}

type OrderDetails struct {
	OrderId        string
	FinalPrice     float64
	ShipmentStatus string
	PaymentStatus  string
}

type OrderProductDetails struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type OrderProducts struct {
	ProductId string `json:"product_id"`
	Stock     int    `json:"stock"`
}

type CombinedOrderDetails struct {
	OrderId        string  `json:"order_id"`
	FinalPrice     float64 `json:"final_price"`
	ShipmentStatus string  `json:"shipment_status"`
	PaymentStatus  string  `json:"payment_status"`
	Firstname      string  `json:"firstname"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	HouseName      string  `json:"house_name" validate:"required"`
	Street         string  `json:"street"`
	City           string  `json:"city"`
	State          string  `json:"state" validate:"required"`
	Pin            string  `json:"pin" validate:"required"`
}