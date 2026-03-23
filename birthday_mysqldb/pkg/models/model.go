package models

import "github.com/jinzhu/gorm"


var db *gorm.DB

//structure of data being stored
type BD struct{
		gorm.Model
		Name string `json:"name"`
		Birthdate string `json:"birthdate"`
}

//set SetDB
func SetDB(d *gorm.DB){

	db=d 

}


//////////////////////////

func ( b *BD) CreateBD()*BD{

	db.Create(b)
	return b

}

func GetAllBD() []BD{
	var BDs []BD
	
	db.Find(&BDs)

	return BDs

}



func GetBDById(Id int64) (*BD,*gorm.DB){
var getBD BD
	db := db.Where("ID=?",Id).Find(&getBD)
	return &getBD,db

}


func DeleteBD(Id int64) BD{
		var bd BD 
		db.Where("Id=?",Id).Find(&bd)
		db.Delete(&bd)
	  return bd}




