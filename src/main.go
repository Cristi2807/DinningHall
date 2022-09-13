package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getDistribution(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/distribution" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Printf("got /distribution request\n\n")
}

func sendPostRequest() {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Toby@example.com",
	})
	responseBody := bytes.NewBuffer(postBody)

	http.Post("http://kitchen:8010/order", "application/json", responseBody)
}

func check() {
	for {
		sendPostRequest()
		time.Sleep(3 * time.Second)

	}
}

func main() {
	http.HandleFunc("/distribution", getDistribution)

	go check()

	fmt.Printf("Server DinningHall started on PORT 8020\n")
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatal(err)
	}
}
