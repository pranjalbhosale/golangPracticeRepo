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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname`
	LastName  string `json:"Lastname`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	//can do w.Header().add too
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {

		if item.Id == params["id"] {

			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {

		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.Id == params["id"] {

			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

		}
	}
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1001", Isbn: "436637", Title: "Harry Potter", Director: &Director{FirstName: "jk", LastName: "rowling"}})
	movies = append(movies, Movie{Id: "1002", Isbn: "763563", Title: "killers of the flower moon", Director: &Director{FirstName: "Martin", LastName: "Scorsese"}})

	//define different API

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("port started")
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
