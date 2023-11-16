package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	Price       int    `json:"price" `
	CategoryId  int    `json:"category"`
	Stock       int    `json:"stock" `
	Size        int    `json:"size" `
	Color       string `json:"color" `
}

type Category struct {
	Category string `json:"category"`
}

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
