package main

import (
	"fmt"
	"net/http"
)

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
	fmt.Fprint(w, "ok")
}
