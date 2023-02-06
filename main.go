package main

import (
	"BackendPkg"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to out Sprint 1 demo!")
	userPantry := BackendPkg.Pantry{
		TimeLastUpdated: time.Now(),
	}
	userPantry.AddToPantry()
	userPantry.DisplayPantry()
	userPantry.RemoveFromPantry()
	userPantry.DisplayPantry()
}
