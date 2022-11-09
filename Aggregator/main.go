package main

import (
	"Lab2/Aggregator/components"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func postOrderProducer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order components.Order
	// Decode data that has been sent with the request and insert it into the new instance order
	_ = json.NewDecoder(r.Body).Decode(&order)
	// append order and return new order
	components.ProducerOrders.Enqueue(order)
	json.NewEncoder(w).Encode(&order)

	fmt.Println("Aggregator service received order from Producer: ", order)
}

func postOrderConsumer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order components.Order
	// Decode data that has been sent with the request and insert it into the new instance order
	_ = json.NewDecoder(r.Body).Decode(&order)
	// append order and return new order
	components.ConsumerOrders.Enqueue(order)
	json.NewEncoder(w).Encode(&order)

	fmt.Println("Aggregator service received order from Consumer: ", order)
}

func main() {
	// create new router
	router := mux.NewRouter()

	router.HandleFunc("/producer", postOrderProducer).Methods("POST")
	router.HandleFunc("/consumer", postOrderConsumer).Methods("POST")

	go components.SendToProducer()
	go components.SendToConsumer()

	// run the server on port 5000
	log.Fatal(http.ListenAndServe(":5000", router))
}
