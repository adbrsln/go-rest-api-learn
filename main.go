package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/twinj/uuid"

	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct (Model)
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//init book var as a slice Book Struct
// slice is basically a variable length array. because u need to define the length of the array
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get all books
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Param
	//loop through books and find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//get all books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = uuid.NewV4().String() //mock id , not safe because its id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//get all books
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			//remove the book from the array
			books = append(books[:index], books[index+1:]...)

			//create new book
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] //mock id , not safe because its id
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//get all books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//init mux router
	router := mux.NewRouter()

	//mock data - @todo - implement DB
	books = append(books, Book{ID: uuid.NewV4().String(), Isbn: "AE1203", Title: "Book 1", Author: &Author{FirstName: "Adib", LastName: "Roslan"}})
	books = append(books, Book{ID: uuid.NewV4().String(), Isbn: "AE12034", Title: "Book 2", Author: &Author{FirstName: "Adib", LastName: "Roslan"}})

	//Route Handler /endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//run server
	log.Fatal(http.ListenAndServe(":8000", router))
}
