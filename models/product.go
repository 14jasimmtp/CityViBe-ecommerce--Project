package models

type Product struct {
	ProductID   int    `json:"productId"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	CategoryId  int    `json:"category" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Size        int    `json:"size" binding:"required"`
	color       string `json:"color" binding:"required"`
}
