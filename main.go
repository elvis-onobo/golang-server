package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func main() {
	greeting := Greeting{Message: "Hello, world", Success: true}
	jsonData, err := json.Marshal(greeting)
	if err != nil {
		return
	}
	port := ":8080"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonData)
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	})
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started at http://localhost:%v", port)
}
