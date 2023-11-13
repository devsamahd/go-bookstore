package routes

import (
	"github.com/gorilla/mux"
	"github.com/devsamahd/go-bookstore/pkg/controllers"
)

var BookstoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetSingleBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateSingleBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteSingleBook).Methods("DELETE")
}