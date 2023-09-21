package BackendPkg

import (
	_ "fmt"
	_ "html"
	_ "os/exec"
	_ "runtime"
	_ "strconv"
	_ "strings"
	_ "time"
	_"unicode/utf8"
)

var UserZipCodePlaceholder string = "32601"

type Scraper struct {
	PublixDeals			[]FoodItem
	WalmartDeals 		[]FoodItem
}

// func (s *Scraper) PublixScrapeDealsPy(){
// 	// run Python script to scrape Publix deals 
// 	name := "python3"
// 	if runtime.GOOS == "windows"{
// 		name = "python"
// 	}
	
// 	cmd := exec.Command(name, "-X", "utf8", "PublixScraper.py")
// 	output, _:= cmd.Output()

// 	outputClean := html.UnescapeString(string(output))

// 	// parse output into FoodItems
// 	lines := strings.Split(string(outputClean), "\n")
// 	products := make([]FoodItem, 0)

// 	for i := 0; i < len(lines)-1; i += 3{
// 		product := FoodItem{
// 			Name: strings.TrimPrefix(lines[i], "Product: "),
// 			SaleDetails: strings.TrimPrefix(lines[i+1], "Deal: "),
// 		}
// 		products = append(products, product)
// 	}

// 	s.PublixDeals = products
// }

// func (s *Scraper) WalmartScrapeDealsPy(){
// 	// run Python script to scrape Walmart deals
// 	name := "python3"
// 	if runtime.GOOS == "windows"{
// 		name = "python"
// 	}
	
// 	cmd := exec.Command(name, "-X", "utf8", "WalmartScraper.py")
// 	output, _ := cmd.Output()
	
// 	// parse output into FoodItems
// 	lines := strings.Split(string(output), "\n")
// 	products := make([]FoodItem, 0)

// 	for i := 0; i < len(lines)-1; i += 3 {
// 		// if statement to filter out items that are incorrectly scraped without a decimal
// 		price, _ := strconv.Atoi(strings.TrimPrefix(lines[i+1], "Price: $"))
// 		if (price < 100) && (strings.TrimPrefix(lines[i+1], "Price: $")[0] != '0'){
// 			product := FoodItem{
// 				Name:  strings.TrimPrefix(lines[i], "Product: "),
// 				SaleDetails: strings.TrimPrefix(lines[i+1], "Price: "),
// 			}
// 			products = append(products, product)
// 		} else if strings.TrimPrefix(lines[i+1], "Price: $")[0:3] == "0.0"{
// 			product := FoodItem{
// 				Name:  strings.TrimPrefix(lines[i], "Product: "),
// 				SaleDetails: strings.TrimPrefix(lines[i+1], "Price: "),
// 			}
// 			products = append(products, product)
// 		}
// 	}

// 	// store products in Scraper struct
// 	s.WalmartDeals = products
// }
