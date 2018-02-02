package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func showHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", showHomePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":9876", router))
}
