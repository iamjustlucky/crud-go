package handler

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func HandleDB()  {
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/mysqlgolang?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		log.Println("Connection Failed", err)
	} else {
		log.Println("Connection Success")
	}
}