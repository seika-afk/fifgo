
package routes

import(
		"github.com/gorilla/mux"
	"gmg/controllers"
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

//get session for mongodb
func getSession() *mgo.Session {
	s, err := mgo.DialWithTimeout("mongodb://127.0.0.1:27017", 5*time.Second)
	if err != nil {
		log.Println("Mongo connection error:", err)
		return nil
	}

	s.SetMode(mgo.Monotonic, true)
	return s
}


func RegisterRoutes(r *mux.Router) {
	session := getSession()

	if session == nil {
		log.Fatal("Failed to connect to MongoDB") // controlled exit
	}

	uc := controllers.NewUserController(session)

	r.HandleFunc("/user/{id}", uc.GetUser).Methods("GET")
	r.HandleFunc("/user", uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", uc.DeleteUser).Methods("DELETE")
}

