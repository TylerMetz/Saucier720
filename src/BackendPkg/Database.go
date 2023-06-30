package BackendPkg

import (
	"database/sql"
	"time"
	"encoding/json"
	"strings"
	_ "github.com/microsoft/go-mssqldb"
	"log"
	"strconv"
	"fmt"
	"context"
	"errors"
)

var db *sql.DB
var server = "mealdealz.database.windows.net"
var port = 1433
var user = "mealdealz-dev"
var password = "Babayaga720"
var database = "MealDealz-db"

type Database struct {
	Name string
}

// initializes application database file OLD FUNCTION
func (d *Database) OpenDatabase() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, database)
	database, _ := sql.Open("sqlserver", connString)
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

	// close db
	database.Close()
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

	// close db
	database.Close()

	return items
}

func (d *Database) ClearPublixDeals() {
	// open the database
	database := d.OpenDatabase()

	// delete the deals table if it exists
	database.Exec("DROP TABLE IF EXISTS PublixData")

	// delete the deals scraped time if it exists
	database.Exec("DROP TABLE IF EXISTS PublixDealsScrapedTime")

	// close db
	database.Close()
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

	// close db
	database.Close()
}

func (d *Database) ReadWalmartDatabase() []FoodItem {
	// calls function to open the database
	database := d.OpenDatabase()

	statement, err := database.Prepare("SELECT Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity FROM WalmartData")
	if err != nil {
		// handle the error, e.g., log or return an empty list
		log.Println("Failed to prepare statement:", err)
		return []FoodItem{}
	}

	rows, err := statement.Query()
	if err != nil {
		// handle the error, e.g., log or return an empty list
		log.Println("Failed to execute query:", err)
		return []FoodItem{}
	}

	var items []FoodItem
	for rows.Next() {
		var item FoodItem
		err := rows.Scan(&item.Name, &item.StoreCost, &item.OnSale, &item.SalePrice, &item.SaleDetails, &item.Quantity)
		if err != nil {
			// handle the error, e.g., log or skip this row
			log.Println("Failed to scan row:", err)
			continue
		}
		items = append(items, item)
	}

	// close db
	database.Close()

	return items
}

func (d *Database) ClearWalmartDeals() {
	// open the database
	database := d.OpenDatabase()

	// delete the deals table if it exists
	database.Exec("DROP TABLE IF EXISTS WalmartData")

	// delete the deals scraped time if it exists
	database.Exec("DROP TABLE IF EXISTS WalmartDealsScrapedTime")

	// close db
	database.Close()
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

	// close db
	database.Close()

	return returnUser

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

	// close db
	database.Close()
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

	// close db
	database.Close()
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

	// close db
	database.Close()
}

func (d *Database) GetUserPantry(userName string) Pantry {
	// calls function to open the database
	database := d.OpenDatabase()

	// create the pantry object
	pantry := Pantry{
		TimeLastUpdated: time.Now(),
		FoodInPantry:    []FoodItem{},
	}

	// query the database for the pantry data
	rows, err := database.Query("SELECT Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity, PantryLastUpdated FROM UserPantries WHERE UserName = ?", userName)

	// handle case where user pantry doesn't exist
	if err != nil{
		return pantry
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

	// close db
	database.Close()

	return pantry
}

func (d *Database) StorePubixScrapedTime(t time.Time) {
	// calls function to open the database
	database := d.OpenDatabase()

	// make table for user data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS PublixDealsScrapedTime (DealsLastScraped DATETIME PRIMARY KEY)")
	statement.Exec()

	// insert into UserData table
	statementTwo, _ := database.Prepare("INSERT INTO PublixDealsScrapedTime (DealsLastScraped) VALUES (?)")

	// store data from this user into table
	statementTwo.Exec(t.Format("2006-01-02 15:04:05"))

	// close db
	database.Close()
}

func (d *Database) ReadPublixScrapedTime() time.Time {
	// calls function to open the database
	database := d.OpenDatabase()

	// make a query to return the last scrape time value
	row := database.QueryRow("SELECT DealsLastScraped FROM PublixDealsScrapedTime")
	var dealsLastScrapedStr string
	row.Scan(&dealsLastScrapedStr)

	// Parse the datetime string into a time.Time object
	dealsLastScraped, _ := time.Parse(time.RFC3339, dealsLastScrapedStr)

	// close db
	database.Close()

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

	// close db
	database.Close()
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

	// close db
	database.Close()

	return dealsLastScraped
}

func (d* Database) WriteJSONRecipes(){
	// Read the recipes from the file
	recipes, _ := GetJSONRecipes()

	// calls function to open the database
	database := d.OpenDatabase()

	// Create a new table for the recipes
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS JSONRecipeData (title, ingredients TEXT, instructions TEXT, recipeID TEXT PRIMARY KEY)")
	statement.Exec()

	// Insert each recipe into the table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO JSONRecipeData (title, ingredients, instructions, recipeID) values (?, ?, ?, ?)")

	idNum := 1

	for _, recipe := range recipes {
		ingredients, _ := json.Marshal(recipe.Ingredients)
		statementTwo.Exec(recipe.Title, string(ingredients), recipe.Instructions, ("json" + strconv.Itoa(idNum)))
	
		idNum++
	}

	

	database.Exec("DELETE FROM JSONRecipeData WHERE ingredients = '[]'")
	database.Exec("DELETE FROM JSONRecipeData WHERE instructions = ''")

	// close db
	database.Close()
}

func (d* Database) DeleteJSONRecipes(){
	// calls function to open the database
	database := d.OpenDatabase()

	// Create a new table for the recipes
	statement, _ := database.Prepare("DROP TABLE JSONRecipeData")
	statement.Exec()

	// close db
	database.Close()

}

func (d* Database) ReadJSONRecipes() []Recipe{
	// calls function to open the database
	database := d.OpenDatabase()

	// Execute a SELECT statement to retrieve all rows from the RecipeData table
	rows, _ := database.Query("SELECT * FROM JSONRecipeData")

	// Iterate through the rows and create a slice of Recipe structs
	var recipes []Recipe
	for rows.Next() {
		var title, ingredientsStr, instructions, recipeID string
		rows.Scan(&title, &ingredientsStr, &instructions, &recipeID)

		// Convert the comma-separated list of ingredients to a slice
		ingredients := strings.Split(ingredientsStr, ",")

		// Create a new Recipe struct and append it to the slice
		recipe := Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeID: 	  recipeID, 
		}
		recipes = append(recipes, recipe)
	}

	// close db
	database.Close()

	return recipes;
}

func (d *Database) WriteNewUserRecipe (currUser User, newRecipe Recipe){
	// calls function to open the database
	database := d.OpenDatabase()

	// Create a new table for the recipes
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserRecipeData (title TEXT, ingredients TEXT, instructions TEXT, recipeID TEXT PRIMARY KEY, username TEXT)")
	statement.Exec()

	// find num of recipes made by user to generaate userID
	idNum := 0
	err := database.QueryRow("SELECT COUNT(*) FROM UserRecipeData WHERE username = ?", currUser.UserName).Scan(&idNum)
	if err != nil {idNum = 1 } // if user has no recipes then idNum = 1

	// Insert each recipe into the table
	ingredients, _ := json.Marshal(newRecipe.Ingredients)
	statementThree, _ := database.Prepare("INSERT OR IGNORE INTO UserRecipeData (title, ingredients, instructions, recipeID, username) values (?, ?, ?, ?, ?)")
	statementThree.Exec(newRecipe.Title, string(ingredients), newRecipe.Instructions, (currUser.UserName + strconv.Itoa(idNum)), currUser.UserName)

	// close db
	database.Close()
}

func (d *Database) DeleteUserRecipe (recipeID string){
	// Calls function to open the database
	database := d.OpenDatabase()

	// Delete based on recipeID
	statement, _ := database.Prepare("DELETE FROM UserRecipeData WHERE recipeID = ?")
	statement.Exec(recipeID)

	// close db
	database.Close()

}

func (d* Database) ReadAllUserRecipes() []Recipe{
	// calls function to open the database
	database := d.OpenDatabase()

	// create the recipes return slice
	var recipes []Recipe

	// Execute a SELECT statement to retrieve all rows from the RecipeData table
	rows, err := database.Query("SELECT * FROM UserRecipeData")
	// handle case where user recipes don't exist
	if err != nil{
		// close db
		database.Close()
		return recipes
	}

	// Iterate through the rows and create a slice of Recipe structs
	for rows.Next() {
		var title, ingredientsStr, instructions, recipeID, userName string
		rows.Scan(&title, &ingredientsStr, &instructions, &recipeID, &userName)

		// Convert the comma-separated list of ingredients to a slice
		ingredients := strings.Split(ingredientsStr, ",")

		// Create a new Recipe struct and append it to the slice
		recipe := Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeID: 	  recipeID, 
		}
		recipes = append(recipes, recipe)
	}

	// close db
	database.Close()

	return recipes;
}

func (d* Database) ReadCurrUserRecipes (currUser User) []Recipe{
	
	// calls function to open the database
	database := d.OpenDatabase()

	// create the recipes return slice
	var recipes []Recipe

	// Execute a SELECT statement to retrieve all rows from the RecipeData table
	statement, err := database.Prepare("SELECT title, ingredients, instructions, recipeID, username FROM UserRecipeData WHERE username = ?")
	// handle case where user recipes don't exist
	if err != nil{
		database.Close()
		return recipes
	}

	rows, err := statement.Query(currUser.UserName)
	// handle case where user recipes don't exist
	if err != nil{
		database.Close()
		return recipes
	}

	// Iterate through the rows and create a slice of Recipe structs
	for rows.Next() {
		var title, ingredientsStr, instructions, recipeID, userName string
		rows.Scan(&title, &ingredientsStr, &instructions, &recipeID, &userName)

		// Convert the comma-separated list of ingredients to a slice
		ingredients := strings.Split(ingredientsStr, ",")

		// Create a new Recipe struct and append it to the slice
		recipe := Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeID: 	  recipeID, 
		}
		recipes = append(recipes, recipe)
	}

	// close db
	database.Close()

	return recipes;
}

func (d* Database) FavoriteRecipe (currUser User, recipeID string){
	// calls function to open the database
	database := d.OpenDatabase()

	// Create a new table for the recipes
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserFavoriteRecipes (recipeID TEXT PRIMARY KEY, username TEXT)")
	statement.Exec()

	// save username and favorited recipe's recipe ID
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserFavoriteRecipes (recipeID, username) values (?, ?)")
	statementTwo.Exec(recipeID, currUser.UserName)

	// close db
	database.Close()
}

func (d* Database) UnfavoriteRecipe (currUser User, recipeID string){
	// calls function to open the database
	database := d.OpenDatabase()

	// delete favorite recipe from table
	statement, _ := database.Prepare("DELETE FROM UserFavoriteRecipes WHERE username = ? AND recipeID = ?")
	statement.Exec(currUser.UserName, recipeID)

	// close db
	database.Close()

}

func (d* Database) ReadFavoriteRecipes (currUser User) []Recipe{
	
	// calls function to open the database
	database := d.OpenDatabase()

	var recipes []Recipe

	// Retrieve recipeIDs from UserFavoriteRecipes table based on the given username
	statement, err := database.Prepare("SELECT recipeID FROM UserFavoriteRecipes WHERE username = ?")
	if err != nil {
		database.Close()
		return recipes
	}

	// Execute the statement
	rows, err := statement.Query(currUser.UserName)
	if err != nil {
		database.Close()
		return recipes
	}

	var recipeIDs []string

	for rows.Next() {
		var recipeID string
		err := rows.Scan(&recipeID)
		if err != nil {
			database.Close()
			return recipes
		}
		recipeIDs = append(recipeIDs, recipeID)
	}

	for _, recipeID := range recipeIDs {
		// loop through all recipeIDs from a user ravorites and retrieve the recipe
		recipe, _ := getRecipeByID(database, recipeID)
		if recipe != nil {
			recipes = append(recipes, *recipe)
		}
	}

	// close db and return recipes
	database.Close()
	return recipes

}

func getRecipeByID(db *sql.DB, recipeID string) (*Recipe, error) {
	recipe, _ := getRecipeFromIdUserTable(db, recipeID)

	if recipe == nil {
		recipeJSON, err := getRecipeFromIdJSONTable(db, recipeID)
		if err != nil {
			return nil, err
		}
		recipe = recipeJSON
	}

	return recipe, nil
}

func getRecipeFromIdUserTable(db *sql.DB, recipeID string) (*Recipe, error) {
	query := "SELECT title, ingredients, instructions FROM UserRecipeData WHERE recipeID = ?"
	row := db.QueryRow(query, recipeID)

	var title, ingredientsStr, instructions string
	err := row.Scan(&title, &ingredientsStr, &instructions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Recipe not found in UserRecipeData
		}
		return nil, err
	}

	// Convert the comma-separated list of ingredients to a slice
	ingredients := strings.Split(ingredientsStr, ",")

	// define recipe object based on returned values
	recipe := Recipe{
		Title:        title,
		Ingredients:  ingredients,
		Instructions: instructions,
		RecipeID:     recipeID,
	}

	return &recipe, nil
}

func getRecipeFromIdJSONTable(db *sql.DB, recipeID string) (*Recipe, error) {
	query := "SELECT title, ingredients, instructions FROM JSONRecipeData WHERE recipeID = ?"
	row := db.QueryRow(query, recipeID)

	var title, ingredientsStr, instructions string
	err := row.Scan(&title, &ingredientsStr, &instructions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Recipe not found in JSONRecipeData
		}
		return nil, err
	}

	// Convert the comma-separated list of ingredients to a slice
	ingredients := strings.Split(ingredientsStr, ",")

	// define recipe object based on returned values
	recipe := Recipe{
		Title:        title,
		Ingredients:  ingredients,
		Instructions: instructions,
		RecipeID:     recipeID,
	}

	return &recipe, nil
}

func (d *Database) FindFavoriteRecipes(currUser User, routingRecipes []Recommendation) []Recommendation{
	// open the database file
	database := d.OpenDatabase()

	var count int
	for i := range routingRecipes {
		
		// Check if the recipe is a favorite for the user
		err := database.QueryRow("SELECT COUNT(*) FROM UserFavoriteRecipes WHERE recipeID = ? AND username = ?", routingRecipes[i].R.RecipeID, currUser.UserName).Scan(&count)
		if err != nil {
			count = 0
		}
		// Set the UserFavorite field based on the query result
		if count > 0 {
			routingRecipes[i].R.UserFavorite = true
		}

	}
	return routingRecipes
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

	// close db
	database.Close()

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

	// close db
	database.Close()

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

	// close db
	database.Close()

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

	// close db
	database.Close()

	return returnUser
}

func (d *Database) ReadList(currUser User) List{
	// FUNC OVERVIEW: returns a user's list based on User (User.Username)
	database := d.OpenDatabase()

	list := List{
		ListOwner:    currUser,
		TimeUpdated:  time.Now(),
		ShoppingList: make([]FoodItem, 0),
	}

	rows, err := database.Query("SELECT FoodItemName, Quantity FROM UserLists WHERE Username = ?", currUser.UserName)
	if err != nil {
		return list
	}

	for rows.Next() {
		var foodItemName string
		var quantity int
		rows.Scan(&foodItemName, &quantity)

		foodItem := FoodItem{
			Name:     foodItemName,
			Quantity: quantity,
		}
		list.ShoppingList = append(list.ShoppingList, foodItem)
	}

	return list
}

func (d *Database) WriteList(newItem FoodItem, currUser User){
	// FUNC OVERVIEW: adds a new item to the user's list
	database := d.OpenDatabase()

	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS UserLists (
			Username TEXT,
			FoodItemName TEXT,
			Quantity INT,
			TimeUpdated TIMESTAMP,
			PRIMARY KEY (Username, FoodItemName)
		)
	`)
	if err != nil {
		// Handle error
	}

	statement, _ := database.Prepare("INSERT INTO UserLists (Username, FoodItemName, Quantity, TimeUpdated) VALUES (?, ?, ?, ?)")
	statement.Exec(currUser.UserName, newItem.Name, newItem.Quantity, time.Now())
}

func (d *Database) UpdateListItem(newItem FoodItem, currUser User){
	// FUNC OVERVIEW: updates the quantity of an item in the list, if it == 0, delete it from user's list
	database := d.OpenDatabase()

	if newItem.Quantity == 0 {
		statement, _ := database.Prepare("DELETE FROM UserLists WHERE Username = ? AND FoodItemName = ?")
		statement.Exec(currUser.UserName, newItem.Name)
	} else {
		statement, _ := database.Prepare("UPDATE UserLists SET Quantity = ?, TimeUpdated = ? WHERE Username = ? AND FoodItemName = ?")
		statement.Exec(newItem.Quantity, time.Now(), currUser.UserName, newItem.Name)
	}
}

////////////////////////////////////////////////////////////// AZURE FUNCTIONS //////////////////////////////////////////////////////////////

// OpenDatabase initializes the database connection
func AzureOpenDatabase() error {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Println("Failed to open database connection:", err)
		return err
	}
	return nil
}
// CloseDatabase closes the database connection
func AzureSQLCloseDatabase() {
	err := db.Close()
	if err != nil {
		log.Println("Failed to close database connection:", err)
	}
}
// StoreUserDatabase stores user data in the UserData table
func StoreUserDatabase(u User) error {
	ctx := context.Background()
	var err error

	if db == nil {
		err = errors.New("StoreUserDatabase: db is null")
		return err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	tsql := `
		INSERT INTO DevSchema.UserData (FirstName, LastName, Email, UserName, Password)
		VALUES (@FirstName, @LastName, @Email, @UserName, @Password);
	`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("FirstName", u.FirstName),
		sql.Named("LastName", u.LastName),
		sql.Named("Email", u.Email),
		sql.Named("UserName", u.UserName),
		sql.Named("Password", u.Password),
	)

	if err != nil {
		return err
	}

	return nil
}
