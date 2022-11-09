package main

import (
	"Lab2/Consumer/components"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

var aggregatorOrders components.Queue

func postOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order components.Order
	// Decode data that has been sent with the request and insert it into the new instance order
	_ = json.NewDecoder(r.Body).Decode(&order)
	// append order and return new order
	aggregatorOrders.Enqueue(order)
	json.NewEncoder(w).Encode(&order)

	fmt.Println("Consumer service received order from Aggregator: ", order)
}

func sendToAggregator() {
	for {
		var m sync.Mutex
		for aggregatorOrders.IsEmpty() != true {
			go sendOrders(&m)
			time.Sleep(time.Second * 1)
		}
	}
}

func sendOrders(m *sync.Mutex) {
	m.Lock()
	order := aggregatorOrders.Dequeue()
	m.Unlock()
	const myURL = "http://aggregator:5000/consumer"

	var requestBody, _ = json.Marshal(order)
	fmt.Printf("Order %v was sent to the Aggregator\n", string(requestBody))
	response, err := http.Post(myURL, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	time.Sleep(time.Second * 1)
}

func main() {
	// create new router
	router := mux.NewRouter()

	router.HandleFunc("/aggregator", postOrder).Methods("POST")

	go sendToAggregator()

	// run the server on port 5050
	log.Fatal(http.ListenAndServe(":5050", router))
}
