package BackendPkg

import (
	"database/sql"
	"time"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	//"fmt"
)

type Database struct {
	Name string
}

// initializes application database file
func (d *Database) OpenDatabase() *sql.DB {
	database, _ := sql.Open("sqlite3", "./MealDealz.db")
	return database
}

// need to pass in the inventory slice from the grocery store item
func (d *Database) StorePublixDatabase(f []FoodItem) {

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS PublixData (Name TEXT PRIMARY KEY, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER)")
	statement.Exec()

	// insert into food item table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO PublixData (Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, ?, ?, ?, ?, ?)")

	for _, item := range f {
		statementTwo.Exec(item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}
}

func (d *Database) ReadPublixDatabase() []FoodItem {
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

func (d *Database) ReadUserDatabase(userName string) User {
	// return user data from a unique username
	// used to validate password
	database := d.OpenDatabase()

	var returnUser User

	stmt, err := database.Prepare("SELECT FirstName, LastName, Email, UserName, Password FROM UserData WHERE UserName=?")
	if err != nil {
		// handle error
	}
	defer stmt.Close()

	row := stmt.QueryRow(userName)
	row.Scan(&returnUser.FirstName, &returnUser.LastName, &returnUser.Email, &returnUser.UserName, &returnUser.Password)

	return returnUser

}

func (d *Database) StoreUserDatabase(u User) {

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for user data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserData (FirstName TEXT, LastName TEXT, Email TEXT, UserName TEXT PRIMARY KEY, Password TEXT)")
	statement.Exec()

	// insert into UserData table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserData (FirstName, LastName, Email, UserName, Password) VALUES (?, ?, ?, ?, ?)")

	// store data from this user into table
	statementTwo.Exec(u.FirstName, u.LastName, u.Email, u.UserName, u.Password)

}

func (d *Database) StoreUserPantry(u User) {

	// calls function to open the database
	database := d.OpenDatabase() // need to open database

	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserPantries (UserName TEXT, PantryLastUpdated DATETIME, Name TEXT, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER, PRIMARY KEY (UserName, Name))")
	statement.Exec()

	// insert into food item table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserPantries (UserName, PantryLastUpdated, Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, datetime(?), ?, ?, ?, ?, ?, ?)")

	for _, item := range u.UserPantry.FoodInPantry {
		statementTwo.Exec(u.UserName, u.UserPantry.TimeLastUpdated.Format("2006-01-02 15:04:05"), item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}
}

func (d *Database) InsertPantryItemPost (u FoodItem){

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserPantries (UserName TEXT, PantryLastUpdated DATETIME, Name TEXT, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER, PRIMARY KEY (UserName, Name))")
	statement.Exec();

	var exampleUser = "RileySmellsLol"
	var exampleTime = "2023-03-24 22:56:28"

	// insert into food item table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserPantries (UserName, PantryLastUpdated, Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, datetime(?), ?, ?, ?, ?, ?, ?)")
	statementTwo.Exec(exampleUser, exampleTime, u.Name, u.StoreCost, u.OnSale, u.SalePrice, u.SaleDetails, u.Quantity)
}

func (d *Database) GetUserPantry(userName string) Pantry {
	// calls function to open the database
	database := d.OpenDatabase()

	// query the database for the pantry data
	rows, _ := database.Query("SELECT Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity, PantryLastUpdated FROM UserPantries WHERE UserName = ?", userName)

	// create the pantry object
	pantry := Pantry{
		TimeLastUpdated: time.Now(),
		FoodInPantry:    []FoodItem{},
	}

	// loop through each row and add the food item to the pantry
	for rows.Next() {
		var name, saleDetails string
		var storeCost, salePrice float64
		var onSale bool
		var quantity int
		var pantryLastUpdated string
		if err := rows.Scan(&name, &storeCost, &onSale, &salePrice, &saleDetails, &quantity, &pantryLastUpdated); err != nil {
			return Pantry{}
		}
		pantry.FoodInPantry = append(pantry.FoodInPantry, FoodItem{
			Name:        name,
			StoreCost:   storeCost,
			OnSale:      onSale,
			SalePrice:   salePrice,
			SaleDetails: saleDetails,
			Quantity:    quantity,
		})

		// set the time last updated
		pantry.TimeLastUpdated, _ = time.Parse("2006-01-02 15:04:05", pantryLastUpdated)

	}

	return pantry
}

func (d *Database) ClearPublixDeals() {
	// open the database
	database := d.OpenDatabase()

	// delete the deals table if it exists
	database.Exec("DROP TABLE IF EXISTS PublixData")

	// delete the deals scraped time if it exists
	database.Exec("DROP TABLE IF EXISTS DealsScrapedTime")
}

func (d *Database) StoreDealsScrapedTime(t time.Time) {
	// calls function to open the database
	database := d.OpenDatabase()

	// make table for user data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS DealsScrapedTime (DealsLastScraped DATETIME PRIMARY KEY)")
	statement.Exec()

	// insert into UserData table
	statementTwo, _ := database.Prepare("INSERT INTO DealsScrapedTime (DealsLastScraped) VALUES (?)")

	// store data from this user into table
	statementTwo.Exec(t.Format("2006-01-02 15:04:05"))
}

func (d *Database) ReadDealsScrapedTime() time.Time {
	// calls function to open the database
	database := d.OpenDatabase()

	// make a query to return the last scrape time value
	row := database.QueryRow("SELECT DealsLastScraped FROM DealsScrapedTime")
	var dealsLastScrapedStr string
	row.Scan(&dealsLastScrapedStr)

	// Parse the datetime string into a time.Time object
	dealsLastScraped, _ := time.Parse(time.RFC3339, dealsLastScrapedStr)

	return dealsLastScraped
}

func (d* Database) WriteRecipes(){
	// Read the recipes from the file
	recipes, _ := GetRecipes()

	// calls function to open the database
	database := d.OpenDatabase()

	// Create a new table for the recipes
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS RecipeData (title TEXT PRIMARY KEY, ingredients TEXT, instructions TEXT)")
	statement.Exec()

	// Insert each recipe into the table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO RecipeData (title, ingredients, instructions) values (?, ?, ?)")

	for _, recipe := range recipes {
		ingredients, _ := json.Marshal(recipe.Ingredients)
		statementTwo.Exec(recipe.Title, string(ingredients), recipe.Instructions)
	}

	database.Exec("DELETE FROM RecipeData WHERE ingredients = '[]'")

}