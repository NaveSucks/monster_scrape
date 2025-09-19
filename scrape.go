package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

type Offer struct {
	Discounter string
	Price      string
	Date       string
}

func main() {
	// Load environment variables from config.env
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env file")
	}

	url := os.Getenv("SCRAPE_URL")
	if url == "" {
		log.Fatal("SCRAPE_URL not set in config.env")
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var offers []Offer
	currentDate := time.Now().Format("2006-01-02")

	doc.Find("div[role='listitem']").Each(func(i int, s *goquery.Selection) {
		discounter := s.Find("p.text-dark1.truncate.text-sm").Last().Text()
		discounter = strings.TrimSpace(discounter)

		price := s.Find("p.text-primary.text-base.font-bold").Text()
		price = strings.TrimSpace(price)

		if strings.Contains(price, "ab") {
			return
		}

		offer := Offer{
			Discounter: discounter,
			Price:      price,
			Date:       currentDate,
		}
		offers = append(offers, offer)
	})

	for _, o := range offers {
		fmt.Printf("Discounter: %s | Price: %s | Date: %s\n", o.Discounter, o.Price, o.Date)
	}
}
