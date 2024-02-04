// handlers/book_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

var books = []Book{
	{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", PublishedAt: "1979-10-12"},
	{Title: "1984", Author: "George Orwell", PublishedAt: "1949-06-08"},
	{Title: "To Kill a Mockingbird", Author: "Harper Lee", PublishedAt: "1960-07-11"},
	{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", PublishedAt: "1925-04-10"},
	{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", PublishedAt: "1954-07-29"},
	{Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", PublishedAt: "1997-06-26"},
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range books {
		if book.Title == params["title"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}
