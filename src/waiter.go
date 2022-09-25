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
	tables[tableNr].m.Lock()
	defer tables[tableNr].m.Unlock()

	tables[tableNr].wantsToOrder = false
	tables[tableNr].waitingOrder = true
	tables[tableNr].order = createOrder(4)
	tables[tableNr].order.TableId = tableNr
	tables[tableNr].order.WaiterId = waiterId
	tables[tableNr].order.PickUpTime = time.Now()

	time.Sleep(time.Duration(2*TIMEUNIT) * time.Millisecond)

	orderMarshalled, _ := json.Marshal(tables[tableNr].order)
	responseBody := bytes.NewBuffer(orderMarshalled)

	http.Post("http://kitchen:8010/order", "application/json", responseBody)

}

func GiveOrderToTable(order Order) {
	tables[order.TableId].m.Lock()
	defer tables[order.TableId].m.Unlock()

	tables[order.TableId].orderReceived = true
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
