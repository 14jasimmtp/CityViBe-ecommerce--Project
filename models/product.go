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

type UpdateProduct struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Size        string `json:"size"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Color       string `json:"color"`
}

type AddProduct struct {
	ID          int    `json:"-"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  int    `json:"category" validate:"required,numeric"`
	Size        int    `json:"size" validate:"required,numeric"`
	Stock       int    `json:"stock" validate:"required,numeric"`
	Price       int    `json:"price" validate:"required,numeric"`
	Color       string `json:"color" validate:"required"`
}

type Category struct {
	Category string `json:"category" binding:"required"`
}

type SetNewName struct {
	Current string `json:"current" validate:"required"`
	New     string `json:"new" validate:"required"`
}
