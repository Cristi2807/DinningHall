package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var jobsGive [NrOfWaiters]chan Order
var jobsTake = make(chan int, NrOfTables)

func InitWaiterChs() {
	for i := 0; i < NrOfWaiters; i++ {
		jobsGive[i] = make(chan Order, NrOfTables)
	}
}

func ReceiveOrderFromTable(tableNr, waiterId int) {

	value, _ := tables.Load(tableNr)
	table := value.(Table)

	table.wantsToOrder = false
	table.waitingOrder = true
	table.order = createOrder(4)
	table.order.TableId = tableNr
	table.order.WaiterId = waiterId
	table.order.PickUpTime = time.Now().UnixMilli()

	tables.Store(tableNr, table)

	time.Sleep(time.Duration(2*TIMEUNIT) * time.Millisecond)

	orderMarshalled, _ := json.Marshal(table.order)
	responseBody := bytes.NewBuffer(orderMarshalled)

	http.Post("http://kitchen:8010/order", "application/json", responseBody)

}

func GiveOrderToTable(order Order) {
	value, _ := tables.Load(order.TableId)
	table := value.(Table)

	table.orderReceived = true

	tables.Store(order.TableId, table)
}

func HandleWaiter(WaiterId int, jobsTake <-chan int, jobsGive <-chan Order) {

	for {
		select {
		case tableGive := <-jobsGive:
			{
				GiveOrderToTable(tableGive)
			}
		default:
			select {
			case tableGive := <-jobsGive:
				{
					GiveOrderToTable(tableGive)
				}
			case tableTake := <-jobsTake:
				{
					s1 := rand.NewSource(time.Now().UnixNano())
					r1 := rand.New(s1)

					randomInt := r1.Intn(3) + 2

					time.Sleep(time.Duration(randomInt*TIMEUNIT) * time.Millisecond)

					ReceiveOrderFromTable(tableTake, WaiterId)

				}
			}
		}
	}

}
