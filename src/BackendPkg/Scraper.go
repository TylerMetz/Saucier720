package BackendPkg

import (
	"fmt"
	"html"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	_"unicode/utf8"
)

var UserZipCodePlaceholder string = "32601"

type Scraper struct {
	PublixDeals			[]FoodItem
	WalmartDeals 		[]FoodItem
}

func (s *Scraper)CheckIfScrapeNewDeals(d Database){

	// EST
	location, _ := time.LoadLocation("America/New_York")

	// create a time object for last Thurday at 8am
	daysToSubtract := (int(time.Now().Weekday()) - 4 + 7) % 7
	previousThursday := time.Now().AddDate(0, 0, -daysToSubtract)
	previousThursday8am := time.Date(previousThursday.Year(), previousThursday.Month(), previousThursday.Day(), 8, 0, 0, 0, location)

	// Check if last Publix scrape occurred before the previous Thursday at 8am EST
	publixScrapeTime, _ := d.ReadPublixScrapedTime()
	if publixScrapeTime.In(location).Before(previousThursday8am) {

		// deletes old weekly deals from .db file
		d.ClearPublixDeals()

		// scrape publix data
		s.PublixScrapeDealsPy() // put py func here
		fmt.Println("Publix Deals Scraped!")
		
		// store publix data to .db file
		d.StorePublixDatabase(s.PublixDeals)
		d.StorePubixScrapedTime(time.Now())
		
	}

	// create a time object for the first of the current month
	year, month, _ := time.Now().Date()
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)

	// check if last Walmart scrape occured over a month ago
	walmartScrapeTime, _ :=d .ReadWalmartScrapedTime()

	if walmartScrapeTime.In(location).Before(firstDayOfMonth) {
		
		// deletes old weekly deals from .db file
		d.ClearWalmartDeals()

		// scrape walmart data
		s.WalmartScrapeDealsPy();
		fmt.Println("Walmart Deals Scraped!")

		// store walmart data to .db file
		d.StoreWalmartDatabase(s.WalmartDeals)
		d.StoreWalmartScrapedTime(time.Now())

	}
}


func (s *Scraper) PublixScrapeDealsPy(){
	// run Python script to scrape Publix deals 
	name := "python3"
	if runtime.GOOS == "windows"{
		name = "python"
	}
	
	cmd := exec.Command(name, "-X", "utf8", "PublixScraper.py")
	output, _:= cmd.Output()

	outputClean := html.UnescapeString(string(output))

	// parse output into FoodItems
	lines := strings.Split(string(outputClean), "\n")
	products := make([]FoodItem, 0)

	for i := 0; i < len(lines)-1; i += 3{
		product := FoodItem{
			Name: strings.TrimPrefix(lines[i], "Product: "),
			SaleDetails: strings.TrimPrefix(lines[i+1], "Deal: "),
		}
		products = append(products, product)
	}

	s.PublixDeals = products
}

func (s *Scraper) WalmartScrapeDealsPy(){
	// run Python script to scrape Walmart deals
	name := "python3"
	if runtime.GOOS == "windows"{
		name = "python"
	}
	
	cmd := exec.Command(name, "-X", "utf8", "WalmartScraper.py")
	output, _ := cmd.Output()
	
	// parse output into FoodItems
	lines := strings.Split(string(output), "\n")
	products := make([]FoodItem, 0)

	for i := 0; i < len(lines)-1; i += 3 {
		// if statement to filter out items that are incorrectly scraped without a decimal
		price, _ := strconv.Atoi(strings.TrimPrefix(lines[i+1], "Price: $"))
		if (price < 100) && (strings.TrimPrefix(lines[i+1], "Price: $")[0] != '0'){
			product := FoodItem{
				Name:  strings.TrimPrefix(lines[i], "Product: "),
				SaleDetails: strings.TrimPrefix(lines[i+1], "Price: "),
			}
			products = append(products, product)
		} else if strings.TrimPrefix(lines[i+1], "Price: $")[0:3] == "0.0"{
			product := FoodItem{
				Name:  strings.TrimPrefix(lines[i], "Product: "),
				SaleDetails: strings.TrimPrefix(lines[i+1], "Price: "),
			}
			products = append(products, product)
		}
	}

	// store products in Scraper struct
	s.WalmartDeals = products
}