package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"book_store/pkg/utils"
	"book_store/pkg/models"

	
)

var newBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
		newBook := models.GetAllBooks()
		res, _ := json.Marshal(newBook)


		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		
		w.Write(res)


}

func GetBookById(w http.ResponseWriter,r *http.Request){

	vars := mux.Vars(r)
	bookId := vars["bookId"]	
	ID,err := strconv.ParseInt(bookId,0,0)

	if err != nil{

			fmt.Println("error while parsing")

	}

	bookDetails,_ := models.GetBookById(ID)

	res,_ = := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.Write(res)
}

