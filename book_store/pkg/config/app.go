package config

import (
"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"


)

var db * gorm.DB


func Connect(){
	url := "gouser:pop090@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
	d,err :=gorm.Open("mysql",url)

	if err!= nil{
			panic(err)
			
	}

	db=d
}

func GetDB() *gorm.DB{
	return db


}
