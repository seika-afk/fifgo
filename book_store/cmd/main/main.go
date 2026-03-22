package main

import (
	"book_store/pkg/config"
	"book_store/pkg/models"
	"book_store/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.Connect()

	db := config.GetDB()
	db.AutoMigrate(&models.Book{})

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
log.Println("Server running at http://localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
