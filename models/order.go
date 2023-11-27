package models

type CheckOutInfo struct {
	Address     []AddressRes `json:"address"`
	Cart        []Cart       `json:"cart"`
	TotalAmount float64      `json:"Total Amount"`
}

type OrderSuccessResponse struct {
	OrderID        uint   `json:"order_id"`
	ShipmentStatus string `json:"shipment_status"`
}

type ViewOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}

type OrderDetails struct {
	Id            string
	FinalPrice    float64
	OrderStatus   string
	PaymentMethod string
	PaymentStatus string
}

type OrderProductDetails struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type OrderProducts struct {
	ProductId string `json:"product_id"`
	Stock     int    `json:"-"`
}

type CombinedOrderDetails struct {
	Id                  string `json:"order_id"`
	FinalPrice          float64 `json:"final_price"`
	OrderStatus         string  `json:"order_status"`
	PaymentStatus       string  `json:"payment_status"`
	Firstname           string  `json:"firstname"`
	Email               string  `json:"email"`
	Phone               string  `json:"phone"`
	HouseName           string  `json:"house_name" validate:"required"`
	Street              string  `json:"street"`
	City                string  `json:"city"`
	State               string  `json:"state" validate:"required"`
	Pin                 string  `json:"pin" validate:"required"`
}

type CheckOut struct{
	AddressID uint `json:"address_id" binding:"required"`
	PaymentID uint `json:"payment_id"`
}
