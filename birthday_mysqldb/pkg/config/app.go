package config

import (
	"birthday_mysql/pkg/models"
	"github.com/jinzhu/gorm"
)

var db * gorm.DB


func Connect(){
	url := "gouser:pop090@tcp(127.0.0.1:3306)/birthday_store?charset=utf8&parseTime=True&loc=Local"

	d,err := gorm.Open("mysql",url)

	if err !=nil{
		panic(err)
	}

	db=d 
	models.SetDB(db)


}


func GetDB() *gorm.DB{

	return db 

}
