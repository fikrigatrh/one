package models

import (
	"time"
)

type Merchant struct {
	ID           uint      `gorm:"primary_key" json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	UserId       int       `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	CreatedBy    int       `json:"created_by"`
	UpdatedBy    int       `json:"updated_by"`
}

type Outlet struct {
	ID         uint      `gorm:"primary_key" json:"-"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	MerchantId int       `json:"merchant_id"`
	OutletName string    `json:"outlet_name"`
	CreatedBy  int       `json:"created_by"`
	UpdatedBy  int       `json:"updated_by"`
}
