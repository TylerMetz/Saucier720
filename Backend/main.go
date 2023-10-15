package main

import (
	"fmt"
	_"log"
)

func main() {

	fmt.Println("Backend Starting: ")

	db, err := NewAzureDatabase()
	if err != nil {
		fmt.Println("Hi")
	}
	
	if err := db.Init(); err !=nil {
		fmt.Println("Hello")
	}
}
