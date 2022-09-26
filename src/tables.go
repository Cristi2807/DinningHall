package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Table struct {
	isFree        bool
	wantsToOrder  bool
	waitingOrder  bool
	orderReceived bool
	order         Order
}

var tables sync.Map

func InitTables() {
	for i := 0; i < NrOfTables; i++ {
		tables.Store(i, Table{isFree: true})
	}
}

func OperateTable(tableNr int) {

	value, _ := tables.Load(tableNr)
	table := value.(Table)

	if table.isFree {
		table.isFree = false
		tables.Store(tableNr, table)
		time.Sleep(time.Duration(2*TIMEUNIT) * time.Millisecond)

	} else if !table.wantsToOrder && !table.waitingOrder {
		table.wantsToOrder = true
		tables.Store(tableNr, table)
		fmt.Printf("Table %d wants to order\n", tableNr)
		jobsTake <- tableNr

	} else if table.orderReceived {
		fmt.Printf("Table %d received order after %d\n", tableNr, time.Now().UnixMilli()-table.order.PickUpTime)
		fmt.Printf("RESTAURANT AVG RATING: %f\n\n\n", CalcAvgRating(float32(table.order.MaxWait), float32(time.Now().UnixMilli()-table.order.PickUpTime)))
		time.Sleep(time.Duration(20*TIMEUNIT) * time.Millisecond)
		table.order = Order{}
		table.orderReceived = false
		table.isFree = true
		table.waitingOrder = false
		tables.Store(tableNr, table)
	}

}

func HandleTable(tableNr int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randomInt := r1.Intn(21)

	time.Sleep(time.Duration(randomInt*TIMEUNIT) * time.Millisecond)

	for {
		OperateTable(tableNr)
	}
}

var sum float32 = 0
var n float32 = 0

func CalcAvgRating(MaxWait float32, RealTime float32) float32 {

	if RealTime < MaxWait {
		sum += 5
		n++
	} else if RealTime <= MaxWait*1.1 {
		sum += 4
		n++
	} else if RealTime <= MaxWait*1.2 {
		sum += 3
		n++
	} else if RealTime <= MaxWait*1.3 {
		sum += 2
		n++
	} else if RealTime <= MaxWait*1.4 {
		sum += 1
		n++
	} else {
		sum += 0
		n++
	}

	avg := sum / n
	return float32(avg)
}
