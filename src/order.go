package main

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type PreparedFood struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type Order struct {
	Id             int            `json:"order_id"`
	Items          []int          `json:"items"`
	MaxWait        int            `json:"max_wait"`
	TableId        int            `json:"table_id"`
	WaiterId       int            `json:"waiter_id"`
	Priority       int            `json:"priority"`
	PickUpTime     int64          `json:"pick_up_time"`
	CookingDetails []PreparedFood `json:"cooking_details"`
	CookingTime    int            `json:"cooking_time"`
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
		if menu[order.Items[i]-1].PreparationTime > MaxPrepTime {
			MaxPrepTime = menu[order.Items[i]-1].PreparationTime
		}
	}

	order.MaxWait = int(float32(MaxPrepTime) * 1.3 * float32(TIMEUNIT))

	return
}
