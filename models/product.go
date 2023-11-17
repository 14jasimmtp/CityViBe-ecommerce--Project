package models

type Product struct {
	ID          int    `json:"-"`
	Name        string `json:"name" `
	Description string `json:"description"`
	Category    string `json:"category"`
	Size        string `json:"size"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Color       string `json:"color"`
}

type AddProduct struct {
	ID          int    `json:"-"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CategoryID  int    `json:"category" binding:"required"`
	Size        int    `json:"size" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Color       string `json:"color" binding:"required"`
}

type Category struct {
	Category string `json:"category" binding:"required"`
}

type SetNewName struct {
	Current string `json:"current" binding:"required"`
	New     string `json:"new" binding:"required"`
}
