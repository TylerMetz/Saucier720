package main

import (
	"BackendPkg"
	"fmt"
	"time"
)

func main(){
	
	// testFoodItem := BackendPkg.FoodItem{
	// 	Name:        "peanut butter",
	// 	StoreCost:   369.99,
	// 	OnSale:      true,
	// 	SaleDetails: "BOGO",
	// 	Quantity:    10,
	// }
	// testFoodItem2 := BackendPkg.FoodItem{
	// 	Name:        "jelly",
	// 	StoreCost:   1.0,
	// 	OnSale:      false,
	// 	SaleDetails: "N/A",
	// 	Quantity:    30,
	// }
	// testFoodItem3 := BackendPkg.FoodItem{
	// 	Name:        "bread",
	// 	StoreCost:   10.69,
	// 	OnSale:      true,
	// 	SaleDetails: "$2 for 2",
	// 	Quantity:    2,
	// }

	//testFoodSlice := []BackendPkg.FoodItem{testFoodItem, testFoodItem2, testFoodItem3}
	

	// test scraper
	//Putting variables outsde for the sake of testing db
	userPublix := BackendPkg.GroceryStore{
		Name:    "Publix",
		ZipCode: "32601",
	}

	// setup user groccery store
	programScraper := BackendPkg.Scraper{
		Store: userPublix,
	}
	runScraper := true
	if runScraper {
		// create new groccery store
		/*userPublix := BackendPkg.GroceryStore{
			Name:    "Publix",
			ZipCode: "32601",
		}

		// setup user groccery store
		programScraper := BackendPkg.Scraper{
			Store: userPublix,
		}*/

		// scrape all data
		programScraper.Scrape()

		// print unparsed data
		//fmt.Println(programScraper.DealsHTML)

		// Takes 48634 'Words' to get to the first items name...
		// Testing to see if we can grab the name and deal from the function 
		//testFoodSlice = programScraper.Store.OrganizeDeals(programScraper.DealsHTML, 48634)
		fmt.Println("Finished")

		//Print the scraper data
		//fmt.Println(programScraper.DealsHTML)
	}

	testFoodSlice := programScraper.Store.OrganizeDeals(programScraper.DealsHTML, 48640)

	// test database
	testDatabase := BackendPkg.Database{
		Name: "MealDealz Database",
	}
	// store publix data to .db file
	testDatabase.StorePublixDatabase(testFoodSlice)

	// create a test user and store their pantry
	testUser := BackendPkg.User{
		FirstName: "Eddie",
		LastName: "Menello",
		Email: "Edward@gmail.com",
		UserName: "Eddiefye69",
		Password: "ILoveGraham420",
		UserPantry: BackendPkg.Pantry{
			FoodInPantry: testFoodSlice,
			TimeLastUpdated: time.Now(),
		},
	}
	testDatabase.StoreUserDatabase(testUser)
	testDatabase.StoreUserPantry(testUser)

	// read from .db file and output test user's pantry to frontend
	var testFoodInterface []interface{}
	for i := 0; i < len(testDatabase.GetUserPantry(testUser.UserName).FoodInPantry); i++{
		testFoodInterface = append(testFoodInterface, testDatabase.GetUserPantry(testUser.UserName).FoodInPantry[i])
	}

	// test router
	programRouter := BackendPkg.Router{
		Name:             "testRouter",
		ItemsToBeEncoded: testFoodInterface,
	}
	programRouter.Rout()

}