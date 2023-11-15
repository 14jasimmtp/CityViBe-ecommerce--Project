package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	CategoryId  int    `json:"category" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Size        int    `json:"size" binding:"required"`
	Color       string `json:"color" binding:"required"`
}

type Category struct {
	Category string `json:"category"`
}

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
