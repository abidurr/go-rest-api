package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

type Author struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Initialize books variable as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	
}

// Create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	
}


func main() {

	// Initialize router as r
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", ISBN: "456654654", Title: "End of Race War", Author: &Author {FirstName: "Mr.", LastName: "Man"}})
	books = append(books, Book{ID: "2", ISBN: "656788654", Title: "World Peace", Author: &Author {FirstName: "John", LastName: "Titor"}})

	// Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}