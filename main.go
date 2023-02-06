package main

import (
	"fmt"
	"https://github.com/TylerMetz/Saucier720/tree/develop-backend/BackendPkg"
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
