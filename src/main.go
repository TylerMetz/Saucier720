package main

import (
	"BackendPkg"
	"fmt"
	"time"
	"context"
)

// global vars
var sessionCookie string
var cookieChanged bool
var sessionUser BackendPkg.User
var programDatabase BackendPkg.Database
var prevUser BackendPkg.User

func main() {

	// Reads recipes dataset in not read in yet and stores in DB
	programDatabase.WriteRecipes()

	// runs scraper if new deals at publix
	CheckIfScrapeNewDeals(programDatabase)

	// create a new context object
	ctx, cancel := context.WithCancel(context.Background())

	// listen for user in a separate goroutine, and wait for session cookie to be defined
	go BackendPkg.ListenForUser(ctx, &sessionCookie, &cookieChanged)
	for sessionCookie == "" {}

	// always check if cookie is changed
	go func(){
		for{
			if(cookieChanged){
				// determine session user based on cookies
				for(BackendPkg.CurrentUser.UserName == prevUser.UserName){
					BackendPkg.CurrentUser = programDatabase.UserFromCookie(sessionCookie)
				}
				// store prev user 
				prevUser = BackendPkg.CurrentUser

				// reset cookie change
				cookieChanged = false
			}
		}
	}()
	// rout and listen for all data actively with the defined session user
	go BackendPkg.RoutData(ctx)
	go BackendPkg.ListenForData(ctx)

	// run infinitely
	select{}

	cancel()
	
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

		// Testing to see if we can grab the name and deal from the function 
		fmt.Println("Finished Scraping")

		testFoodSlice := programScraper.Store.OrganizeDeals(programScraper.DealsHTML)
		
		// store publix data to .db file
		d.StorePublixDatabase(testFoodSlice)
		d.StoreDealsScrapedTime(programScraper.TimeLastDealsScraped)
	}
}