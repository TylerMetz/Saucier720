package main

import (
	"fmt"
	"example.com/package/BackendPkg"
	"time"
)

func main() {
	fmt.Println("Welcome to out Sprint 1 demo!")
	userPantry := BackendPkg.Pantry{
		TimeLastUpdated: time.Now(),
	}
	userPantry.AddToPantry()
	userPantry.AddToPantry()
	userPantry.AddToPantry()
	userPantry.DisplayPantry()
	userPantry.RemoveFromPantry()
	userPantry.RemoveFromPantry()
	userPantry.DisplayPantry()
}
