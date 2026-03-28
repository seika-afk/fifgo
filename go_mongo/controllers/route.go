package controllers

import (
	"encoding/json"
	"gmg/models"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{

		Session  *mgo.Session
}


func NewUserController(s * mgo.Session) * UserController{

return &UserController{s}
}
 


///////////////////////////////////////////////// GET USER 


func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}

	session := uc.Session.Copy()
	defer session.Close()

	err := session.DB("mongo-golang").
		C("users").
		FindId(oid).
		One(&u)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}



/////////////////////////////////////////////// CREATE USER 
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	session := uc.Session.Copy()
	defer session.Close()

	var u models.User

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate new ObjectId
	u.ID = bson.NewObjectId()

	// Insert into DB
	err = session.DB("mongo-golang").
		C("users").
		Insert(u)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(u)
}
//////////////////////////////////////////////  DELETE USER 
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	session := uc.Session.Copy()
	defer session.Close()

	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	err := session.DB("mongo-golang").
		C("users").
		RemoveId(oid)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(map[string]string{
	"message": "User deleted",
})
}

