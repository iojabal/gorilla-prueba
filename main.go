package main

import (
	"log"
	"net/http"
	"prueba/handlers"
	"prueba/models"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	models.MigrateArt()
	models.MigrarUser()
	//endpoints
	r.HandleFunc("/user/", handlers.InitHandler).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user/", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/products/", handlers.CreateProducts).Methods("POST")
	r.HandleFunc("/products/", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", handlers.UpdateArti).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3003", r))
}
