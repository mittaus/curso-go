package main

import (
	// "fmt"

	"database/sql"
	"log"
	"net/http"

	// "./controllers"
	"example.com/me/controllers"
	"example.com/me/driver"
	// "./driver"
	// "./models"
	"example.com/me/models"

	"os"

	"github.com/gorilla/mux"
	// "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var books []models.Book

var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}
	// port := os.Getenv("PORT")
	port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }
	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	done := make(chan bool)
	go http.ListenAndServe(":" + port, router)
	log.Printf("Server started at port %v", port)
	<-done
	log.Fatal(http.ListenAndServe(":" + port, router))
}
