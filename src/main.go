package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"BackendPkg"
)

func main(){
	
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

	testFoodSlice := []BackendPkg.FoodItem{testFoodItem, testFoodItem2, testFoodItem3}
	
	
	
	database, _ := sql.Open("sqllite3", "./Publix.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS FoodItems (Name TEXT, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER)")
	statement.Exec();
	statementTwo, _ := database.Prepare("INSERT INTO FoodItems (Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, ?, ?, ?, ?, ?)")

	for _, item := range testFoodSlice {
		statementTwo.Exec(item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}

	//rows, _ := database.Query("SELECT Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity FROM FoodItems")

}