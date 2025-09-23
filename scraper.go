package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Offer struct {
	Discounter string `json:"discounter"`
	Price      string `json:"price"`
	Date       string `json:"date"`
}

// FetchOffers scrapes the SCRAPE_URL and returns single-can offers
func FetchOffers(url string) ([]Offer, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var offers []Offer
	currentDate := time.Now().Format("2006-01-02")

	doc.Find("div[role='listitem']").Each(func(i int, s *goquery.Selection) {
		discounter := s.Find("p.text-dark1.truncate.text-sm").Last().Text()
		discounter = strings.TrimSpace(discounter)

		price := s.Find("p.text-primary.text-base.font-bold").Text()
		price = strings.TrimSpace(price)

		if price == "" || strings.Contains(price, "ab") {
			return
		}

		offer := Offer{
			Discounter: discounter,
			Price:      price,
			Date:       currentDate,
		}
		offers = append(offers, offer)
	})

	log.Printf("Fetched %d offers", len(offers))
	return offers, nil
}
