package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

// Delete via append
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

// Get specific movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.IntN(100000000))
	movies = append(movies, movie)

	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}

}

// Removes existing movies then appends updated record (not how to work in proper DB systems)
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("content-type", "application/json")
	//params
	params := mux.Vars(r)
	//range over movies
	for index, item := range movies {
		if item.ID == params["id"] {
			//remove matching record
			movies = append(movies[:index], movies[index+1:]...)

			//insert new record with same id
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			err := json.NewEncoder(w).Encode(movies)
			if err != nil {
				return
			}
		}
	}
}

// Main func
func main() {
	r := mux.NewRouter()

	//Preload some data
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438227",
		Title: "Movie One",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe"},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "45455",
		Title: "Movie Two",
		Director: &Director{
			FirstName: "Stephen",
			LastName:  "Smith"},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
