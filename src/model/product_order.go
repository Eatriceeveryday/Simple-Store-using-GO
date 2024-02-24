package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductOrder struct {
	ID        string         `json:"-" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	OrderID   string         `json:"orderId"`
	Order     Order          `json:"-"`
	ProductID string         `json:"-"`
	Product   Product        `json:"product" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
}
