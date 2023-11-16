package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CategoryId  int    `json:"category"`
	Size        int    `json:"size"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Color       string `json:"color"`
}

type Category struct {
	Category string `json:"category"`
}

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
