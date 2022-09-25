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
	m             sync.Mutex
}

var tables [NrOfTables]Table

func InitTables() {
	for i := 0; i < NrOfTables; i++ {
		tables[i].isFree = true
	}
}

func OperateTable(tableNr int) {
	tables[tableNr].m.Lock()
	defer tables[tableNr].m.Unlock()

	if tables[tableNr].isFree {
		tables[tableNr].isFree = false
		time.Sleep(time.Duration(2*TIMEUNIT) * time.Millisecond)

	} else if !tables[tableNr].wantsToOrder && !tables[tableNr].waitingOrder {
		tables[tableNr].wantsToOrder = true
		fmt.Printf("Table %d wants to order\n", tableNr)
		jobsTake <- tableNr

	} else if tables[tableNr].orderReceived {
		fmt.Printf("Table %d received order after %s\n", tableNr, time.Now().Sub(tables[tableNr].order.PickUpTime))
		time.Sleep(time.Duration(10*TIMEUNIT) * time.Millisecond)
		tables[tableNr].order = Order{}
		tables[tableNr].orderReceived = false
		tables[tableNr].isFree = true
		tables[tableNr].waitingOrder = false
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
