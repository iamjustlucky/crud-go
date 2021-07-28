package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID int             		`form:"id" json:"id"`
	Code string 			`form:"code" json:"code"`
	Name string 			`form:"name" json:"name"`
	Price decimal.Decimal 	`form:"price" json:"price" sql:"type:decimal(16,2);"`
	gorm.Model
}