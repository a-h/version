package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var versionFlag = flag.Bool("v", false, "Displays the version and then quits.")

func main() {
	flag.Parse()

	// If the version flag is set, print the version and quit.
	if *versionFlag {
		fmt.Println(version)
		return
	}

	// Otherwise, start the web server.
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/version", versionHandler)

	fmt.Printf("Starting up application version %s\n", version)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello!")
}

func versionHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, version)
}
