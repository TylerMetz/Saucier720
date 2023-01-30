package Backend

import "fmt"

type Pantry struct{
	// List of Ingredients (need to decide data structure)
	dateLastUpdates int; //MMDDYYYY
}

func DisplayPantry(userPantry Pantry){
	// displays all pantry using fmt
}
func AddToPantry(/*List of Ingredients from struct*/, newItem Ingredient){
	// update ingredients list by reference, add new ingredient parament to list
}
func RemoveFromPantry(/*List of Ingredients from struct*/, oldItem Ingredient){
	// update ingredients list by reference, remove specific parameter ingredient from list
}