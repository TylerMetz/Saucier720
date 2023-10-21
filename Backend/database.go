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
	// Signup / Login
	PostSignup(*Account) error
	GetPasswordByUserName(string) (string, error)
	// Pantry
	GetPantry() (Pantry, error)
	GetPantryByUserName(string) (Pantry, error)
	PostPantryIngredient(string, Ingredient) error
	UpdatePantryByUserName(string, Pantry) error
	DeletePantryIngredient(string, Ingredient) error
	// Recipes
	GetRecipes() ([]Recipe, error)
	GetUserCreatedRecipes() ([]Recipe, error)
	GetRecipesByUserName(string) ([]Recipe, error)
	GetRecipesByRecipeID(int) (Recipe, error)
	PostRecipe(string, Recipe) error
	UpdateRecipeByUserName(string, Recipe) error
	DeleteRecipe(string, int) error
	// Favorite Recipes
	GetFavoriteRecipes(string) ([]Recipe, error)
	PostFavoriteRecipe(string, int) error
	DeleteFavorite(string, int) error
	// Deals
	GetDeals() ([]Ingredient, error)
	GetDealsByStore(string) ([]Ingredient, error)
	//List
	GetShoppingListByUserName(string) ([]Ingredient, error)
	PostListIngredient(string, Ingredient) error
	UpdateListByUserName(string, Pantry) error
	DeleteListIngredient(string, Ingredient) error
	// Cookies
	GetCookieByUserName(string) (string, error)
	PostCookieByUserName(string, string) error
	UpdateCookieByUserName(string, string) error
	DeleteCookieByUserName(string) error
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

// GETS

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

func (s *AzureDatabase) GetPantryByUserName(username string) (Pantry, error) {	
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

func (s *AzureDatabase) GetRecipes() ([]Recipe, error){

	recipes := []Recipe{
	}

	tsql := fmt.Sprintf(`
	SELECT RecipeID, Title, Ingredients, Instructions, UserName from dbo.recipes;
	`)


	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
	)
	if err != nil {
		fmt.Println("error on recipe query")
		return []Recipe{}, err
	}

	//Create Recipe
	var title, ingredientsStr, instructions, userName string
	var recipeID int
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&recipeID, &title, &ingredientsStr, &instructions, &userName)
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
			RecipeID:     recipeID,
			RecipeAuthor: userName,
		}

		recipes = append(recipes, recipe)
	}

	// return recipes
	return recipes, nil
}

func (s *AzureDatabase) GetUserCreatedRecipes() ([]Recipe, error){
	recipes := []Recipe{
	}

	tsql := fmt.Sprintf(`
	SELECT RecipeID, Title, Ingredients, Instructions, UserName from dbo.recipes
	WHERE UserName != @UserName;
	`)


	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", "MealDealz Classic Recipes"),
	)

	//Create Recipe
	var title, ingredientsStr, instructions, userName string
	var recipeID int
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&recipeID, &title, &ingredientsStr, &instructions, &userName)
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
			RecipeID:     recipeID,
			RecipeAuthor: userName,
		}

		recipes = append(recipes, recipe)
	}

	// return recipes
	return recipes, nil
}

func (s *AzureDatabase) GetRecipesByUserName(username string) ([]Recipe, error) {
	recipes := []Recipe{
	}

	tsql := fmt.Sprintf(`
	SELECT RecipeID, Title, Ingredients, Instructions, UserName from dbo.recipes
	WHERE UserName = @UserName;
	`)


	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username),
	)

	//Create Recipe
	var title, ingredientsStr, instructions, userName string
	var recipeID int
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&recipeID, &title, &ingredientsStr, &instructions, &userName)
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
			RecipeID:     recipeID,
			RecipeAuthor: userName,
		}

		recipes = append(recipes, recipe)
	}

	// return recipes
	return recipes, nil
}

func (s *AzureDatabase) GetRecipesByRecipeID(id int) (Recipe, error) {
	recipe := Recipe{}

	tsql := fmt.Sprintf(`
	SELECT Title, Ingredients, Instructions, UserName from dbo.recipes
	WHERE RecipeID = @RecipeID;
	`)

	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("RecipeID", id),
	)

	//Create Recipe
	var title, ingredientsStr, instructions, userName string
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&title, &ingredientsStr, &instructions, &userName)
		if err != nil {
			return Recipe{}, err
		}

		var ingredients []string
		err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
		if err != nil {
			return Recipe{}, err
		}

		recipe = Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeID:     id,
			RecipeAuthor: userName,
		}
	}

	// return recipes
	return recipe, nil
}

func (s *AzureDatabase) GetFavoriteRecipes(username string) ([]Recipe, error) {
	recipes := []Recipe{
	}

	tsql := fmt.Sprintf(`
	SELECT RecipeID, Title, Ingredients, Instructions, UserName from dbo.recipes
	WHERE RecipeID IN (
		SELECT RecipeID FROM dbo.user_favorite_recipes
		WHERE UserName = @UserName
	);
	`)

	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username),
	)
	if err != nil {
		fmt.Println("error on favorite recipe query")
		return []Recipe{}, err
	}

	//handleLogin

	var title, ingredientsStr, instructions, userName string
	var recipeID int
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&recipeID, &title, &ingredientsStr, &instructions, &userName)
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
			RecipeID:     recipeID,
			RecipeAuthor: userName,
		}

		recipes = append(recipes, recipe)
	}

	// return recipes
	return recipes, nil
}

func (s *AzureDatabase) GetDeals() ([]Ingredient, error) {
	deals := []Ingredient{}

	tsql := fmt.Sprintf(`
	SELECT Store, FoodName, SaleDetails from dbo.deals_data;
	`)

	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
	)
	if err != nil {
		fmt.Println("error on deals query")
		return []Ingredient{}, err
	}

	var store, foodName, saleDetails string
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&store, &foodName, &saleDetails)
		if err != nil {
			return []Ingredient{}, err
		}

		ingredient := Ingredient{
			Name: 	  foodName,
			SaleDetails: saleDetails,
			FoodType: "Food", // will need to be updated when food typing introduced
			Quantity: 1,
		}

		deals = append(deals, ingredient)
	}

	return []Ingredient{}, nil
}

func (s *AzureDatabase) GetDealsByStore(storeName string) ([]Ingredient, error) {
	deals := []Ingredient{}

	tsql := fmt.Sprintf(`
	SELECT Store, FoodName, SaleDetails from dbo.deals_data
	WHERE Store = @Store;
	`)

	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("Store", storeName),
	)
	if err != nil {
		fmt.Println("error on deals by store query")
		return []Ingredient{}, err
	}

	var store, foodName, saleDetails string
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&store, &foodName, &saleDetails)
		if err != nil {
			return []Ingredient{}, err
		}

		ingredient := Ingredient{
			Name: 	  foodName,
			SaleDetails: saleDetails,
			FoodType: "Food", // will need to be updated when food typing introduced
			Quantity: 1,
		}

		deals = append(deals, ingredient)
	}

	return []Ingredient{}, nil
}

func (s *AzureDatabase) GetShoppingListByUserName(username string) ([]Ingredient, error) {
	shoppingList := []Ingredient{}

	tsql := fmt.Sprintf(`
	SELECT UserName, FoodName, FoodType, Quantity FROM dbo.list
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username),
	)
	if err != nil {
		fmt.Println("error on shopping list query")
		return []Ingredient{}, err
	}

	var name, foodName, foodType string
	var quantity int
	for rows.Next() {
		//append to recipe to get all
		err = rows.Scan(&name, &foodName, &foodType, &quantity)
		if err != nil {
			return []Ingredient{}, err
		}

		ingredient := Ingredient{
			Name: 	  foodName,
			FoodType: foodType,
			Quantity: quantity,
			SaleDetails: "",
		}

		shoppingList = append(shoppingList, ingredient)
	}

	return shoppingList, nil
}

func (s *AzureDatabase) GetCookieByUserName(username string) (string, error) {
	var cookie string

	tsql := fmt.Sprintf(`
	SELECT Cookie from dbo.deals_data
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	rows, err := s.db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username),
	)
	if err != nil {
		fmt.Println("error on cookie query")
		return "", err
	}

	for rows.Next() {
		err = rows.Scan(&cookie)
		if err != nil {
			return "", err
		}
	}

	return cookie, nil
}

// POSTS

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

func (s *AzureDatabase) PostPantryIngredient(username string, newPantryItem Ingredient) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
		INSERT INTO dbo.user_ingredients (UserName, FoodName, FoodType, Quantity)
		VALUES (@UserName, @FoodName, @FoodType, @Quantity);
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", username),
		sql.Named("FoodName", newPantryItem.Name),
		sql.Named("FoodType", newPantryItem.FoodType),
		sql.Named("Quantity", newPantryItem.Quantity),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AzureDatabase) PostRecipe(username string, recipe Recipe) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
	INSERT INTO dbo.recipes (Title, Ingredients, Instructions, UserName)
	VALUES (@Title, @Ingredients, @Instructions, @UserName);
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("Title", recipe.Title),
		sql.Named("Ingredients", recipe.Ingredients),
		sql.Named("Instructions", recipe.Ingredients),
		sql.Named("UserName", username),
	)
	if err != nil {
		fmt.Println("error on recipe post")
		return err
	}

	return nil
}

func (s *AzureDatabase) PostListIngredient(username string, ingredient Ingredient) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
	INSERT INTO dbo.user_lists (UserName, FoodName, FoodType, Quantity)
	VALUES (@UserName, @FoodName, @FoodType, @Quantity);
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", username),
		sql.Named("FoodName", ingredient.Name),
		sql.Named("FoodType", ingredient.FoodType),
		sql.Named("Quantity", ingredient.Quantity),
	)
	if err != nil {
		fmt.Println("error on list post")
		return err
	}

	return nil
}

func (s *AzureDatabase) PostFavoriteRecipe(username string, id int) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
	INSERT INTO dbo.user_favorite_recipes (UserName, RecipeID)
	VALUES (@UserName, @RecipeID);
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", username),
		sql.Named("RecipeID", id),
	)
	if err != nil {
		fmt.Println("error on favorite recipe post")
		return err
	}

	return nil
}

func (s *AzureDatabase) PostCookieByUserName(username string, cookie string) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
	INSERT INTO dbo.user_cookies (UserName, Cookie)
	VALUES (@UserName, @Cookie);
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	fmt.Println("username: ", username)
	fmt.Println("cookie: ", cookie)

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", username),
		sql.Named("Cookie", cookie),
	)
	if err != nil {
		fmt.Println("error on cookie post")
		updateCookieerr := s.UpdateCookieByUserName(username, cookie)
		if updateCookieerr == nil {
			return nil
		}
		return err
	}

	return nil
}

// DELETES
func (s *AzureDatabase) DeletePantryIngredient(username string, ingredient Ingredient) error {
	ctx := context.Background()

    tsql := fmt.Sprintf(`
	DELETE from dbo.user_ingredients
	WHERE UserName = @UserName
	AND FoodName = @FoodName;
	`)

	stmt, err := s.db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", username),
        sql.Named("FoodName", ingredient.Name),
    )

	if err != nil {
        return err
    }

	return nil
}

func (s *AzureDatabase) DeleteListIngredient(username string, ingredient Ingredient) error {
	ctx := context.Background()

    tsql := fmt.Sprintf(`
	DELETE from dbo.user_lists
	WHERE UserName = @UserName
	AND FoodName = @FoodName;
	`)

	stmt, err := s.db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", username),
        sql.Named("FoodName", ingredient.Name),
    )

	if err != nil {
        return err
    }

	return nil
}

func (s *AzureDatabase)	DeleteFavorite(username string, id int) error {
	ctx := context.Background()

    tsql := fmt.Sprintf(`
    	DELETE FROM dbo.user_favorite_recipes
       	WHERE UserName = @UserName 
		AND RecipeID = @RecipeID;
    `)

	stmt, err := s.db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		sql.Named("RecipeID", id),
		sql.Named("UserName", username),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AzureDatabase)	DeleteRecipe(username string, id int) error {
	ctx := context.Background()

    tsql := fmt.Sprintf(`
    	DELETE FROM dbo.recipes
       	WHERE UserName = @UserName 
		AND RecipeID = @RecipeID;
    `)

	stmt, err := s.db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		sql.Named("RecipeID", id),
		sql.Named("UserName", username),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AzureDatabase) DeleteCookieByUserName(username string) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
		DELETE FROM dbo.user_cookies
	   	WHERE UserName = @UserName;
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		sql.Named("UserName", username),
	)
	if err != nil {
		return err
	}

	return nil
}

// UPDATES
func (s *AzureDatabase) UpdatePantryByUserName(username string, pantry Pantry) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
		UPDATE dbo.user_ingredients
		SET Quantity = @Quantity AND FoodType = @FoodType
		WHERE UserName = @UserName
		AND FoodName = @FoodName;
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	for _, ingredient := range pantry.Ingredients {
		_, err = stmt.ExecContext(ctx,
			sql.Named("UserName", username),
			sql.Named("FoodName", ingredient.Name),
			sql.Named("FoodType", ingredient.FoodType),
			sql.Named("Quantity", ingredient.Quantity),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *AzureDatabase) UpdateListByUserName(username string, pantry Pantry) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
		UPDATE dbo.user_lists
		SET Quantity = @Quantity
		WHERE UserName = @UserName
		AND FoodName = @FoodName;
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	for _, ingredient := range pantry.Ingredients {
		_, err = stmt.ExecContext(ctx,
			sql.Named("UserName", username),
			sql.Named("FoodName", ingredient.Name),
			sql.Named("Quantity", ingredient.Quantity),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *AzureDatabase) UpdateRecipeByUserName(username string, recipe Recipe) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
		UPDATE dbo.recipes
		SET Title = @Title, Ingredients = @Ingredients, Instructions = @Instructions
		WHERE UserName = @UserName
		AND RecipeID = @RecipeID;
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("Title", recipe.Title),
		sql.Named("Ingredients", recipe.Ingredients),
		sql.Named("Instructions", recipe.Instructions),
		sql.Named("UserName", username),
		sql.Named("RecipeID", recipe.RecipeID),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AzureDatabase) UpdateCookieByUserName(username string, cookie string) error {
	ctx := context.Background()

	tsql := fmt.Sprintf(`
		UPDATE dbo.user_cookies
		SET Cookie = @Cookie
		WHERE UserName = @UserName;
	`)

	stmt, err := s.db.Prepare(tsql)
	if err != nil {
		
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("Cookie", cookie),
		sql.Named("UserName", username),
	)
	if err != nil {
		return err
	}

	return nil
}
