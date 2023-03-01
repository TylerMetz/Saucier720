package main

import (
	"BackendPkg"
	"testing"
	"time"
)

func TestTwo(t *testing.T) {
	// create new groccery store
	userPublix := BackendPkg.GroceryStore{
		Name:    "Publix",
		ZipCode: "32601",
	}

	// setup user groccery store
	programScraper := BackendPkg.Scraper{
		Store: userPublix,
	}

	// scrape all data
	programScraper.Scrape()

	now := time.Now()
	diff := now.Sub(programScraper.TimeLastDealsScraped)
	if diff.Seconds() > 1 {
		t.Errorf("Data was not scraped.")
	}

}
