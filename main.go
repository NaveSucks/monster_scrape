package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	// Load environment variables
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env")
	}

	url := os.Getenv("SCRAPE_URL")
	if url == "" {
		log.Fatal("SCRAPE_URL not set in config.env")
	}

	// --- Fetch once at startup (for debugging & freshness) ---
	runScrape(url)

	// --- Schedule daily scraping at 03:00 AM ---
	c := cron.New()
	_, err = c.AddFunc("0 3 * * *", func() {
		log.Println("Starting scheduled scrape...")
		runScrape(url)
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()

	// --- Start API (blocking) ---
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // fallback
	}
	StartAPI(port)
}

// runScrape encapsulates fetch + save logic
func runScrape(url string) {
	offers, err := FetchOffers(url)
	if err != nil {
		log.Println("Error fetching offers:", err)
		return
	}
	err = SaveOffers(offers)
	if err != nil {
		log.Println("Error saving offers:", err)
	} else {
		log.Printf("Saved %d offers", len(offers))
	}
}
