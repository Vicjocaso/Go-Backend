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
	Director *Director `json:"directo"`
}

type Director struct {
	Firstname string `json:"firsname"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "Application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMoviesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "Application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "Application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//Set json contect type
	w.Header().Set("content_Tyoe", "Application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "Application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "4123", Isbn: "1145223", Title: "Movie One", Director: &Director{Firstname: "Pedro", LastName: "Soto"}})
	movies = append(movies, Movie{ID: "1234", Isbn: "1142323", Title: "Transformers", Director: &Director{Firstname: "Smith", LastName: "Rodriguez"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviesById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
