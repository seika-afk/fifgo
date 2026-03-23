package routes

import(
		"github.com/gorilla/mux"
	"birthday_mysql/pkg/controllers"
)


func RegisterBirthdateRoutes(r *mux.Router){
	r.HandleFunc("/birthday",controllers.CreateBD).Methods("POST")
	r.HandleFunc("/birthday",controllers.GetAllBD).Methods("GET")

	r.HandleFunc("/birthday/{bdId}",controllers.GetBDByID).Methods("GET")
	r.HandleFunc("/birthday/{bdId}",controllers.DeleteBD).Methods("DELETE")




}
