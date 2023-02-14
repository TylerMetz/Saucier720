package main

import (
	"BackendPkg"
	//"fmt"
	//"time"
)

func main() {
	/*
		fmt.Println("Welcome to out Sprint 1 demo!")
		userPantry := BackendPkg.Pantry{
			TimeLastUpdated: time.Now(),
		}
		userPantry.AddToPantry()
		userPantry.AddToPantry()
		userPantry.AddToPantry()
		userPantry.DisplayPantry()
		userPantry.RemoveFromPantry()
		userPantry.RemoveFromPantry()
		userPantry.DisplayPantry()
	*/

	localPublix := BackendPkg.GroceryStore{
		Name: "Publix",
	}
	localPublix.ScrapeDeals()
	testItem := BackendPkg.FoodItem{
		Name:        "test name",
		SaleDetails: "Test details",
	}
	localPublix.Inventory = append(localPublix.Inventory, testItem)
	localPublix.DisplaySales()
}
