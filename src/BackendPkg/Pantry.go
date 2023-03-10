package BackendPkg

import "fmt"
import "time"

type Pantry struct {
	FoodInPantry    []FoodItem // slice for all items in pantry
	TimeLastUpdated time.Time  // MMDDYYYY
}

// displays all pantry using fmt
func (p Pantry) DisplayPantry() {
	fmt.Println("User's Current Pantry: (Last Updated on", p.TimeLastUpdated.Format("September 16, 2006, at 15:04"), ")")

	for i, value := range p.FoodInPantry {
		fmt.Println("Item", i+1, ":", value.Name)
	}
}

// update ingredients list by reference, add new ingredient parament to list
func (p *Pantry) AddToPantry() {

	// gets name of item in pantry from user
	fmt.Println("What Item are you looking to add to your pantry?")
	var name string
	fmt.Scanln(&name)

	// adds new item to fooditem slice

	p.FoodInPantry = append(p.FoodInPantry, NewPantryItem(name)) // called as a function of FoodItem
	//get updated time
	p.TimeLastUpdated = time.Now()
}

// update ingredients list by reference, remove specific parameter ingredient from list
func (p *Pantry) RemoveFromPantry() {

	// gets name of item in pantry to remove from user
	fmt.Println("What Item are you looking to remove from your pantry?")
	var name string
	fmt.Scanln(&name)

	// cycles through slice to find item to remove
	var foundIndex int = 0
	for i, value := range p.FoodInPantry {
		// finds and saves index of item getting removed
		if value.Name == name {
			foundIndex = i
		}
	}

	// if index is in range, remove from slice
	if foundIndex != len(p.FoodInPantry) && foundIndex >= 0 {
		p.FoodInPantry = append(p.FoodInPantry[0:foundIndex], p.FoodInPantry[foundIndex+1:len(p.FoodInPantry)]...)
		// get time updated
		p.TimeLastUpdated = time.Now()
	} else {
		fmt.Println("Item couldn't be found in your pantry!")
	}

}


