package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	items := []Item{
		{ID: "1", Name: "Item1", Price: 100},
		{ID: "2", Name: "Item2", Price: 200},
		{ID: "3", Name: "Item3", Price: 300},
		{ID: "4", Name: "Item4", Price: 400},
		{ID: "5", Name: "Item5", Price: 500},
		{ID: "6", Name: "Item6", Price: 600},
		{ID: "7", Name: "Item7", Price: 700},
		{ID: "8", Name: "Item8", Price: 800},
		{ID: "9", Name: "Item9", Price: 900},
		{ID: "10", Name: "Item10", Price: 1000},
	}

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Error marshaling item:", err)
			continue
		}

		resp, err := http.Post("http://localhost:8080/item", "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}

		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Failed to create item %s: %s\n", item.ID, resp.Status)
		} else {
			fmt.Printf("Successfully created item %s\n", item.ID)
		}
	}
}
