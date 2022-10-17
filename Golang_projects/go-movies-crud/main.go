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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// --------------------
// movies list
// --------------------
func moviesHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// ---------------------------------
// Delete movie with given id
// ---------------------------------
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//ranging over movies to delete the requested one
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	//returning remaining movies
	json.NewEncoder(w).Encode(movies)
}

// ---------------------------------
// returning movie with given id
// ---------------------------------
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // getting movie through its id

	//ranging over movies to return a particular one
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// ---------------------------------
// create movie
// ---------------------------------
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// ---------------------------------
// update movie with given id
// ---------------------------------
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set the content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)

	//ranging
	for index, item := range movies {
		if item.ID == params["id"] {
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

func main() {

	// have some movies in the slice
	movies = append(movies, Movie{ID: "1", Isbn: "453234", Title: "Shutter Island", Director: &Director{Firstname: "Chris", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "838211", Title: "Prestige", Director: &Director{Firstname: "Steve", Lastname: "Harvey"}})
	movies = append(movies, Movie{ID: "3", Isbn: "546343", Title: "Fight Club", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	//creating routes
	r := mux.NewRouter()
	r.HandleFunc("/movies", moviesHandle).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting port at 6969")

	//creating port
	if err := http.ListenAndServe(":6969", r); err != nil {
		log.Fatal(err)
	}
}
