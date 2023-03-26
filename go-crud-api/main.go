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
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
}

var movies []Movie

func handleMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			fmt.Print("yolo")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func newMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		if item.ID == params["id"] {
			movie.ID = item.ID
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Director: &Director{Name: "John Smokes"}})
	r.HandleFunc("/", handleMovies).Methods("GET")
	r.HandleFunc("/getmovie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/updatemovie/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/deletemovie/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/createmovie", newMovie).Methods("PUT")

	fmt.Printf("Starting Server at 5001")
	if err := http.ListenAndServe(":5001", r); err != nil {
		log.Fatal(err)
	}
}
