# Delete these lol

```
func (d *Database) StorePublixDatabase(f []FoodItem) error {
	var err error
	db, err := AzureOpenDatabase()

	ctx := context.Background()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	tsql := `
		INSERT INTO dbo.deals_data (Store, foodName, saleDetails)
		VALUES (@store, @foodName, @saleDetails);
	`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range f {
		_, err = stmt.ExecContext(ctx,
			sql.Named("store", "Publix"),
			sql.Named("foodName", item.Name),
			sql.Named("saleDetails", item.SaleDetails),
		)
	}

	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) StoreWalmartDatabase(f []FoodItem) error {
	var err error
	db, err := AzureOpenDatabase()

	ctx := context.Background()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	tsql := `
		INSERT INTO dbo.deals_data (Store, foodName, saleDetails)
		VALUES (@store, @foodName, @saleDetails);
	`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range f {
		_, err = stmt.ExecContext(ctx,
			sql.Named("store", "Walmart"),
			sql.Named("foodName", item.Name),
			sql.Named("saleDetails", item.SaleDetails),
		)
	}

	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}
```

```
func (d *Database) ClearPublixDeals() error {
	var err error
	db, err := AzureOpenDatabase()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	// Define the SQL DELETE statement
	query := "DELETE FROM dbo.deals_data WHERE store = 'Publix'"

	// Execute the DELETE statement
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}

func (d *Database) ClearWalmartDeals() error {
	var err error
	db, err := AzureOpenDatabase()

	if db == nil {
		fmt.Println("Failed to open database")
		return err
	}

	// Define the SQL DELETE statement
	query := "DELETE FROM dbo.deals_data WHERE store = 'Walmart'"

	// Execute the DELETE statement
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	AzureSQLCloseDatabase();
	return nil
}
```

```
func (d *Database) StorePubixScrapedTime(t time.Time) error {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// Define the SQL INSERT statement for the "dbo.deals_data" table
	query := "INSERT INTO dbo.deals_data (store, foodName, saleDetails) VALUES (?, ?, ?)"
	store := "Publix" // Assuming "Publix" is the store name

	// Execute the INSERT statement
	_, err = db.Exec(query, store, t.Format("2006-01-02 15:04:05"), "Scraped time")
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) StoreWalmartScrapedTime(t time.Time) error {
	// Establish a connection to the Azure SQL Database
	db, err := AzureOpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// Define the SQL INSERT statement for the "dbo.deals_data" table
	query := "INSERT INTO dbo.deals_data (store, foodName, saleDetails) VALUES (?, ?, ?)"
	store := "Walmart" // Assuming "Walmart" is the store name

	// Execute the INSERT statement
	_, err = db.Exec(query, store, t.Format("2006-01-02 15:04:05"), "Scraped time")
	if err != nil {
		return err
	}

	return nil
}
```


