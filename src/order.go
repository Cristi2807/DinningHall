package main

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type Order struct {
	Id      int   `json:"order_id"`
	Items   []int `json:"items"`
	MaxWait int   `json:"max_wait"`
}

var OrderID uint64

func incOrderID() uint64 {
	return atomic.AddUint64(&OrderID, 1)
}

func createOrder(NrOfItems int) (order Order) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

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
