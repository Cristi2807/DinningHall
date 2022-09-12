package main

import (
	"fmt"
	"log"
	"net/http"
)

func getDistribution(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/distribution" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!\n")
	fmt.Printf("got /distribution request\n")
}

func main() {
	http.HandleFunc("/distribution", getDistribution) // Update this line of code

	fmt.Printf("Starting server at port 8020\n")
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatal(err)
	}
}
