package domain

type Wishlist struct {
	ID        uint    `gorm:"UNIQUE;NOT NULL"`
	User      User    `gorm:"foriegnkey:UserID"`
	UserID    uint    `gorm:"NOT NULL"`
	Product   Product `gorm:"foriegnkey:ProductID"`
	ProductID uint    `gorm:"NOT NULL"`
}
