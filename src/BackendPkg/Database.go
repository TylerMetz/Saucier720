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
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS PublixData (Name TEXT PRIMARY KEY, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER)")
	statement.Exec();

	// insert into food item table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO PublixData (Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, ?, ?, ?, ?, ?)")

	for _, item := range f {
		statementTwo.Exec(item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}
}

func (d* Database) ReadPublixDatabase() []FoodItem{
	// calls function to open the database
	database := d.OpenDatabase()

	statement, _ := database.Prepare("SELECT Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity FROM PublixData")

	rows, _ := statement.Query()

	var items []FoodItem
	for rows.Next() {
		var item FoodItem
		rows.Scan(&item.Name, &item.StoreCost, &item.OnSale, &item.SalePrice, &item.SaleDetails, &item.Quantity)
    	items = append(items, item)
	}

	return items
}

// func (d* Database) ReadUserDatabase(userName string) User{
// 	// return user data from a unique username
// 	// used to validate password
// }

func (d *Database) StoreUserDatabase (u User){

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for user data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserData (FirstName TEXT, LastName TEXT, Email TEXT, UserName TEXT PRIMARY KEY, Password TEXT)")
	statement.Exec();

	// insert into UserData table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserData (FirstName, LastName, Email, UserName, Password) VALUES (?, ?, ?, ?, ?)")

	// store data from this user into table
	statementTwo.Exec(u.FirstName, u.LastName, u.Email, u.UserName, u.Password)
	
}