package model

import (
	"time"

	"gorm.io/gorm"
)

type Carts struct {
	ID         string         `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
	CustomerID string         `json:"-"`
	Customer   Customer       `json:"-"`
	ProductID  string         `json:"-"`
	Product    Product        `json:"product" gorm:"not null"`
	Quantity   int            `json:"quantity"`
}
