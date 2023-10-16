package main

import (
	"github.com/microsoft/go-mssqldb/azuread"
	"database/sql"
	"context"
    "log"
    "fmt"
	"time"
	"encoding/json"
    _"errors"

)

var server = "mealdealz.database.windows.net"
var port = 1433
var user = "mealdealz-dev"
var password = "Babayaga720"
var database = "MealDealz-db"

type Storage interface {
	GetPantry() (Pantry, error)
	GetPantryByUser(string) (Pantry, error)
	GetRecipes() ([]Recipe, error)
	PostSignup(*Account) error
	GetPasswordByUserName(string) (string, error)
	CheckPassword(string, string) bool
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
	
			pantry.Ingredients = append(pantry.Ingredients, Ingredient{
				Name:        foodName,
				FoodType: foodType,
				SaleDetails: "",
				Quantity: quantity,
			})
	
		}
	
		return pantry, nil
}

func (s *AzureDatabase) GetPantryByUser(username string) (Pantry, error) {	
	// Create the pantry object
	pantry := Pantry{
		Ingredients:    []Ingredient{},
	}

	tsql := fmt.Sprintf(`
	SELECT UserName, FoodName, FoodType, Quantity FROM dbo.user_ingredients
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username),
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

		pantry.Ingredients = append(pantry.Ingredients, Ingredient{
			Name:        foodName,
			FoodType: foodType,
			SaleDetails: "",
			Quantity: quantity,
		})

	}

	return pantry, nil
}

func (s *AzureDatabase) PostSignup(user *Account) error{
	ctx := context.Background()

	tsql := `
		INSERT INTO dbo.user_data (FirstName, LastName, Email, UserName, Password, DateJoined)
		VALUES (@FirstName, @LastName, @Email, @UserName, @Password, @DateJoined);
	`

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("FirstName", user.FirstName),
		sql.Named("LastName", user.LastName),
		sql.Named("Email", user.Email),
		sql.Named("UserName", user.UserName),
		sql.Named("Password", user.Password),
		sql.Named("DateJoined", time.Now()),
	)

	return nil
}

func (s *AzureDatabase) GetPasswordByUserName(userName string) (string, error){
	var password string

	tsql := fmt.Sprintf(`
	SELECT Password FROM dbo.user_data
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
    // Execute query
    rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", userName))
	
    if err != nil {
		fmt.Println("error on user password query")
        return "", err
    }
	for rows.Next() {
		err = rows.Scan(&password)
	}

	return password, nil
}

func (s *AzureDatabase) CheckPassword(username, password string) bool {
	dbPassword, _ := s.GetPasswordByUserName(username)
	if(password == dbPassword){
		return true
	}
	return false
}

func (s *AzureDatabase) GetRecipes() ([]Recipe, error){

	recipes := []Recipe{
	}

	tsql := fmt.Sprintf(`
	SELECT Title, Ingredients, Instructions, UserName from dbo.recipes;
	`)


	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
	)

	//Create Recipe
	var title, ingredientsStr, instructions, userName string
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&title, &ingredientsStr, &instructions, &userName)
		if err != nil {
			return []Recipe{}, err
		}

		var ingredients []string
		err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
		if err != nil {
			return []Recipe{}, err
		}

		recipe := Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeID:     "1",
			RecipeAuthor: userName,
		}

		recipes = append(recipes, recipe)
	}

	// return recipes
	return recipes, nil
}
