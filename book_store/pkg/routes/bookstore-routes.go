package routes

import (
	"github.com/gorilla/mux"
	"book_store/pkg/controllers"
)

func RegisterBookStoreRoutes(r *mux.Router) {
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
