package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/simple-crud-go-api/pkg/db"
	"github.com/simple-crud-go-api/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("Api is running")
	http.ListenAndServe(":4000", router)
}
