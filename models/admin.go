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

type TimePeriod struct {
	Year      string
	Month     string
	Week      string
	Startdate string
	EndDate   string
}

type DashBoardUser struct {
	TotalUsers  int `json:"Totaluser"`
	BlockedUser int `json:"Blockuser"`
}
type DashBoardProduct struct {
	TotalProducts     int `json:"Totalproduct"`
	OutofStockProduct int `json:"Outofstock"`
}
type DashboardOrder struct {
	DeliveredOrderProducts int
	PendingOrderProducts   int
	CancelledOrderProducts int
	TotalOrderItems        int
	TotalOrderQuantity     int
}
type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}
type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}
type CompleteAdminDashboard struct {
	DashboardUser    DashBoardUser
	DashboardProduct DashBoardProduct
	DashboardOrder   DashboardOrder
	DashboardRevenue DashboardRevenue
	DashboardAmount  DashboardAmount
}
