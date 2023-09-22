package BackendPkg

import (
	"database/sql"
	// "time"
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

// SEAL OF APPROVAL
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

// SEAL OF APPROVAL
func (d *Database) InsertPantryItemPost(currUser User, f FoodItem) error{
	var err error
	db, err := AzureOpenDatabase()

	ctx := context.Background()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	tsql := fmt.Sprintf(`
		INSERT INTO dbo.user_ingredients (UserName, FoodName, FoodType, Quantity)
		VALUES (@UserName, @FoodName, @FoodType, @Quantity);
	`)

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		sql.Named("UserName", currUser.UserName),
		sql.Named("FoodName", f.Name),
		sql.Named("FoodType", f.FoodType),
		sql.Named("Quantity", f.Quantity),
	)

	if err != nil {
		return err
	}

	return nil
}

// STORE OF APPROVAL
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

// SEAL OF APPROVAL
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

// SEAL OF APPROVAL
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
		sql.Named("UserName", username),
	)

	if err != nil {
        fmt.Println("error on user password query")
        return User{}, err
    }

	for rows.Next() {
        err = rows.Scan(&returnUser.FirstName, &returnUser.LastName, &returnUser.Email, &returnUser.UserName, &returnUser.Password)
    }

	fmt.Println("passed in userfrom for db query:" + username)
	fmt.Println("returned username from db query: " + returnUser.UserName)
	return returnUser, nil
}

// SEAL OF APPROVAL
func (d *Database) UpdatePantry(currUser User, f []FoodItem) error {
	fmt.Println("updating Pantry")
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

	// Insert all items in the received pantry to the user's pantry
	updateQuery := fmt.Sprintf(`
		UPDATE dbo.user_ingredients
		SET Quantity = @Quantity
		WHERE FoodName = @FoodName
		AND UserName = @UserName;
		`)

	stmt, err := db.Prepare(updateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range f {
		fmt.Println("Updating Pantry with: " + item.Name)
		_, err = stmt.ExecContext(
			ctx,
			sql.Named("UserName", currUser.UserName),
			sql.Named("FoodName", item.Name),
			sql.Named("Quantity", item.Quantity),
		)
		if err != nil {
			fmt.Println("Error on inserting user pantry")
			return err
		}
	}
	return nil
}

// SEAL OF APPROVAL
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
	SELECT UserName, FoodName, FoodType, Quantity FROM dbo.user_ingredients 
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", username),
	)
	if err != nil {
		fmt.Println("error on user pantry query, username was: " + username)
		return Pantry{}, err
	}

	// Loop through each row and add the food item to the pantry
	for rows.Next() {
		var name, foodName, foodType string
		var quantity int

		err := rows.Scan(&name, &foodName, &foodType, &quantity)
		if err != nil {
			return Pantry{}, err
		}

		pantry.FoodInPantry = append(pantry.FoodInPantry, FoodItem{
			Name:        foodName,
			FoodType: foodType,
			SaleDetails: "",
			Quantity: quantity,
		})

	}

	return pantry, nil
}

// SEAL OF APPROVAL
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
	INSERT INTO dbo.recipes (Title, Ingredients, Instructions, UserName)
	VALUES (@Title, @Ingredients, @Instructions, @UserName);
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
		sql.Named("UserName", "MealDealz Classic Recipe"),
		)

		// if err != nil {
		// 	fmt.Println("Query failure")
		// 	return err
		// }
	}

	AzureSQLCloseDatabase();
	return nil
}

// SEAL OF APPROVAL
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
	SELECT RecipeID, Title, Ingredients, Instructions FROM dbo.recipes
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	
	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", "MealDealz Classic Recipe"),
	)

	if err != nil {
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
			RecipeAuthor: "MealDealz Classic Recipe", // set to null for all recipes from the JSON file
		}
		recipes = append(recipes, recipe)
	}

	fmt.Println(recipes)
	return recipes, nil
}

// SEAL OF APPROVAL
func (d *Database) WriteNewUserRecipe(currUser User, newRecipe Recipe) error {
	fmt.Println("inserting new user recipe")
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

    tsql := fmt.Sprintf(`
        INSERT INTO dbo.recipes (Title, Ingredients, Instructions, UserName)
        VALUES (@Title, @Ingredients, @Instructions, @UserName);
    `)

	stmt, err := db.Prepare(tsql)
    if err != nil {
        fmt.Println("error inserting user recipe")
        return err
    }
    defer stmt.Close()

	// Insert the new recipe into the table
	ingredientsJSON, _ := json.Marshal(newRecipe.Ingredients)
	_, err = stmt.ExecContext(
		ctx,
		sql.Named("Title", newRecipe.Title),
		sql.Named("Ingredients", ingredientsJSON),
		sql.Named("Instructions", newRecipe.Instructions),
		sql.Named("UserName", currUser.UserName),
	)
	if err != nil {
		return err
	}

	return nil
}

// SEAL OF APPROVAL
func (d *Database) DeleteUserRecipe(recipeID string) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

    tsql := fmt.Sprintf(`
	DELETE from dbo.recipes
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

	return nil
}

// SEAL OF APPROVAL
func (d *Database) ReadAllUserRecipes() ([]Recipe, error) {
	fmt.Println("read all user recipes")
	// Establish a connection to the Azure SQL Database
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return []Recipe{}, err
    }

    tsql := fmt.Sprintf(`
    SELECT Title, Ingredients, Instructions, RecipeID, UserName FROM dbo.recipes;
    `)

	ctx := context.Background()
	rows, err := db.QueryContext(
		ctx,
		tsql,
	)

	if err != nil {
		fmt.Println("error on all user recipe read")
		return []Recipe{}, err
	}

	var recipes []Recipe

	// Iterate through the rows and create a slice of Recipe structs
	for rows.Next() {
		var title, ingredientsStr, instructions, recipeID, userName string
		err := rows.Scan(&title, &ingredientsStr, &instructions, &recipeID, &userName)
		if err != nil {
			fmt.Println("row error")
			return []Recipe{}, err
		}

		// Convert the JSON string of ingredients to a slice
		var ingredients []string
		err = json.Unmarshal([]byte(ingredientsStr), &ingredients)
		if err != nil {
			fmt.Println("Unmarshal failure")
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

	fmt.Println(recipes)
	return recipes, nil
}

// SEAL OF APPROVAL
func (d *Database) ReadCurrUserRecipes(currUser User) ([]Recipe, error) {
	// Establish a connection to the Azure SQL Database
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return []Recipe{}, err
    }

    tsql := fmt.Sprintf(`
    SELECT Title, Ingredients, Instructions, RecipeID, UserName FROM dbo.recipes
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
		// return []Recipe{}, err
		panic(err)
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

//waiting for jason and recipes combination
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

func (d *Database) FindFavoriteRecipes(currUser User, routingRecipes []Recommendation) []Recommendation {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return routingRecipes // Return the original list if there's a database error
	}
	defer db.Close()
	
	var count int
	tsql := fmt.Sprintf(`
	SELECT COUNT(*) FROM dbo.user_favorite_recipes
	WHERE RecipeID = @RecipeID
	AND UserName = @UserName;
	`)
	ctx := context.Background()

	// Iterate through the routingRecipes and check if each recipe is a favorite for the user
	for i := range routingRecipes {
		rows, err := db.QueryContext(
			ctx,
			tsql,
			sql.Named("RecipeID", routingRecipes[i].R.RecipeID),
			sql.Named("UserName", currUser.UserName),
		)
		for rows.Next() {
			err = rows.Scan(&count)
		}
		if err != nil {
			count = 0
		}
		// Set the UserFavorite field based on the query result
		routingRecipes[i].R.UserFavorite = count > 0
	}

	return routingRecipes
}

// SEAL OF APPROVAL
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

// SEAL OF APPROVAL
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

// SEAL OF APPROVAL
func (d *Database) UserFromCookie(cookie string) (User, error) {
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return User{}, err
    }

	var returnUser User
	var userName string

	tsql := fmt.Sprintf(`
	SELECT UserName FROM dbo.user_cookies
	WHERE Cookie = @Cookie;
	`)

	ctx := context.Background()
    // Execute query
    row, err := db.QueryContext(
        ctx,
        tsql,
        sql.Named("Cookie", cookie),
	)
	
	if err != nil {
        return User{}, err
    }
	for row.Next() {
		err = row.Scan(&userName)
	}
	// Retrieve user details based on the username
	returnUser, err = d.ReadUserDatabase(userName)
	return returnUser, nil
}

// SEAL OF APPROVAL
func (d *Database) ReadList(currUser User) (List, error) {
	// Establish a connection to the Azure SQL Database
	var err error
    db, err := AzureOpenDatabase()

    if db == nil {
        fmt.Println("Failed to open database")
        return List{}, err
    }

	list := List{
		ShoppingList: make([]FoodItem, 0),
	}

	tsql := fmt.Sprintf(`
	SELECT FoodName, Quantity FROM dbo.user_lists
	WHERE UserName = @UserName;
	`)

	ctx := context.Background()
	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("UserName", currUser.UserName),
	)

	for rows.Next() {
		var foodName string
		var quantity int
		err = rows.Scan(&foodName, &quantity)

		foodItem := FoodItem{
			Name:     foodName,
			Quantity: quantity,
		}
		list.ShoppingList = append(list.ShoppingList, foodItem)
	}

	return list, err
}

// SEAL OF APPROVAL
func (d *Database) WriteList(newItem FoodItem, currUser User) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return nil
    }

    tsql := `
        INSERT INTO dbo.user_lists (UserName, FoodName, FoodType, Quantity)
        VALUES (@UserName, @FoodName, @FoodType, @Quantity);
    `

    stmt, err := db.Prepare(tsql)
    if err != nil {
        return nil
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		sql.Named("UserName", currUser.UserName),
		sql.Named("FoodName", newItem.Name),
		sql.Named("FoodType", newItem.FoodType),
		sql.Named("Quantity", newItem.Quantity),
	)

	if err != nil {
        return err
    }

    return nil
}

// Return when list has updating functionality
func (d *Database) UpdateListItem(newItem FoodItem, currUser User) error {
	var err error
    db, err := AzureOpenDatabase()

    ctx := context.Background()

    if db == nil {
        fmt.Println("Failed to open database")
        return err
    }

    tsql := fmt.Sprintf(`
        UPDATE dbo.user_lists
		SET Quantity = @Quantity
		Where UserName = @UserName
		AND FoodName = @FoodName;
    `)

	stmt, err := db.Prepare(tsql)
    if err != nil {
        
        return err
    }
    defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
        sql.Named("Quantity", newItem.Quantity),
        sql.Named("UserName", currUser.UserName),
        sql.Named("FoodName", newItem.Name),
    )

    if err != nil {
        return err
    }

    AzureSQLCloseDatabase();
    return nil


}
