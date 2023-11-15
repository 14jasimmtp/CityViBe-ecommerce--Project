package domain

type Product struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	Size        int      `json:"size"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Color       string   `json:"color"`
	Deleted     bool     `json:"delete" gorm:"default:false"`
}
type Category struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Category string `json:"category" gorm:"unique; not null"`
}
