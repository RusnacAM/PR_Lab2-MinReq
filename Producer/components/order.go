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

// Function to generate order
func GetOrder() Order {
	foods, maxTime := GetItems()
	return Order{
		Id:       Random(1, 100000),
		Items:    foods,
		Priority: Random(1, 5),
		MaxWait:  maxTime,
	}
}

func GetMaxTime(foods []int) float32 {
	var maxTime int
	for _, v := range foods {
		var curr = v - 1
		var tempTime = Menu[curr].preparationTime

		if tempTime > maxTime {
			maxTime = tempTime
		}
	}

	var finalTime = float32(maxTime) * 1.3

	return finalTime
}

func GetItems() ([]int, float32) {
	var amount = Random(1, 5)
	var items []int
	for i := 0; i <= amount; i++ {
		items = append(items, Random(1, 13))
	}

	var maxTime = GetMaxTime(items)

	return items, maxTime
}

func MakeOrder() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 500)
		go SendOrderFoodOrd(&wg)
	}
	wg.Wait()
}

func SendOrderFoodOrd(wg *sync.WaitGroup) {
	const myURL = "http://aggregator:5000/producer"

	newOrder := GetOrder()
	requestBody, _ := json.Marshal(newOrder)

	fmt.Printf("Order %v was sent to the Aggregator.\n", string(requestBody))

	response, err := http.Post(myURL, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	wg.Done()
}
