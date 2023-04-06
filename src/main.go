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
	testFoodItem4 := BackendPkg.FoodItem{
		Name:        "eggs",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testFoodItem5 := BackendPkg.FoodItem{
		Name:        "milk",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testFoodItem6 := BackendPkg.FoodItem{
		Name:        "sugar",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testFoodItem7 := BackendPkg.FoodItem{
		Name:        "flour",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testFoodItem8 := BackendPkg.FoodItem{
		Name:        "vanilla",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testFoodItem9 := BackendPkg.FoodItem{
		Name:        "butter",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testFoodItem10 := BackendPkg.FoodItem{
		Name:        "apple",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testUserFoodSlice := []BackendPkg.FoodItem{testFoodItem, testFoodItem2, testFoodItem3, testFoodItem4, testFoodItem5, testFoodItem6, testFoodItem7, testFoodItem8, testFoodItem9, testFoodItem10}

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

	testUserTwo := BackendPkg.User{
		FirstName: "Sam",
		LastName: "Forsnot",
		Email: "samuel@gmail.com",
		UserName: "SameHatesBigWordsXXX",
		Password: "ILoveJess420",
		UserPantry: BackendPkg.Pantry{
			FoodInPantry: testUserFoodSlice,
			TimeLastUpdated: time.Now(),
		},
	}

	// store Eddie
	testDatabase.StoreUserDatabase(testUser)
	testDatabase.StoreUserPantry(testUser)

	// store Eddie version of Sam
	testDatabase.StoreUserDatabase(testUserTwo)
	testDatabase.StoreUserPantry(testUserTwo)

	// Reads recipes dataset
	testDatabase.WriteRecipes()

	// runs scraper if new deals at publix
	CheckIfScrapeNewDeals(testDatabase)

	// routs deals to deals page
	go RoutWeeklyDeals(testDatabase)

	// routs Eddie's pantry, lol
	go RoutUserPantry(testDatabase, testUser)

	testReccList := BackendPkg.BestRecipes(testDatabase.GetUserPantry("Eddiefye69"), testDatabase.ReadRecipes(), testDatabase.ReadPublixDatabase())

	for i := 0; i < len(testReccList); i++{
		fmt.Println(testReccList[i].R.Title)
		for j := 0; j < len(testReccList[i].ItemsInPantry); j++{
			fmt.Println(testReccList[i].ItemsInPantry[j].Name)
		}
	}

	ListenForPost(testDatabase);
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

	// Set the location to Eastern Standard Time (EST)
	est, _ := time.LoadLocation("America/New_York")

	// Get the current time in EST
	now := time.Now().In(est)

	// Get the previous Thursday at 8am EST
	previousThursday := now.AddDate(0, 0, -int(now.Weekday()+3)%7)
	previousThursday8am := time.Date(previousThursday.Year(), previousThursday.Month(), previousThursday.Day(), 8, 0, 0, 0, est)

	// Check if scrapeTime occurred before the previous Thursday at 8am EST
	if d.ReadDealsScrapedTime().In(est).Before(previousThursday8am) {

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

		testFoodSlice := programScraper.Store.OrganizeDeals(programScraper.DealsHTML)
		
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
	go programRouter.Rout("/api/Deals", ":8081")
}

func ListenForPost(d BackendPkg.Database){
	programRouter2 := BackendPkg.Router{
		Name:             "testRouter",
	}
	programRouter2.Listen("/api/NewPantryItem", ":8082") //should pass in database here
}