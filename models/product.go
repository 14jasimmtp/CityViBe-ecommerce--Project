package models

type Product struct {
	ProductID   int    `json:"productId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Size        string `json:"size" binding:"required"`
	color       string `json:"color" binding:"required"`
}
