package main

import (
	"github.com/microsoft/go-mssqldb/azuread"
	"database/sql"
	"context"
    "log"
    "fmt"
    _"errors"

)

var server = "mealdealz.database.windows.net"
var port = 1433
var user = "mealdealz-dev"
var password = "Babayaga720"
var database = "MealDealz-db"

type Storage interface {
	GetPantry() (Pantry, error)
}

type AzureDatabase struct {
	db *sql.DB
}

func NewAzureDatabase() (*AzureDatabase, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	server, user, password, port, database)

	var db *sql.DB
	var err error

	db, err = sql.Open(azuread.DriverName, connString)
	if err != nil {
        log.Fatal("Error creating connection pool: ", err.Error())
    }
    ctx := context.Background()
    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal(err.Error())
    }
    fmt.Printf("Connected!\n")

	return &AzureDatabase{
		db: db,
	}, nil
}

func (s *AzureDatabase) GetPantry() (Pantry, error) {	
		// Create the pantry object
		pantry := Pantry{
			Ingredients:    []Ingredient{},
		}
	
		tsql := fmt.Sprintf(`
		SELECT UserName, FoodName, FoodType, Quantity FROM dbo.user_ingredients
		`)
	
		ctx := context.Background()
		rows, err := s.db.QueryContext(
			ctx,
			tsql,
		)
		if err != nil {
			log.Fatal(err)
		}
	
		// Loop through each row and add the food item to the pantry
		for rows.Next() {
			var name, foodName, foodType string
			var quantity int
	
			err := rows.Scan(&name, &foodName, &foodType, &quantity)
			if err != nil {
				return Pantry{}, err
			}

			fmt.Println(foodName)
	
			pantry.Ingredients = append(pantry.Ingredients, Ingredient{
				Name:        foodName,
				FoodType: foodType,
				SaleDetails: "",
				Quantity: quantity,
			})
	
		}
	
		return pantry, nil
	}
