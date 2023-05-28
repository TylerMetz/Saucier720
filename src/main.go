package main

import (
	"BackendPkg"
	//"fmt"
	"time"
)

// global vars
var sessionCookie string
var prevCookie string
var cookieChanged bool
var programDatabase BackendPkg.Database
var programScraper BackendPkg.Scraper
var prevUser BackendPkg.User

func main() {

	// Reads recipes dataset in not read in yet and stores in DB
	programDatabase.WriteRecipes()

	// runs scraper if new deals at publix
	CheckIfScrapeNewDeals(programDatabase)

	// listen for user in a separate goroutine, and wait for session cookie to be defined
	go BackendPkg.ListenForUser(&sessionCookie, &cookieChanged)
	for sessionCookie == "" && !cookieChanged {}

	// always check if cookie is changed
	go func(){
		for{
			if(cookieChanged){
				// determine session user based on cookies
				for(BackendPkg.CurrentUser.UserName == prevUser.UserName){
					BackendPkg.CurrentUser = programDatabase.UserFromCookie(sessionCookie)
					if(prevCookie == sessionCookie){
						BackendPkg.CurrentUser = programDatabase.UserFromCookie(sessionCookie)
						break;
					}
				}
				// store prev user 
				prevUser = BackendPkg.CurrentUser

				// reset cookie change
				cookieChanged = false
			}
		}
	}()

	
	// rout and listen for all data actively with the defined session user
	go BackendPkg.RoutData()
	go BackendPkg.ListenForData()

	// goroutine to set the previous cookie to the session cookie while the session cookie isn't being changed
	go func(){
		for{
			if(!cookieChanged){
				prevCookie = sessionCookie
			}
		}
	}()

	// run infinitely
	for{}
	
}

func CheckIfScrapeNewDeals(d BackendPkg.Database){

	// for testing
	d.ClearWalmartDeals()

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
		d.ClearWalmartDeals()

		// scrape all data
		programScraper.Scrape()
		
		// store publix data to .db file
		d.StorePublixDatabase(programScraper.PublixDeals)
		d.StoreWalmartDatabase(programScraper.WalmartDeals)
		d.StoreDealsScrapedTime(programScraper.TimeLastDealsScraped)
	}
}