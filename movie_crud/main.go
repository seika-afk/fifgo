package main

import (
"fmt"
"log"
"encoding/json"
"math/rand"
"net/http"
"strconv"
"github.com/gorilla/mux"

)

type Movie struct{

	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"Director"`

}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}


var movies[]Movie

////////////////////////////////////////// FUNCTIONS 

func getMovies(w http.ResponseWriter,r *http.Request){

		w.Header().Set("Content-Type","application/json")
 		fmt.Println("Called : getMovies")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter,r *http.Request){

		w.Header().Set("Content-Type","application/json")
 		params := mux.Vars(r)
	  for index, item := range movies{
				if item.ID == params["id"] {
			movies=append(movies[:index],movies[index+1:]...)
				break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	 
	for _,item := range movies{
			if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
	}}
}


func createMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(1000))
	movies= append(movies, movie)
		json.NewEncoder(w).Encode(movie)


}

func updateMovie(w http.ResponseWriter,r *http.Request){
	//set json content type 
		w.Header().Set("Content-Type","application/json")

	// param 
	params := mux.Vars(r)

	//loop over the movies , range  
		  for index, item := range movies{
				if item.ID == params["id"] {
			movies=append(movies[:index],movies[index+1:]...)
		// delete the movei with id that has been sent 
	

	// add a new movie -> the movie we send in body postman 



			var movie Movie
					_= json.NewDecoder(r.Body).Decode(&movie)
					movie.ID=params["id"]
					movies= append(movies, movie)
						json.NewEncoder(w).Encode(movie)
						return

		}
	}




}

func main(){

	r:=mux.NewRouter()
	//appending demo movies
	movies = append(movies, Movie{ID: "1",Isbn:"12213",Title:"POKEMON",Director: &Director{Firstname:"John",Lastname: "Smith"} })
  
	movies = append(movies, Movie{ID: "2",Isbn:"12313",Title:"DORAEMON",Director: &Director{Firstname:"not John",Lastname: "Smith"} })



	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")


	


	fmt.Printf("starting server at port 8080 \n")

	log.Fatal(http.ListenAndServe(":8080",r))

}


