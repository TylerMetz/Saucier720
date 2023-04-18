package main

import (
	"database/sql"
	"testing"
	"BackendPkg"
)

func TestUpdatePantry(t *testing.T) {
	// Initialize test variables
	d := BackendPkg.Database{} // assuming Database struct has been defined
	currUser := BackendPkg.User{UserName: "testUser"}
	f := []BackendPkg.FoodItem{
		{Name: "apple", StoreCost: 1.50, Quantity: 2},
		{Name: "banana", StoreCost: 0.75, Quantity: 3},
	}

	// Create a temporary in-memory database for testing
	db, err := sql.Open("sqlite3", "MealDealz.db")
	if err != nil {
		t.Fatalf("Failed to create temporary database: %v", err)
	}
	defer db.Close()

	// Create UserPantries table for testing
	createTableStmt := `
		CREATE TABLE IF NOT EXISTS UserPantries (
			UserName TEXT,
			PantryLastUpdated DATETIME,
			Name TEXT,
			StoreCost REAL,
			OnSale BOOL,
			SalePrice REAL,
			SaleDetails TEXT,
			Quantity INT,
			PRIMARY KEY (UserName, Name)
		);
	`
	_, err = db.Exec(createTableStmt)
	if err != nil {
		t.Fatalf("Failed to create UserPantries table: %v", err)
	}

	// Call the function being tested
	d.UpdatePantry(currUser, f)

	// Verify that the pantry was updated correctly
	rows, err := db.Query("SELECT Name, StoreCost, Quantity FROM UserPantries WHERE UserName = ?", currUser.UserName)
	if err != nil {
		t.Fatalf("Failed to query UserPantries table: %v", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var name string
		var storeCost float64
		var quantity int
		err = rows.Scan(&name, &storeCost, &quantity)
		if err != nil {
			t.Fatalf("Failed to scan row from UserPantries table: %v", err)
		}

		// Verify that the item was inserted correctly
		if name == "apple" {
			if storeCost != 1.50 {
				t.Errorf("Expected apple store cost to be 1.50, but got %v", storeCost)
			}
			if quantity != 2 {
				t.Errorf("Expected apple quantity to be 2, but got %v", quantity)
			}
			count++
		} else if name == "banana" {
			if storeCost != 0.75 {
				t.Errorf("Expected banana store cost to be 0.75, but got %v", storeCost)
			}
			if quantity != 3 {
				t.Errorf("Expected banana quantity to be 3, but got %v", quantity)
			}
			count++
		} else {
			t.Errorf("Unexpected item found in UserPantries table: %v", name)
		}
	}

	// Verify that all items were inserted
	if count != len(f) {
		t.Errorf("Expected %v items to be inserted, but %v were inserted", len(f), count)
	}
}