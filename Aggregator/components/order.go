package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Order struct {
	Id       int     `json:"id"`
	Items    []int   `json:"items"`
	Priority int     `json:"priority"`
	MaxWait  float32 `json:"max_wait"`
}

var ProducerOrders Queue
var ConsumerOrders Queue

func SendToProducer() {
	for {
		var m sync.Mutex
		for ConsumerOrders.IsEmpty() != true {
			go sendOrdersProducer(&m)
			time.Sleep(time.Second * 1)
		}
	}
}

func SendToConsumer() {
	for {
		var m sync.Mutex
		for ProducerOrders.IsEmpty() != true {
			go sendOrdersConsumer(&m)
			time.Sleep(time.Second * 1)
		}
	}
}

func sendOrdersProducer(m *sync.Mutex) {
	m.Lock()
	order := ConsumerOrders.Dequeue()
	m.Unlock()
	const myURL = "http://producer:8000/aggregator"

	var requestBody, _ = json.Marshal(order)
	fmt.Printf("Order %v was sent to the Producer\n", string(requestBody))
	response, err := http.Post(myURL, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	time.Sleep(time.Second * 1)
}

func sendOrdersConsumer(m *sync.Mutex) {
	m.Lock()
	order := ProducerOrders.Dequeue()
	m.Unlock()
	var requestBody, _ = json.Marshal(order)
	const myURL = "http://consumer:5050/aggregator"

	fmt.Printf("Order %v was sent to the Consumer\n", string(requestBody))

	response, err := http.Post(myURL, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	time.Sleep(time.Millisecond * 500)
}
