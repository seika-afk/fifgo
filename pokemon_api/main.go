package main
//imports 
import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

//Pokemon Struct -> id,name,type
type Pokemon struct{

	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
//trainer Struct ->Name
type Trainer struct{
	Name string `json:"name"`
}

var pokemons[]Pokemon

// FUNCTIONS

func getPokemons(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(pokemons)

}

func deletePokemon(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","application/json")
		params := mux.Vars(r)

		for index,item := range pokemons{

				if item.ID== params["id"]{
			pokemons=append(pokemons[:index],pokemons[index+1:]...)
			break
		}

	}

json.NewEncoder(w).Encode(pokemons)
}

func getPokemon(w http.ResponseWriter,r *http.Request){
w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	 
	for _,item := range pokemons{
			if item.Type == params["type"]{
			json.NewEncoder(w).Encode(item)
			return 
	}}


}

func addPokemon(){}


















//MAIN Function 
func main(){
	r := mux.NewRouter()

	pokemons=append(pokemons,Pokemon{ID:"1",Name:"Pikachu",Type:"electric"})
	pokemons=append(pokemons,Pokemon{ID:"2",Name:"Raichu",Type:"electric"})

	pokemons=append(pokemons,Pokemon{ID:"1",Name:"Charmander",Type:"Fire"})



//getPokemons
	r.HandleFunc("/all_pokemons",getPokemons).Methods("GET")
//get pokemon by index
	r.HandleFunc("/pokemon/{type}",getPokemon).Methods("GET")
	r.HandleFunc("/pokemons/{id}",deletePokemon).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")

	log.Fatal(http.ListenAndServe(":8080",r))

}














