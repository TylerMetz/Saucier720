package BackendPkg

import (
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

type Database struct{
	Name string
	Store GroceryStore
}

func (d *Database) FoodItemSliceTest (f []FoodItem){
	database, _ := sql.Open("sqlite3", "./Publix.db")
	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS FoodItems (Name TEXT, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER)")
	statement.Exec();
	//insert into food item table
	statementTwo, _ := database.Prepare("INSERT INTO FoodItems (Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, ?, ?, ?, ?, ?)")

	for _, item := range f {
		statementTwo.Exec(item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}
}