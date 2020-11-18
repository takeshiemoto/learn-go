package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID int    `json:"id"`
	Do string `json:"do"`
}

var (
	todoStore = map[int]Todo{}
	IDCount   = 0
)

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var todo Todo
			if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			IDCount++
			todo.ID = IDCount
			todoStore[todo.ID] = todo

			fmt.Printf("%#+v\n", todoStore)
		case http.MethodGet:
			// まずスライスを用意する
			todos := []Todo{}
			for _, todo := range todoStore {
				todos = append(todos, todo)
			}
			if err := json.NewEncoder(w).Encode(&todos); err != nil {
				http.Error(w, err.Error(), 500)
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
