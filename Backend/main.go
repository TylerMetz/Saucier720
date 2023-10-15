package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Backend Starting: ")

	db, err := NewAzureDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	server := NewMealDealzServer(":1433", db)
	server.Run()
}
