package main

import (
	"encoding/json"
	"os"
	"time"
)

var offersFile = "offers.json"

// SaveOffers merges today's offers into JSON file
func SaveOffers(newOffers []Offer) error {
	today := time.Now().Format("2006-01-02")

	// Load existing offers
	existing, err := LoadOffers()
	if err != nil {
		return err
	}

	// Keep only offers that are NOT from today
	var filtered []Offer
	for _, o := range existing {
		if o.Date != today {
			filtered = append(filtered, o)
		}
	}

	// Add new ones (overwrite today's)
	updated := append(filtered, newOffers...)

	// Write file
	data, err := json.MarshalIndent(updated, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(offersFile, data, 0644)
}

// LoadOffers
func LoadOffers() ([]Offer, error) {
	if _, err := os.Stat(offersFile); os.IsNotExist(err) {
		return []Offer{}, nil
	}

	data, err := os.ReadFile(offersFile)
	if err != nil {
		return nil, err
	}

	var offers []Offer
	err = json.Unmarshal(data, &offers)
	return offers, err
}
