package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

const TIMEUNIT int = 300

type MenuItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	PrepTime int    `json:"preparation-time"`
}

type Order struct {
	Id      int   `json:"order_id"`
	Items   []int `json:"items"`
	MaxWait int   `json:"max_wait"`
}

var Menu = []MenuItem{
	{Id: 1, Name: "Pizza", PrepTime: 20 * TIMEUNIT},
	{Id: 2, Name: "Salad", PrepTime: 10 * TIMEUNIT},
	{Id: 3, Name: "Zeama", PrepTime: 7 * TIMEUNIT},
	{Id: 4, Name: "Scallop Sashimi with Meyer Lemon Confit", PrepTime: 32 * TIMEUNIT},
	{Id: 5, Name: "Island Duck with Mulberry Mustard", PrepTime: 35 * TIMEUNIT},
	{Id: 6, Name: "Waffles", PrepTime: 10 * TIMEUNIT},
	{Id: 7, Name: "Aubergine", PrepTime: 20 * TIMEUNIT},
	{Id: 8, Name: "Lasagna", PrepTime: 30 * TIMEUNIT},
	{Id: 9, Name: "Burger", PrepTime: 15 * TIMEUNIT},
	{Id: 10, Name: "Gyros", PrepTime: 15 * TIMEUNIT},
	{Id: 11, Name: "Kebab", PrepTime: 15 * TIMEUNIT},
	{Id: 12, Name: "Unagi Maki", PrepTime: 20 * TIMEUNIT},
	{Id: 13, Name: "Tobacco Chicken", PrepTime: 30 * TIMEUNIT},
}

var OrderID uint64

func incOrderID() uint64 {
	return atomic.AddUint64(&OrderID, 1)
}

func createOrder(NrOfItems int) (order Order) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//OrderID = (OrderID + 1) % 100000
	order.Id = int(incOrderID())

	var MaxPrepTime = 0
	for i := 0; i < NrOfItems; i++ {
		order.Items = append(order.Items, r1.Intn(13)+1)
		if Menu[order.Items[i]-1].PrepTime > MaxPrepTime {
			MaxPrepTime = Menu[order.Items[i]-1].PrepTime
		}
	}

	order.MaxWait = int(float32(MaxPrepTime) * 1.3)

	return
}

func RandOrders() {

	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		randomInt := r1.Intn(7) + 4

		orderMarshalled, _ := json.Marshal(createOrder(4))
		responseBody := bytes.NewBuffer(orderMarshalled)

		http.Post("http://kitchen:8010/order", "application/json", responseBody)

		time.Sleep(time.Duration(randomInt) * time.Second)
	}
}

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
	fmt.Printf("%d %v %d \n\n", order.Id, order.Items, order.MaxWait)
}

func main() {
	http.HandleFunc("/distribution", getDistribution)

	go RandOrders()
	go RandOrders()
	go RandOrders()
	//go RandOrders()

	fmt.Printf("Server DinningHall started on PORT 8020\n")
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatal(err)
	}
}
