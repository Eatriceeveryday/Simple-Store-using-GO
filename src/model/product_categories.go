package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategories struct {
	ID                 string         `json:"-" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
	ProductDescription string         `json:"description" gorm:"not null"`
	CategoryID         string         `json:"-"`
	Category           Category       `json:"-" gorm:"not null"`
	ProductID          string         `json:"-"`
	Product            Product        `json:"product"`
}
