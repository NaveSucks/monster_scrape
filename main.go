package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func init() {
	// Only load local config.env if it exists
	_ = godotenv.Load("config.env")
}

func main() {
	url := os.Getenv("SCRAPE_URL") // <- This works with Docker Compose env
	if url == "" {
		log.Fatal("SCRAPE_URL not set")
	}

	// --- Fetch once at startup (for debugging & freshness) ---
	runScrape(url)

	// --- Schedule daily scraping at 03:00 AM ---
	c := cron.New()
	_, err := c.AddFunc("0 3 * * *", func() {
		log.Println("Starting scheduled scrape...")
		runScrape(url)
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()

	// --- Start API (blocking) ---
	//port := os.Getenv("API_PORT")
	//if port == "" {
	//port = "8080" // fallback
	//}
	//StartAPI(port)
	StartAPI("8080")
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
