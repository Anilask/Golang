package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "Golang pointer", Author: "Mr Golang", Year: "2010"},
		Book{ID: 2, Title: "Gorotine", Author: "Mr Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang router", Author: "Mr Router", Year: "2012"},
		Book{ID: 4, Title: "Golang Concurrency", Author: "Mr currency", Year: "2013"},
		Book{ID: 5, Title: "good part", Author: "Mr. good", Year: "2014"})
	router.HandleFunc("/books", getBooks).Methods("Get")
	router.HandleFunc("/books/{id}", getBook).Methods("Get")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	//log.Println("Get all books")

	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Gets one book")
	params := mux.Vars(r)
	//log.Println(params)
	i, _ := strconv.Atoi(params["id"])
	// log.Println(reflect.TypeOf(params["id"]))
	log.Println(reflect.TypeOf(i))
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(books)
	log.Println(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("update the book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("Delete Book")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:1], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}
