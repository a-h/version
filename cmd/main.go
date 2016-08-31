package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/version", version)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello!")
}

func version(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "0.0.1")
}
