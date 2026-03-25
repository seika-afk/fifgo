
package routes

import(
		"github.com/gorilla/mux"
	"gmg/controllers"
)


func RegisterRoutes(r *mux.Router){
	r.HandleFunc("/user/{id}",controllers.GetUser).Methods("POST")
	r.HandleFunc("/user",controllers.CreateUser).Methods("GET")
	r.HandleFunc("/user/{id}",controllers.DeleteUser).Methods("DELETE")



 
}

