package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                 string         `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
	Name               string         `json:"name" gorm:"not null"`
	Price              int            `json:"price" gorm:"not null"`
	ProductDescription string         `json:"description" gorm:"not null"`
}
