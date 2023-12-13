package domain

type Product struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	CategoryID  uint     `json:"category_id" form:"category_id"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID;"`
	SizeID      int      `json:"size_id" form:"size_id"`
	Size        Size     `json:"-" gorm:"foriegnkey:SizeID;"`
	Stock       int      `json:"stock" form:"stock"`
	Price       float64  `json:"price" form:"price"`
	OfferPrize  float64  `json:"offerprice"`
	Color       string   `json:"color" form:"color"`
	Deleted     bool     `json:"delete" gorm:"default:false"`
	ImageURL    string   `json:"imageurl"`
}
type Category struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Category string `json:"category" gorm:"unique; not null"`
}

type Size struct {
	ID   uint   `json:"id" gorm:"unique; not null"`
	Size string `json:"size" gorm:"unique; not null"`
}

type ProductImage struct {
	ID        uint   `json:"id" gorm:"primarykey;unique; not null"`
	ImageUrl  string `json:"image_url"`
	ProductID uint   `json:"product_id"`
}
