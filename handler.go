package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var (
	receipts = make(map[string]Receipt)
	points   = make(map[string]int)
	mu       sync.Mutex
)

func generateID() string {
	return uuid.New().String()
}

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid receipt data", http.StatusBadRequest)
		return
	}

	id := generateID()
	points[id] = calculatePoints(receipt)

	mu.Lock()
	receipts[id] = receipt
	mu.Unlock()

	response := map[string]string{"id": id}
	json.NewEncoder(w).Encode(response)
}

// GetPoints handles GET /receipts/{id}/points
func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/receipts/") : len(r.URL.Path)-len("/points")]

	mu.Lock()
	pointsAwarded, exists := points[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": pointsAwarded}
	json.NewEncoder(w).Encode(response)
}
