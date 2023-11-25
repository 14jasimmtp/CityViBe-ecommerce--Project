package domain

import (
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Email    string `gorm:"UNIQUE"`
	Name     string
	AdminID  uuid.UUID `gorm:"PRIMARY KEY"`
	Password string
}
