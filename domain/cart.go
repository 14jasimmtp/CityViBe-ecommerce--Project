package domain

type Cart struct{
	CartID uint
	CartItems CartItems `json:"-"`
}

type CartItems struct{
	CartID uint `json:"cartId"`
	ProductID uint `json:"productId"`
	UserID uint `json:"userId"`
}