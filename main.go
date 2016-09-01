package main

import (
	"fmt"
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
	r.HandleFunc("/version/", versionHandler)
	r.HandleFunc("/Version/", versionHandler)

	fmt.Printf("Starting up application version %s\n", version)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello!")
}

func versionHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, version)
}
