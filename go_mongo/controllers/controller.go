package controllers

import (
	"encoding/json"
	"gmg/models"
	"net/http"

	"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

type UserController struct{

		Session  *mgo.Session
}


func NewUserController(s * mgo.Session) * UserController{

return &UserController{s}
}
 

func (uc *UserController )GetUser(w http.ResponseWriter, r *http.Request){

	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id){
http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}



	oid := bson.ObjectIdHex(id)
	u:=models.User{}

	err := uc.Session.DB("mongo-golang").
    C("users").
    FindId(oid).
    One(&u)

if err != nil {
    w.WriteHeader(http.StatusNotFound)
    return
}
uj, err := json.Marshal(u)
if err != nil {
    http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
    return
}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
w.Write(uj)


}

func CreateUser(){}

func DeleteUser(){}

