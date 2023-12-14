package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" `
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
	OfferPrize  float64 `json:"offerprice"`
	Color       string  `json:"color"`
}

type UpdateProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
	OfferPrice  float64 `json:"offer_price"`
	Color       string  `json:"color"`
}

type AddProduct struct {
	ID          int    `json:"-"`
	Name        string `json:"name" validate:"required" form:"name"`
	Description string `json:"description" validate:"required" form:"description"`
	CategoryID  int    `json:"category" validate:"required,numeric" form:"category_id"`
	Size        int    `json:"size" validate:"required,numeric" form:"size_id"`
	Stock       int    `json:"stock" validate:"required,numeric" form:"stock"`
	Price       int    `json:"price" validate:"required,numeric" form:"price"`
	Color       string `json:"color" validate:"required" form:"color"`
	ImageURL    string `json:"imageurl"`
}

type Category struct {
	Category string `json:"category" binding:"required"`
}

type SetNewName struct {
	Current string `json:"current" validate:"required"`
	New     string `json:"new" validate:"required"`
}

type Offer struct {
	gorm.Model `json:"-"`
	Id         int       `json:"-" gorm:"primarykey"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Amount     int       `json:"amount"`
	MinPrice   int       `json:"minprice"`
	ValidFrom  time.Time `json:"-"`
	ValidUntil time.Time `json:"valid_until"`
	UsageLimit int       `json:"usage_limit"`
	UsedCount  int       `json:"-"`
	Category   int       `json:"category"`
	ProductId  int       `json:"product_id"`
}
