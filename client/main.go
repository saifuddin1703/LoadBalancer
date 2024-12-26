package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Request structure to parse the incoming JSON
type Request struct {
	Name string `json:"name"`
}

// Handler for the GET endpoint
func getNameHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON from the request body
	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Respond with the name
	response := fmt.Sprintf("Hello, %s! from server 1", req.Name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/get-name", getNameHandler)
	port := 6000
	// Start the server on port 8080
	fmt.Printf("Starting server on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
