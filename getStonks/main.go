package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getAAPLStockPrice() (string, error) {
	url := "https://finance.yahoo.com/quote/AAPL/"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers to mimic a browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %v", err)
	}

	// Try multiple selectors in order of likelihood
	selectors := []string{
		`fin-streamer[data-symbol="AAPL"][data-field="regularMarketPrice"]`, // Modern selector
		`[data-test="qsp-price"]`, // Alternative test attribute
		`#quote-header-info [data-field="regularMarketPrice"]`, // Container-based
		`.Fw\(b\).Fz\(36px\)`, // Class-based fallback
	}

	var price string
	for _, selector := range selectors {
		price = doc.Find(selector).First().Text()
		if price != "" {
			break
		}
	}

	price = strings.TrimSpace(price)
	if price == "" {
		return "", fmt.Errorf("could not find stock price using any selector")
	}

	return price, nil
}

func main() {
	price, err := getAAPLStockPrice()
	if err != nil {
		log.Fatalf("Error getting AAPL stock price: %v", err)
	}

	fmt.Printf("Current AAPL Stock Price: $%s\n", price)
}
