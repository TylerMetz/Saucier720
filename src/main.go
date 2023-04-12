package main

import (
	"BackendPkg"
	"fmt"
	"time"
	"context"
	//"os"
	//"os/signal"
    //"syscall"
)

// global vars
var sessionCookie string
var sessionUser BackendPkg.User
var programDatabase BackendPkg.Database
var cookieChange bool

func main() {

	// Reads recipes dataset in not read in yet and stores in DB
	programDatabase.WriteRecipes()

	// runs scraper if new deals at publix
	CheckIfScrapeNewDeals(programDatabase)

	for {
		// create a new context with a cancel function
		ctx, cancel := context.WithCancel(context.Background())

		// reset sessionCookie and cookieChanged bool
		sessionCookie = ""
		cookieChange = false

		// run program again
		go runProgram(&cookieChange, ctx)

		// do nothing while waiting for cookie to be changed
		for !cookieChange{}

		// cancel all go routines
		cancel()

		//shutdown all active ListenAndServe functions
		BackendPkg.ShutdownServers()
		
	}
}

func runProgram(cookieChange *bool, ctx context.Context) {
	// testing
	fmt.Println("New User Program running...")

	// wait for user to login and return a cookie
	go BackendPkg.ListenLogin(&sessionCookie, cookieChange, ctx)
	for sessionCookie == "" {}

	// determine session user based on cookies
	sessionUser = programDatabase.UserFromCookie(sessionCookie)

	// routs all data
	BackendPkg.RoutAllData(programDatabase, sessionUser, ctx)

	// listens for data from frontend
	BackendPkg.ListenForAllPosts(sessionUser, sessionCookie, ctx)
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