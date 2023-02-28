package main

import (
	"BackendPkg"
	"fmt"
)

func main() {
	/*

		testFoodItem := BackendPkg.FoodItem{
			Name:        "peanut butter",
			StoreCost:   369.99,
			OnSale:      true,
			SaleDetails: "BOGO",
			Quantity:    10,
		}
		testFoodItem2 := BackendPkg.FoodItem{
			Name:        "jelly",
			StoreCost:   1.0,
			OnSale:      false,
			SaleDetails: "N/A",
			Quantity:    30,
		}
		testFoodItem3 := BackendPkg.FoodItem{
			Name:        "bread",
			StoreCost:   10.69,
			OnSale:      true,
			SaleDetails: "$2 for 2",
			Quantity:    2,
		}

		testFoodSlice := []interface{}{testFoodItem, testFoodItem2, testFoodItem3}

		programRouter := BackendPkg.Router{
			Name:             "testRouter",
			ItemsToBeEncoded: testFoodSlice,
		}
		programRouter.Rout()

	*/
	// runs scraper
	runScraper := false
	if runScraper {
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

		// print unparsed data
		fmt.Println(programScraper.DealsHTML)
	}
}
