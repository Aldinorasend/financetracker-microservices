package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func main() {
	// Creating endpoint and calling handler
	http.HandleFunc("/health", healthHandler)

	// Make it easier to change port
	fmt.Println("Server running on http://localhost:8080")

	// Make error handler to listen port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// Output shown in browser
func healthHandler(w http.ResponseWriter, r *http.Request) {
	res := HealthResponse{
		Status:  "OK",
		Service: "transaction-service",
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(res)
}
