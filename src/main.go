package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const TIMEUNIT int = 300
const NrOfTables = 5
const NrOfWaiters = 3

func getDistribution(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/distribution" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var order Order
	json.NewDecoder(r.Body).Decode(&order)

	jobsGive[order.WaiterId] <- order
}

func main() {
	http.HandleFunc("/distribution", getDistribution)

	ParseMenu()

	InitWaiterChs()
	InitTables()

	for i := 0; i < NrOfWaiters; i++ {
		go HandleWaiter(i, jobsTake, jobsGive[i])
	}

	for i := 0; i < NrOfTables; i++ {
		go HandleTable(i)
	}

	fmt.Printf("Server Dinning-Hall started on PORT 8020\n")
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatal(err)
	}

}
