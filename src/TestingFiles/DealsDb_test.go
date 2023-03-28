package main

import (
	"BackendPkg"
	"testing"
	"strconv"
)

func TestFour(t *testing.T){
	// creates test database
	testDatabase := BackendPkg.Database{
		Name: "Testing Database",
	}

	// clear old database
	testDatabase.ClearPublixDeals()

	// creates test groccery store
	testPublix := BackendPkg.GroceryStore{
		Name:    "Publix",
		ZipCode: "32601",
	}
	// setup user groccery store
	testScraper := BackendPkg.Scraper{
		Store: testPublix,
	}
	// scrape all data
	testScraper.Scrape()

	newString := testScraper.Store.FindStart("view all results",(testScraper.DealsHTML))

	// organize scraped data
	testFoodSlice := testScraper.Store.OrganizeDeals(newString)
	
	// store publix data to .db file
	testDatabase.StorePublixDatabase(testFoodSlice)
	testDatabase.StoreDealsScrapedTime(testScraper.TimeLastDealsScraped)

	// compare scrape times
	if(testDatabase.ReadDealsScrapedTime().Format("2006-01-02 15:04:05") != testScraper.TimeLastDealsScraped.Format("2006-01-02 15:04:05")){
		t.Errorf("Scrape times don't match.")
	}

	// compare slices of food items
	for i := 0; i < len(testDatabase.ReadPublixDatabase()); i++ {
		if testDatabase.ReadPublixDatabase()[i].Name == "" {
			t.Errorf("Failed to return item #" + strconv.Itoa(i + 1) + " correctly.")
		}
	}
}