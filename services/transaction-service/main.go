package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Defining structure healthResponse to respond to healthcheck request
type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

// defining structure transaction to respond to transaction request
type Transaction struct {
	ID          int     `json:"id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

// Slice for saving transaction
var transactions []Transaction
var nextID = 1

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	// Creating endpoint and calling handler
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/transactions", createTransactionHandler)

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

	// Delivering structured response
	res := HealthResponse{
		Status:  "OK",
		Service: "transaction-service",
		Version: "1.0.0",
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(res)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// Delivering structured response
	res := HelloResponse{
		Message: "Hello from Financial Tracker",
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(res)
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Defining variable transaction
	var transaction Transaction

	// decoding request body
	err := json.NewDecoder(r.Body).Decode(&transaction)

	// checking error during decoding
	if err != nil {
		http.Error(w, "Invalid Req Body", http.StatusBadRequest)
		return
	}

	// Creating ID and increasing nextID
	transaction.ID = nextID
	nextID++

	// Appending transaction to slice
	transactions = append(transactions, transaction)

	// Returning the same transaction as response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Encoding and delivering response
	json.NewEncoder(w).Encode(transaction)

}
