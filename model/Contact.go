package model

import "gorm.io/gorm"

type Contact struct{
	Id int 					`form:"id" json:"id"`
	Contactname string 		`form:"contact_name" json:"contact_name"`
	Phonenumber string 		`form:"phonenumber" json:"phone_number"`
	Email string 			`form:"email" json:"email"`
	gorm.Model
}