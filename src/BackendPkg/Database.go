package BackendPkg

import (
	"database/sql"
	"time"
	"encoding/json"
	// "strings"
	_ "github.com/microsoft/go-mssqldb"
	"log"
	// "strconv"
	"fmt"
	"context"
	_ "errors"
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

////////////////////////////////////////////////////////////// AZURE FUNCTIONS //////////////////////////////////////////////////////////////

// OpenDatabase initializes the database connection
func AzureOpenDatabase() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Printf("Failed to open database connection: %v", err)
		return nil, err
	}

	// Set up connection pooling and other database configurations here

	return db, nil
}

// CloseDatabase closes the database connection
func AzureSQLCloseDatabase() {
	err := db.Close()
	if err != nil {
		log.Println("Failed to close database connection:", err)
	}
}

func StoreUserDatabase(u User) error {
	var err error
	db, err := AzureOpenDatabase()

	ctx := context.Background()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	tsql := `
		INSERT INTO dbo.user_data (FirstName, LastName, Email, UserName, Password)
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

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) StoreUserPantry(u User) error {
	var err error
	db, err := AzureOpenDatabase()

	ctx := context.Background()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	tsql := `
		INSERT INTO dbo.user_ingredients (UserName, FoodName, FoodType, Quantity)
		VALUES (@UserName, @FoodName, @FoodType, @Quantity);
	`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range u.UserPantry.FoodInPantry {
		_, err = stmt.ExecContext(ctx,
			sql.Named("UserName", u.UserName),
			sql.Named("FoodName", item.Name),
			sql.Named("FoodType", item.FoodType),
			sql.Named("Quantity", item.Quantity),
		)
	}

	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) InsertPantryItemPost(currUser User, f FoodItem) error{
	var err error
	db, err := AzureOpenDatabase()

	ctx := context.Background()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	tsql := `
		INSERT INTO dbo.user_ingredients (UserName, FoodName, FoodType, Quantity)
		VALUES (@UserName, @FoodName, @FoodType, @Quantity);
	`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", currUser.UserName),
		sql.Named("Name", f.Name),
		sql.Named("FoodType", f.FoodType),
		sql.Named("Quantity", f.Quantity),
	)

	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) StoreCookie(username string, cookie string) error {
	var err error
	db, err := AzureOpenDatabase()

	ctx := context.Background()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	tsql := `
		INSERT INTO dbo.user_cookies (UserName, Cookie)
		VALUES (@UserName, @Cookie);
	`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", username),
		sql.Named("Cookie", cookie),
	)

	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) ReadPublixDatabase() ([]FoodItem, error) {
	var err error
	db, err := AzureOpenDatabase()

	if db == nil {
		fmt.Println("Failed to open database")
		return nil, err
	}

	var items []FoodItem

	tsql := `
	SELECT foodName, saleDetails FROM dbo.deals_data 
	WHERE Store = @Store;
	`

	ctx := context.Background()

	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("Store", "Publix"),
	)

	if err != nil {
        fmt.Println("error on read publix query")
        return nil, err
    }

	for rows.Next() {
		var item FoodItem
		err := rows.Scan(&item.Name, &item.SaleDetails)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (d *Database) ReadWalmartDatabase() ([]FoodItem, error) {
	var err error
	db, err := AzureOpenDatabase()

	if db == nil {
		fmt.Println("Failed to open database")
		return nil, err
	}

	var items []FoodItem

	tsql := `
	SELECT foodName, saleDetails FROM dbo.deals_data 
	WHERE Store = @Store;
	`

	ctx := context.Background()

	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("Store", "Walmart"),
	)

	if err != nil {
        fmt.Println("error on read walmart query")
        return nil, err
    }

	for rows.Next() {
		var item FoodItem
		err := rows.Scan(&item.Name, &item.SaleDetails)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (d *Database) ReadUserDatabase(username string) (User, error) {
	var err error
	db, err := AzureOpenDatabase()

	if db == nil {
		fmt.Println("Failed to open database")
		return User{}, err
	}

	var returnUser User

	tsql := fmt.Sprintf(`
	SELECT FirstName, LastName, Email, UserName, Password FROM dbo.user_data 
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username,
	))

	if err != nil {
        fmt.Println("error on user password query")
        return User{}, err
    }

	for rows.Next() {
        err = rows.Scan(&returnUser.FirstName, &returnUser.LastName, &returnUser.Email, &returnUser.UserName, &returnUser.Password)
    }

	return returnUser, nil
}

func (d *Database) UpdatePantry(currUser User, f []FoodItem) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }
	
	// Clear all of the user's current pantry
	queryDelete := fmt.Sprintf(`
	DELETE FROM dbo.user_ingredients WHERE UserName = @UserName
	`)

	stmt, err := db.Prepare(queryDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		sql.Named("UserName", currUser.UserName,
	))

	if err != nil {
        return err
    }

	// Insert all items in the received pantry to the user's pantry
	queryInsert := fmt.Sprintf(`
		INSERT INTO dbo.user_ingredients (UserName, FoodName, Foodtype, Quantity)
		VALUES (@UserName, @FoodName, @FoodType, @Quantity,)
		`)

	stmt, err = db.Prepare(queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range f {
		_, err = stmt.ExecContext(
			ctx,
			sql.Named("UserName", currUser.UserName),
			sql.Named("FoodName", item.Name),
			sql.Named("FoodType", item.FoodType),
			sql.Named("Quantity", item.Quantity),
		)
		if err != nil {
			return err
		}
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) GetUserPantry(username string) (Pantry, error) {
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return Pantry{}, err
    }

	// Create the pantry object
	pantry := Pantry{
		FoodInPantry:    []FoodItem{},
	}

	tsql := fmt.Sprintf(`
	SELECT UserName, FoodName, FoodType, Quantity, FROM dbo.user_ingredients 
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username),
	)
	if err != nil {
		fmt.Println("error on user pantry query")
		return Pantry{}, err
	}

	// Loop through each row and add the food item to the pantry
	for rows.Next() {
		var name, foodType, saleDetails string
		var quantity int

		err := rows.Scan(&name, &foodType, &saleDetails, &quantity)
		if err != nil {
			return Pantry{}, err
		}

		pantry.FoodInPantry = append(pantry.FoodInPantry, FoodItem{
			Name:        name,
			FoodType: foodType,
			SaleDetails: saleDetails,
			Quantity: quantity,
		})

	}

	return pantry, nil
}

func (d *Database) ReadPublixScrapedTime() (time.Time, error) {
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return time.Time{}, err
    }

	var dealsLastScraped time.Time

	tsql := fmt.Sprintf(`
	SELECT MAX(CAST(SaleDetails AS DATETIME)) FROM dbo.deals_data 
	WHERE Store = @Store;
	`)
	
	ctx := context.Background()
    // Execute query
    rows, err := db.QueryContext(
        ctx,
        tsql,
		sql.Named("Store", "Publix"),
	)

	if err != nil {
		return time.Time{}, nil
	}
	for rows.Next() {
		err = rows.Scan(&dealsLastScraped)
	}

	return dealsLastScraped, nil
}

func (d *Database) ReadWalmartScrapedTime() (time.Time, error) {
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return time.Time{}, err
    }

	var dealsLastScraped time.Time

	tsql := fmt.Sprintf(`
	SELECT MAX(CAST(SaleDetails AS DATETIME)) FROM dbo.deals_data 
	WHERE Store = @Store;
	`)
	
	ctx := context.Background()
    // Execute query
    rows, err := db.QueryContext(
        ctx,
        tsql,
		sql.Named("Store", "Walmart"),
	)

	if err != nil {
		return time.Time{}, nil
	}
	for rows.Next() {
		err = rows.Scan(&dealsLastScraped)
	}

	return dealsLastScraped, nil
}

func (d *Database) WriteJSONRecipes() error {
	fmt.Println("Writing Recipes")
	// Read the recipes from the file
	recipes, err := GetJSONRecipes()
	if err != nil {
		fmt.Println("Jason Failed")
		return err
	}

	db, err := AzureOpenDatabase()
	
	ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

	fmt.Println("Inserting")
	tsql := (`
	INSERT INTO dbo.jason_recipes (Title, Ingredients, Instructions)
	VALUES (@Title, @Ingredients, @Instructions);
	`)

	stmt, err := db.Prepare(tsql)
    if err != nil {
        return err
    }
    defer stmt.Close()

	// Insert each recipe into the table
	for _, recipe := range recipes {
		ingredientsJSON, _ := json.Marshal(recipe.Ingredients)
		
		_, err = stmt.ExecContext(ctx,
		sql.Named("Title", recipe.Title),
		sql.Named("Ingredients", ingredientsJSON),
		sql.Named("Instructions", recipe.Instructions),
		)

		// if err != nil {
		// 	fmt.Println("Query failure")
		// 	return err
		// }
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) ReadJSONRecipes() ([]Recipe, error) {
	// Establish a connection to the Azure SQL Database
	var recipes []Recipe
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return recipes, err
    }


	tsql := fmt.Sprintf(`
	"SELECT RecipeID, Title, Ingredients, Instructions 
	FROM dbo.jason_recipes"
	`)

	ctx := context.Background()
	
	rows, err := db.QueryContext(
		ctx,
		tsql,
	)

	if err != nil {
        fmt.Println("error on user password query")
        return recipes, err
    }

	// Iterate through the rows and create a slice of Recipe structs
	for rows.Next() {
		var title, ingredientsStr, instructions string
		var recipeID string //this needs to be an int
		err := rows.Scan(&recipeID, &title, &ingredientsStr, &instructions)
		if err != nil {
			return nil, err
		}

		// Convert the JSON string of ingredients to a slice
		var ingredients []string
		err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
		if err != nil {
			return nil, err
		}

		// Create a new Recipe struct and append it to the slice
		recipe := Recipe{
			RecipeID:     recipeID,
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeAuthor: "", // set to null for all recipes from the JSON file
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (d *Database) WriteNewUserRecipe(currUser User, newRecipe Recipe) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

    tsql := fmt.Sprintf(`
        INSERT INTO dbo.user_recipes (Title, Ingredients, Instructions, UserName)
        VALUES (@Title, @Ingredients, @Instructions, @UserName);
    `)

	stmt, err := db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	// Insert the new recipe into the table
	ingredientsJSON, _ := json.Marshal(newRecipe.Ingredients)
	_, err = stmt.ExecContext(
		ctx,
		sql.Named("Title", newRecipe.Title),
		sql.Named("Ingreidents", string(ingredientsJSON)),
		sql.Named("Instructions", newRecipe.Instructions),
		sql.Named("UserName", currUser.UserName),
	)
	if err != nil {
		return err
	}

	AzureSQLCloseDatabase()
	return nil
}

func (d *Database) DeleteUserRecipe(recipeID string) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

    tsql := fmt.Sprintf(`
	Delete from dbo.user_recipes
	WHERE RecipeID = @RecipeID
	`)

	stmt, err := db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
        sql.Named("RecipeID", recipeID),
    )

	if err != nil {
        return err
    }

    AzureSQLCloseDatabase();
	return nil
}

func (d *Database) ReadAllUserRecipes() ([]Recipe, error) {
	// Establish a connection to the Azure SQL Database
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return []Recipe{}, err
    }

    tsql := fmt.Sprintf(`
    SELECT Title Ingredients, Instructions, RecipeID, UserName FROM dbo.recipes;
    `)

	ctx := context.Background()
	rows, err := db.QueryContext(
		ctx,
		tsql,
	)

	if err != nil {
		fmt.Println("error on user recipe read")
		return []Recipe{}, err
	}

	var recipes []Recipe

	// Iterate through the rows and create a slice of Recipe structs
	for rows.Next() {
		var title, ingredientsStr, instructions, recipeID, userName string
		err := rows.Scan(&title, &ingredientsStr, &instructions, &recipeID, &userName)
		if err != nil {
			return []Recipe{}, err
		}

		// Convert the JSON string of ingredients to a slice
		var ingredients []string
		err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
		if err != nil {
			return []Recipe{}, err
		}

		// Create a new Recipe struct and append it to the slice
		recipe := Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeID:     recipeID,
			RecipeAuthor: userName,
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (d *Database) ReadCurrUserRecipes(currUser User) ([]Recipe, error) {
	// Establish a connection to the Azure SQL Database
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return []Recipe{}, err
    }

    tsql := fmt.Sprintf(`
    SELECT Title Ingredients, Instructions, RecipeID, UserName FROM dbo.recipes
	WHERE UserName = @UserName;
    `)

	ctx := context.Background()
	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", currUser.UserName),
	)

	if err != nil {
		fmt.Println("error on currUser recipe read")
		return []Recipe{}, err
	}

	var recipes []Recipe

	// Iterate through the rows and create a slice of Recipe structs
	for rows.Next() {
		var title, ingredientsStr, instructions, recipeID, userName string
		err := rows.Scan(&title, &ingredientsStr, &instructions, &recipeID, &userName)
		if err != nil {
			return []Recipe{}, err
		}

		// Convert the JSON string of ingredients to a slice
		var ingredients []string
		err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
		if err != nil {
			return []Recipe{}, err
		}

		// Create a new Recipe struct and append it to the slice
		recipe := Recipe{
			Title:        title,
			Ingredients:  ingredients,
			Instructions: instructions,
			RecipeID:     recipeID,
			RecipeAuthor: userName,
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (d *Database) FavoriteRecipe(currUser User, recipeID string) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

    tsql := fmt.Sprintf(`
        INSERT INTO dbo.user_favorite_recipes (RecipeID, UserName)
        VALUES (@RecipeID, UserName);
    `)

	stmt, err := db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		tsql, 
		sql.Named("RecipeID", recipeID),
		sql.Named("UserName", currUser.UserName),
	)
	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) UnfavoriteRecipe(currUser User, recipeID string) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

    tsql := fmt.Sprintf(`
    	DELETE FROM dbo.user_favorite_recipes
       	WHERE UserName = @UserName AND RecipeID = @RecipeID;
    `)

	stmt, err := db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		tsql, 
		sql.Named("RecipeID", recipeID),
		sql.Named("UserName", currUser.UserName),
	)
	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) ReadFavoriteRecipes(currUser User) ([]Recipe, error) {
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return []Recipe{}, err
    }

	var recipes []Recipe

	tsql := fmt.Sprintf(`
	SELECT RecipeID FROM dbo.user_favorite_recipes
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
    // Execute query
    rows, err := db.QueryContext(
        ctx,
        tsql,
        sql.Named("UserName", currUser.UserName))
    
    if err != nil {
        fmt.Println("error on user password query")
        return []Recipe{}, err
    }

	var recipeIDs []string

	for rows.Next() {
		var recipeID string
		err := rows.Scan(&recipeID)
		if err != nil {
			return []Recipe{}, err
		}
		recipeIDs = append(recipeIDs, recipeID)
	}

	for _, recipeID := range recipeIDs {
		// Retrieve the recipe based on recipeID
		recipe, err := d.getRecipeByID(recipeID)
		if err != nil {
			return []Recipe{}, err
		}
		if recipe != nil {
			recipes = append(recipes, *recipe)
		}
	}

	return recipes, nil
}

// Helper function to get a recipe by RecipeID
func (d *Database) getRecipeByID(recipeID string) (*Recipe, error) {
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return &Recipe{}, err
    }

	tsql := fmt.Sprintf(`
	SELECT Title, Ingreidents, Instructions from dbo.user_recipes
	WHERE RecipeID = @RecipeID;
	`)
	ctx := context.Background()
	row, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("RecipeID", recipeID),
	)

	//Create Recipe
	var title, ingredientsStr, instructions string
	err = row.Scan(&title, &ingredientsStr, &instructions)

	// Convert the JSON string of ingredients to a slice
	var ingredients []string
	err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
	if err != nil {
		return nil, err
	}

	recipe := Recipe{
		Title:        title,
		Ingredients:  ingredients,
		Instructions: instructions,
		RecipeID:     recipeID,
	}

	return &recipe, nil
}

func (d *Database) getRecipeFromUserTable(recipeID string) (*Recipe, error) {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT Title, Ingredients, Instructions FROM dbo.user_recipes WHERE RecipeID = ?"
	row := db.QueryRow(query, recipeID)

	var title, ingredientsStr, instructions string
	err = row.Scan(&title, &ingredientsStr, &instructions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Recipe not found in user_recipes
		}
		return nil, err
	}

	// Convert the JSON string of ingredients to a slice
	var ingredients []string
	err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
	if err != nil {
		return nil, err
	}

	// Define a Recipe object based on returned values
	recipe := Recipe{
		Title:        title,
		Ingredients:  ingredients,
		Instructions: instructions,
		RecipeID:     recipeID,
	}

	return &recipe, nil
}

func (d *Database) getRecipeFromJSONTable(recipeID string) (*Recipe, error) {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT Title, Ingredients, Instructions FROM dbo.jason_recipes WHERE recipeID = ?"
	row := db.QueryRow(query, recipeID)

	var title, ingredientsStr, instructions string
	err = row.Scan(&title, &ingredientsStr, &instructions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Recipe not found in JSONRecipeData
		}
		return nil, err
	}

	// Convert the JSON string of ingredients to a slice
	var ingredients []string
	err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
	if err != nil {
		return nil, err
	}

	// Define a Recipe object based on returned values
	recipe := Recipe{
		Title:        title,
		Ingredients:  ingredients,
		Instructions: instructions,
		RecipeID:     recipeID,
	}

	return &recipe, nil
}

func (d *Database) FindFavoriteRecipes(currUser User, routingRecipes []Recommendation) []Recommendation {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return routingRecipes // Return the original list if there's a database error
	}
	defer db.Close()

	// Iterate through the routingRecipes and check if each recipe is a favorite for the user
	for i := range routingRecipes {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM dbo.user_favorite_recipes WHERE RecipeID = ? AND UserName = ?", routingRecipes[i].R.RecipeID, currUser.UserName).Scan(&count)
		if err != nil {
			count = 0
		}
		// Set the UserFavorite field based on the query result
		routingRecipes[i].R.UserFavorite = count > 0
	}

	return routingRecipes
}

func (d *Database) GetUserPassword(username string) (string, error) {
	// Establish a connection to the Azure SQL Database
	var err error
	db, err := AzureOpenDatabase()

	if db == nil {
		fmt.Println("Failed to open database")
		return "", err
	}

	var password string

	tsql := fmt.Sprintf(`
	SELECT Password FROM dbo.user_data
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
    // Execute query
    rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username))
	
    if err != nil {
		fmt.Println("error on user password query")
        return "", err
    }
	for rows.Next() {
		err = rows.Scan(&password)
	}

	return password, nil
}

func (d *Database) ReadCookie(username string) (string, error) {
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return "", err
    }

    var cookie string

	tsql := fmt.Sprintf(`
	SELECT Cookie from dbo.user_cookies
	WHERE UserName = @UserName;
	`)
	
	ctx := context.Background()

	rows, err := db.QueryContext(
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
	}

	return cookie, nil
}

func (d *Database) UserFromCookie(cookie string) (User, error) {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	var returnUser User
	var userName string

	stmt, err := db.Prepare("SELECT UserName FROM dbo.user_cookies WHERE Cookie=?")
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(cookie)
	err = row.Scan(&userName)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case where the cookie is not found
			return User{}, nil
		}
		return User{}, err
	}

	// Retrieve user details based on the username
	returnUser, err = d.ReadUserDatabase(userName)

	return returnUser, nil
}

func (d *Database) ReadList(currUser User) List {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return List{
			ShoppingList: make([]FoodItem, 0),
		}
	}
	defer db.Close()

	list := List{
		ShoppingList: make([]FoodItem, 0),
	}

	rows, err := db.Query("SELECT FoodName, Quantity FROM dbo.user_lists WHERE UserName = ?", currUser.UserName)
	if err != nil {
		return list
	}

	for rows.Next() {
		var foodName string
		var quantity int
		err := rows.Scan(&foodName, &quantity)
		if err != nil {
			continue
		}

		foodItem := FoodItem{
			Name:     foodName,
			Quantity: quantity,
		}
		list.ShoppingList = append(list.ShoppingList, foodItem)
	}

	return list
}

func (d *Database) WriteList(newItem FoodItem, currUser User) {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		// Handle error
		return
	}
	defer db.Close()

	// Insert or update the item in the "user_lists" table
	statement, err := db.Prepare("MERGE INTO user_lists AS target "+
		"USING (SELECT ? AS UserName, ? AS FoodName, ? AS FoodType, ? AS Quantity) AS source "+
		"ON (target.UserName = source.UserName AND target.FoodName = source.FoodName) "+
		"WHEN MATCHED THEN "+
		"UPDATE SET target.Quantity = source.Quantity "+
		"WHEN NOT MATCHED THEN "+
		"INSERT (UserName, FoodName, FoodType, Quantity) "+
		"VALUES (source.UserName, source.FoodName, source.FoodType, source.Quantity);")

	if err != nil {
		// Handle error
		return
	}

	_, err = statement.Exec(currUser.UserName, newItem.Name, newItem.FoodType, newItem.Quantity)
	if err != nil {
		// Handle error
		return
	}
}

func (d *Database) UpdateListItem(newItem FoodItem, currUser User) {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		// Handle error
		return
	}
	defer db.Close()

	// Update the item in the "user_lists" table
	statement, err := db.Prepare("UPDATE dbo.user_lists SET Quantity = ? WHERE UserName = ? AND FoodName = ?")
	if err != nil {
		// Handle error
		return
	}

	_, err = statement.Exec(newItem.Quantity, currUser.UserName, newItem.Name)
	if err != nil {
		// Handle error
		return
	}
}

