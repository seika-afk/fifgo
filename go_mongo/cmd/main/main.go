package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"gopkg.in/mgo.v2"

	"gmg/controllers"

	"gmg/routes"


)


func main(){


	r:=mux.NewRouter()
	uc := controllers.NewUserController(getSession())
	//route all the routes
	routes.RegisterRouter(r)


	
	//setup server
	log.Println("Server Running at http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000",r))



}

func getSession() *mgo.Session{
	s,err:=mgo.Dial("mongodb://localhost:27107")	
	if err != nil{
			panic(err)
	}
return s
}
