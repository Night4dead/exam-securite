package main

import (
	"connection"
	"customhttp"

	"encoding/json"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"

	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	customhttp.HandleFunc(router, "/comment", AddComments).Methods("POST")
	customhttp.HandleFunc(router, "/connect", connection.Connect).Methods("POST")
	customhttp.HandleFunc(router, "/connectSafe", connection.ConnectSafe).Methods("POST")

	log.Println("Server start !")
	log.Fatal(http.ListenAndServe(":8080", cors.Default().Handler(router)))
}

type Comment struct {
	Comment string `json:"comment"`
}

var comments = make([]Comment, 0)

func AddComments(w http.ResponseWriter, r *http.Request) {
	log.Println("add a comment")
	var comment Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)
	comments = append(comments, comment)
	w.WriteHeader(200)
}
