package controllers


import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"birthday_mysql/pkg/models"

	
)

var newBD models.BD
func CreateBD(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("Name")
    birthdate := r.URL.Query().Get("Birthdate")

    if name == "" || birthdate == "" {
        http.Error(w, "missing parameters", http.StatusBadRequest)
        return
    }

    bd := &models.BD{
        Name:      name,
        Birthdate: birthdate,
    }

    bd.CreateBD()

    res, _ := json.Marshal(bd)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
func GetAllBD(w http.ResponseWriter,r *http.Request){

	new_BD:=models.GetAllBD()
	res,_:=json.Marshal(new_BD)

		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		
		w.Write(res)



}


func GetBDByID(w http.ResponseWriter,r *http.Request){

	vars := mux.Vars(r)
	bdId := vars["bdId"]	
	ID,err := strconv.ParseInt(bdId,0,0)
	if err != nil{

			fmt.Println("error while parsing")

	}

	bD,_ := models.GetBDById(ID)

	res,_ := json.Marshal(bD)
	w.Header().Set("Content-Type","application/json")
	w.Write(res)

}
func DeleteBD(w http.ResponseWriter,r *http.Request){
    
		
  vars := mux.Vars(r)
	bdId := vars["bdId"]
	ID,err := strconv.ParseInt(bdId,0,0)
	 
	if err != nil{
			fmt.Println("error while parsing")

	}
		b := models.DeleteBD(ID)
		res,_ := json.Marshal(b)


		w.Header().Set("Content-Type","application/json")
	w.Write(res)


}



