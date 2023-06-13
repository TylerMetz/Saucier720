package BackendPkg

import (
	"database/sql"
	"time"
	"encoding/json"
	"strings"
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

func (d *Database) ClearPublixDeals() {
	// open the database
	database := d.OpenDatabase()

	// delete the deals table if it exists
	database.Exec("DROP TABLE IF EXISTS PublixData")

	// delete the deals scraped time if it exists
	database.Exec("DROP TABLE IF EXISTS DealsScrapedTime")
}

func (d *Database) StoreWalmartDatabase(f []FoodItem) {

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS WalmartData (Name TEXT PRIMARY KEY, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER)")
	statement.Exec()

	// insert into food item table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO WalmartData (Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, ?, ?, ?, ?, ?)")

	for _, item := range f {
		statementTwo.Exec(item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}
}

func (d *Database) ReadWalmartDatabase() []FoodItem {
	// calls function to open the database
	database := d.OpenDatabase()

	statement, _ := database.Prepare("SELECT Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity FROM WalmartData")

	rows, _ := statement.Query()

	var items []FoodItem
	for rows.Next() {
		var item FoodItem
		rows.Scan(&item.Name, &item.StoreCost, &item.OnSale, &item.SalePrice, &item.SaleDetails, &item.Quantity)
		items = append(items, item)
	}

	return items
}

func (d *Database) ClearWalmartDeals() {
	// open the database
	database := d.OpenDatabase()

	// delete the deals table if it exists
	database.Exec("DROP TABLE IF EXISTS WalmartData")

	// delete the deals scraped time if it exists
	database.Exec("DROP TABLE IF EXISTS DealsScrapedTime")
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

func (d *Database) InsertPantryItemPost (currUser User, f FoodItem){

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserPantries (UserName TEXT, PantryLastUpdated DATETIME, Name TEXT, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER, PRIMARY KEY (UserName, Name))")
	statement.Exec();

	// insert into food item table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserPantries (UserName, PantryLastUpdated, Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, datetime(?), ?, ?, ?, ?, ?, ?)")
	statementTwo.Exec(currUser.UserName, time.Now().Format("2006-01-02 15:04:05"), f.Name, f.StoreCost, f.OnSale, f.SalePrice, f.SaleDetails, f.Quantity)
}

func (d *Database) UpdatePantry(currUser User, f []FoodItem){
	
	// calls function to open the database
	database := d.OpenDatabase()

	// clear all of user's current pantry
	statementOne, _ := database.Prepare("DELETE FROM UserPantries WHERE UserName = ?")
	statementOne.Exec(currUser.UserName)

	// insert all items in recieved pantry to user's pantry
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserPantries (UserName, PantryLastUpdated, Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, datetime(?), ?, ?, ?, ?, ?, ?)")
	for _, item := range f {
		statementTwo.Exec(currUser.UserName, time.Now().Format("2006-01-02 15:04:05"), item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
	}
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

func (d *Database) StorePubixScrapedTime(t time.Time) {
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

func (d *Database) ReadPublixScrapedTime() time.Time {
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

func (d *Database) StoreWalmartScrapedTime(t time.Time) {
	// calls function to open the database
	database := d.OpenDatabase()

	// make table for user data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS WalmartDealsScrapedTime (DealsLastScraped DATETIME PRIMARY KEY)")
	statement.Exec()

	// insert into UserData table
	statementTwo, _ := database.Prepare("INSERT INTO WalmartDealsScrapedTime (DealsLastScraped) VALUES (?)")

	// store data from this user into table
	statementTwo.Exec(t.Format("2006-01-02 15:04:05"))
}

func (d *Database) ReadWalmartScrapedTime() time.Time {
	// calls function to open the database
	database := d.OpenDatabase()

	// make a query to return the last scrape time value
	row := database.QueryRow("SELECT DealsLastScraped FROM WalmartDealsScrapedTime")
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

func (d* Database) DeleteRecipes(){
	// calls function to open the database
	database := d.OpenDatabase()

	// Create a new table for the recipes
	statement, _ := database.Prepare("DROP TABLE RecipeData")
	statement.Exec()

}

func (d* Database) ReadRecipes() []Recipe{
	// calls function to open the database
	database := d.OpenDatabase()

	// Execute a SELECT statement to retrieve all rows from the RecipeData table
	rows, _ := database.Query("SELECT * FROM RecipeData")

	// Iterate through the rows and create a slice of Recipe structs
	var recipes []Recipe
	for rows.Next() {
		var title, ingredientsStr, instructions string
		rows.Scan(&title, &ingredientsStr, &instructions)

		// Convert the comma-separated list of ingredients to a slice
		ingredients := strings.Split(ingredientsStr, ",")

		// Create a new Recipe struct and append it to the slice
		recipe := Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
		}
		recipes = append(recipes, recipe)
	}

	return recipes;
}

func (d* Database) GetUserPassword(username string) string{
	database := d.OpenDatabase()
	var password string 

	stmt, err := database.Prepare("SELECT Password FROM UserData WHERE UserName=?")
	if err != nil {
		// handle error
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	row.Scan(&password)

	return password
}

func (d *Database) StoreCookie(username string, cookie string) {

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for user data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS Cookies (UserName TEXT PRIMARY KEY, Cookie TEXT)")
	statement.Exec()

	// insert into UserData table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO Cookies (UserName, Cookie) VALUES (?, ?)")

	// store data from this user into table
	statementTwo.Exec(username, cookie)

}

func (d *Database) ReadCookie(username string) string {
	// return user data from a unique username
	database := d.OpenDatabase()

	var returnCookie string
	stmt, err := database.Prepare("SELECT Cookie FROM Cookies WHERE UserName=?")
	if err != nil {
		// handle error
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	row.Scan(&returnCookie)

	return returnCookie
}

func (d *Database) UserFromCookie(cookie string) User {
	// return user based off of the cookie
	database := d.OpenDatabase()
	var returnUser User
	var userName string
	stmt, err := database.Prepare("SELECT UserName FROM Cookies WHERE Cookie=?")
	if err != nil {
		// handle error
	}
	defer stmt.Close()

	row := stmt.QueryRow(cookie)
	row.Scan(&userName)

	// grabs the user based of the username
	returnUser = d.ReadUserDatabase(userName)

	return returnUser
}


