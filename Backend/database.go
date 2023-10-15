package main

import (
	"github.com/microsoft/go-mssqldb/azuread"
	"database/sql"
	"fmt"

)

var server = "mealdealz.database.windows.net"
var port = 1433

var database = "MealDealz-db"

type Storage interface {
	GetPantry() (*Pantry, error)
}

type AzureDatabase struct {
	db *sql.DB
}

func NewAzureDatabase() (*AzureDatabase, error) {
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;",
	server, password, port, database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	
}