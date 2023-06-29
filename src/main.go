package main

import (
	"BackendPkg"
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

	// runs scraper if new deals at publix or walmart
	programScraper.CheckIfScrapeNewDeals(programDatabase)

	// listen for user in a separate goroutine, and wait for session cookie to be defined
	go BackendPkg.ListenUserInfo(&sessionCookie, &cookieChanged)
	for sessionCookie == "" && !cookieChanged {}
	
	// always check if cookie is changed
	go func(){
		for{
			if(cookieChanged){
				// determine session user based on cookies
				for(BackendPkg.CurrentUser.UserName == prevUser.UserName){
					if sessionCookie != "" {
						BackendPkg.CurrentUser = programDatabase.UserFromCookie(sessionCookie)
						if(prevCookie == sessionCookie){
							BackendPkg.CurrentUser = programDatabase.UserFromCookie(sessionCookie)
							break;
						}
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
