package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"log"

	"gmg/routes"


)


func main(){

//Creating a router 

	r:=mux.NewRouter()
	
//route all the routes
	routes.RegisterRoutes(r)


	
	//setup server
	log.Println("Server Running at http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000",r))



}

