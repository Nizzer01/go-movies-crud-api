package go_movies_crud_api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

func main() {
	r := mux.NewRouter()

	//Preload some data
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438227",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
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
