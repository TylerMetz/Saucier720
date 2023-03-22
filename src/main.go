package main

import (
	"BackendPkg"
	"fmt"
	"time"
)

func main(){
	
	// test food items to test user fxns
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
	testUserFoodSlice := []BackendPkg.FoodItem{testFoodItem, testFoodItem2, testFoodItem3}

	// test database
	testDatabase := BackendPkg.Database{
		Name: "MealDealz Database",
	}

	// create a test user and store their pantry
	testUser := BackendPkg.User{
		FirstName: "Eddie",
		LastName: "Menello",
		Email: "Edward@gmail.com",
		UserName: "Eddiefye69",
		Password: "ILoveGraham420",
		UserPantry: BackendPkg.Pantry{
			FoodInPantry: testUserFoodSlice,
			TimeLastUpdated: time.Now(),
		},
	}
	testDatabase.StoreUserDatabase(testUser)
	testDatabase.StoreUserPantry(testUser)

	CheckIfScrapeNewDeals(testDatabase)
	
	RoutWeeklyDeals(testDatabase)
	go RoutUserPantry(testDatabase, testUser)
 
}

func RoutUserPantry(d BackendPkg.Database, u BackendPkg.User){
	
	// read from .db file and output test user's pantry to frontend
	var testFoodInterface []interface{}
	for i := 0; i < len(d.GetUserPantry(u.UserName).FoodInPantry); i++{
		testFoodInterface = append(testFoodInterface, d.GetUserPantry(u.UserName).FoodInPantry[i])
	}
	// test router
	programRouter := BackendPkg.Router{
		Name:             "testRouter",
		ItemsToBeEncoded: testFoodInterface,
	}
	programRouter.Rout("/api/Pantry", ":8080")
}

func CheckIfScrapeNewDeals(d BackendPkg.Database){

	// Get the current time in the EST timezone
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)

	// Calculate the previous Thursday at 8am EST
	prevThursday := now.Add(-time.Duration(now.Weekday()-time.Thursday-7) * 24 * time.Hour)
	prevThursday8am := time.Date(prevThursday.Year(), prevThursday.Month(), prevThursday.Day(), 8, 0, 0, 0, loc)
	
	// fmt.Println(d.ReadDealsScrapedTime().Format("2006-01-02 15:04:05"))

	// check if the scrape time was before the most recent thursday at 8am EST, rescrape if so
	if d.ReadDealsScrapedTime().Before(prevThursday8am) {

		// deletes old weekly deals from .db file
		d.ClearPublixDeals()

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
		//fmt.Println(programScraper.DealsHTML)

		// Takes 48634 'Words' to get to the first items name...
		// Testing to see if we can grab the name and deal from the function 
		fmt.Println("Finished Scraping")

		//Print the scraper data
		//fmt.Println(programScraper.DealsHTML)

		testFoodSlice := programScraper.Store.OrganizeDeals(programScraper.DealsHTML, 48640)
		
		// store publix data to .db file
		d.StorePublixDatabase(testFoodSlice)
		d.StoreDealsScrapedTime(programScraper.TimeLastDealsScraped)
	}
}

func RoutWeeklyDeals(d BackendPkg.Database){
	
	// read from .db file and output test user's pantry to frontend
	var testFoodInterface []interface{}
	for i := 0; i < len(d.ReadPublixDatabase()); i++{
		testFoodInterface = append(testFoodInterface, d.ReadPublixDatabase()[i])
	}
	// test router
	programRouter := BackendPkg.Router{
		Name:             "testRouter",
		ItemsToBeEncoded: testFoodInterface,
	}
	programRouter.Rout("/api/Deals", ":8080")
}