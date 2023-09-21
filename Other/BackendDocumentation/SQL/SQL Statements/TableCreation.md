# MealDealz Schema V2

## Deals Data

```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'deals_data'),
                BEGIN
                    CREATE TABLE deals_data (
                        store VARCHAR(25) PRIMARY KEY,
                        foodName VARCHAR(50),
                        saleDetails VARCHAR(255),
                    )
                END
```

## Scraped Times

```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'scraped_time'),
                BEGIN
                    CREATE TABLE scraped_time (
                        store VARCHAR(25) PRIMARY KEY,
                        time DATETIME,
                    )
                END
```

## User Data

```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'user_data')
                BEGIN
                    CREATE TABLE user_data (
                        UserName VARCHAR(25) PRIMARY KEY,
                        FirstName VARCHAR(30),
                        LastName VARCHAR(30),
                        Email VARCHAR(50),
                        Password VARCHAR(25),
                        DateJoined DATE,
                    )
                END
```

## User Lists

```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'user_lists')
                BEGIN
                    CREATE TABLE user_lists (
                        UserName VARCHAR(25) PRIMARY KEY,
                        FoodName VARCHAR(30),
                        FoodType VARCHAR(30),
                        Quantity INT,
                    )
                END
```

## User Cookies

```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'user_cookies')
                BEGIN
                    CREATE TABLE user_cookies (
                        UserName VARCHAR(25) PRIMARY KEY,
                        Cookie VARCHAR(255),
                    )
                END
```

## User Ingredients

```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'user_ingredients')
                BEGIN
                    CREATE TABLE user_ingredients (
                        UserName VARCHAR(25),
                        FoodName VARCHAR(25),
                        FoodType VARCHAR(25),
                        Quantity INT,
                        PRIMARY KEY (UserName, FoodName)
                    )
                END
```

## User Recipes

```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'user_recipes')
                BEGIN
                    CREATE TABLE user_recipes (
                        RecipeID INT PRIMARY KEY IDENTITY(1,1),
                        Title VARCHAR(25),
                        Ingredients VARCHAR(255),
                        Instructions VARCHAR(255),
                        UserName VARCHAR(25)
                    )
                END
```

## Jason's Recipes
```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'jason_recipes')
                BEGIN
                    CREATE TABLE dbo.jason_recipes (
                    RecipeID int IDENTITY(1,1) PRIMARY KEY,  
                    Title varchar (30),  
                    Ingredients varchar(1500),  
                    Instructions varchar(8000)
                    );
                END
```

## User Favorite Recipes
```SQL
IF NOT EXISTS (SELECT * FROM dbo.tables WHERE name = 'user_favorite_recipes')
                BEGIN
                    CREATE TABLE user_favorite_recipes (
                        RecipeID INT IDENTITY(1,1),
                        UserName VARCHAR(25),
                        PRIMARY KEY (RecipeID, UserName)
                    )
                END
```


