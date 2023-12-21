package models

type CheckOutInfoDiscount struct {
	Address        []AddressRes `json:"address"`
	Cart           []Cart       `json:"cart"`
	TotalAmount    float64      `json:"Total Amount"`
	DiscountAmount float64      `json:"Discounted Amount"`
}

type CheckOutInfo struct {
	Address     []AddressRes `json:"address"`
	Cart        []Cart       `json:"cart"`
	TotalAmount float64      `json:"Total Amount"`
}

type OrderSuccessResponse struct {
	OrderID       int     `json:"order_id"`
	PaymentMethod string  `json:"payment_method"`
	TotalAmount   float64 `json:"total_amount"`
	PaymentStatus string  `json:"Payment_status"`
}

type ViewOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}

type ViewAdminOrderDetails struct {
	OrderDetails        AdminOrderDetails
	OrderProductDetails []OrderProductDetails
}

type OrderDetails struct {
	Id            int
	FinalPrice    float64
	PaymentMethod string
	PaymentStatus string
}

type AdminOrderDetails struct {
	UserID        int
	Id            string
	FinalPrice    float64
	PaymentMethod string
	PaymentStatus string
}

type OrderProductDetails struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	OrderStatus string  `json:"order_status"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type OrderProducts struct {
	ProductId string `json:"product_id"`
	Stock     int    `json:"-"`
}

type CombinedOrderDetails struct {
	Id            string  `json:"order_id"`
	FinalPrice    float64 `json:"final_price"`
	OrderStatus   string  `json:"order_status"`
	PaymentStatus string  `json:"payment_status"`
	Firstname     string  `json:"firstname"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	HouseName     string  `json:"house_name" validate:"required"`
	Street        string  `json:"street"`
	City          string  `json:"city"`
	State         string  `json:"state" validate:"required"`
	Pin           string  `json:"pin" validate:"required"`
}

type CheckOut struct {
	AddressID uint   `json:"address_id" validate:"required"`
	PaymentID uint   `json:"payment_id" validate:"required"`
}

type CancelDetails struct {
	OrderStatus   string  `json:"order_status"`
	Quantity      int     `json:"quantity"`
	PaymentStatus string  `json:"payment_status"`
	TotalPrice    float64 `json:"total_price"`
	OrderID       int     `json:"order_id"`
	ProductID     int     `json:"product_id"`
}

type SalesReport struct {
	TotalSales   float64
	TotalOrders  int64
	AverageOrder float64
}
