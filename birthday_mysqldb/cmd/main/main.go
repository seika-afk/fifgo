package main

import (
	"birthday_mysql/pkg/config"
	"birthday_mysql/pkg/models"
	"birthday_mysql/pkg/routes"


	"net/http"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)


func main(){
	
	config.Connect()

	//connect to db and check if db has same as struct
	db:=config.GetDB()
	db.AutoMigrate(&models.BD{})

//initiate NewRouter
	r:= mux.NewRouter()

	//register routes 
	routes.RegisterBirthdateRoutes(r)

	//setup server
	log.Println("Server Running at http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000",r))

}
