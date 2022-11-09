package main

import (
	"Lab2/Producer/components"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

	fmt.Println("Producer service received order from Aggregator: ", order)
}

func main() {
	// create new router
	router := mux.NewRouter()

	//components.MakeOrder()
	router.HandleFunc("/aggregator", postOrder).Methods("POST")

	go components.MakeOrder()

	// run the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}
