package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies 	[]Movie 

func getAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func getSingleMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item:= range movies{
		if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
	}
	}
}

func AddNewMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append (movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func EditMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append (movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode (&movie)
			movie.ID = params["id"]
			movies = append (movies, movie)
			json. NewEncoder (w). Encode (movie)
			return
		}	
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:"1", 
		Isbn: "039223", 
		Title:"Walking Dead", 
		Director: &Director{Firstname: "John", Lastname: "wick"}})

	movies = append(movies, Movie{
		ID:"2", 
		Isbn: "036223", 
		Title:"Not Walking Dead", 
		Director: &Director{Firstname: "John", Lastname: "wick"}})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getSingleMovie).Methods("GET")
	r.HandleFunc("/movies", AddNewMovie).Methods("POST")
	r.HandleFunc("/movies", EditMovie).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server on PORT 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}