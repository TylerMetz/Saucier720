package BackendPkg

import (
    "database/sql"
    _"github.com/mattn/go-sqlite3"
	"time"
	//"fmt"
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

func (d *Database) StoreUserPantry (u User){

	// calls function to open the database
	database := d.OpenDatabase()

	// make table for food item data
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS UserPantries (UserName TEXT, PantryLastUpdated DATETIME, Name TEXT, StoreCost REAL, OnSale INTEGER, SalePrice REAL, SaleDetails TEXT, Quantity INTEGER, PRIMARY KEY (UserName, Name))")
	statement.Exec();

	// insert into food item table
	statementTwo, _ := database.Prepare("INSERT OR IGNORE INTO UserPantries (UserName, PantryLastUpdated, Name, StoreCost, OnSale, SalePrice, SaleDetails, Quantity) VALUES (?, datetime(?), ?, ?, ?, ?, ?, ?)")

	for _, item := range u.UserPantry.FoodInPantry {
		statementTwo.Exec(u.UserName, u.UserPantry.TimeLastUpdated.Format("2006-01-02 15:04:05"), item.Name, item.StoreCost, item.OnSale, item.SalePrice, item.SaleDetails, item.Quantity)
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
