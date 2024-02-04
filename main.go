// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"book-app/handlers" 
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.GetBooksHandler).Methods("GET")
	r.HandleFunc("/books/{title}", handlers.GetBookHandler).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
