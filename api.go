package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// StartAPI starts an HTTP server that serves the offers
func StartAPI(port string) {
	http.HandleFunc("/offers", func(w http.ResponseWriter, r *http.Request) {
		data, err := LoadOffers()
		if err != nil {
			http.Error(w, "Error loading offers", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	log.Printf("API running on http://localhost:%s/offers", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
