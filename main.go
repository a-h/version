package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/Version", versionHandler)
	r.HandleFunc("/version", versionHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello!")
}

func versionHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, version)
}
