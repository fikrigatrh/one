package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	MerchantID int     `json:"merchant_id"`
	OutletID   int     `json:"outlet_id"`
	BillTotal  float64 `json:"bill_total"`
	CreatedBy  int     `json:"created_by"`
	UpdatedBy  int     `json:"updated_by"`
}

type Report struct {
	No           int    `json:"no" gorm:"-"`
	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
	Date         string `json:"date"`
	BillTotal    int    `json:"bill_total"`
}

type ReqReportMerchant struct {
	Date       string `json:"date"`
	MerchantId string `json:"merchant_id"`
	OutletId   string `json:"outlet_id"`
}

type ReqReportOutlet struct {
	Date     string `json:"date"`
	OutletId string `json:"outlet_id,omitempty" validate:"required"`
}
