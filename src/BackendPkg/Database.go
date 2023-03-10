package BackendPkg

import (
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

type Database struct{
	Name string
}

// initializes application database file
func (d *Database) OpenDatabase() *sql.DB{
	database, _ := sql.Open("sqlite3", "./MealDealz.db")
	return database
}

// need to pass in the inventory slice from the grocery store item
func (d *Database) StorePublixDatabase (f []FoodItem){

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS FoodItems (Name TEXT, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER)")
	statement.Exec();

	//insert into food item table
	statementTwo, _ := database.Prepare("INSERT INTO FoodItems (Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, ?, ?, ?, ?, ?)")

	for _, item := range f {
		statementTwo.Exec(item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}
}

func (d *Database) StoreUserDatabase (u User){

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for user data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserData (FirstName TEXT, LastName TEXT, Email TEXT, UserName TEXT, Password TEXT)")
	statement.Exec();

	// insert into UserData table
	statementTwo, _ := database.Prepare("INSERT INTO UserData (FirstName, LastName, Email, UserName, Password) VALUES (?, ?, ?, ?, ?)")

	// store data from this user into table
	statementTwo.Exec(u.FirstName, u.LastName, u.Email, u.UserName, u.Password)
	
}