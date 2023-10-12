package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


 type Movie struct{
	ID string `json:"id"`
	Vinay string `json:"vinay"`
	Title string `json:"title"`
	Director *Director `json:"director"`
 }

 type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
 }

 var movies []Movie

func main() {
	r:= mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Vinay: "34255", Title: "Dummy0 Movie", Director: &Director{Firstname: "John", Lastname: "snow"}})
	movies = append(movies, Movie{ID: "2", Vinay: "34256", Title: "Dummy1 Movie", Director: &Director{Firstname: "Boney", Lastname: "snow"}})
	movies = append(movies, Movie{ID: "3", Vinay: "34257", Title: "Dummy2 Movie", Director: &Director{Firstname: "Johny", Lastname: "Deph"}})


	r.HandleFunc("/movies", getMovies).Methods("GET")
r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
r.HandleFunc("/movies", createMovie).Methods("POST")
r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}


func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies{

		if items.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, items := range movies {
		if items.ID == params["id"] {
			json.NewEncoder(w).Encode(items)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
// to update the movie we will simply delete the old movie and create a new movie in it
func updateMovie(w http.ResponseWriter, r *http.Request)  {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movie
	//delete the movie with the id that we have sent
	//add a new movie - the movie that we send in the body of postman
	for index, items := range movies {
		if items.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}