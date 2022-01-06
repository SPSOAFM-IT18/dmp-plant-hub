package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// IDK how to make it work
// needs to reload when data changes
type LiveControl struct {
	Restart   bool `json:"restart"`
	PumpState bool `json:"pumpState"`
}

func CreateLiveControl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(LiveControl{Restart: false, PumpState: false})
	fmt.Println("live control")
}

func handleRequests() {
	http.HandleFunc("/live/control", CreateLiveControl)
}

func main() {
	handleRequests()
	log.Fatal(http.ListenAndServe(":5000", nil))
}
