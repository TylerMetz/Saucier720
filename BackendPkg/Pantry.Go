package Backend

import "fmt"

type Pantry struct{
	foodInPantry []FoodItem; //slice for all items in pantry
	dateLastUpdates int; //MMDDYYYY
}

// displays all pantry using fmt
func (p Pantry) DisplayPantry(){
	fmt.Println("User's Current Pantry: ")
	for i, value := range p.foodInPantry{
		fmt.Println("Item", i+1, ":", value.name)
	}
}

// update ingredients list by reference, add new ingredient parament to list
func (p *Pantry) AddToPantry(){
	
	// gets name of item in pantry from user
	fmt.Println("What Item are you looking to add to your pantry?")
	var name string
	fmt.Scanln(name)

	// adds new item to fooditem slice
	p.foodInPantry = append(p.foodInPantry, NewPantryItem(new))

}

// update ingredients list by reference, remove specific parameter ingredient from list
func (p *Pantry) RemoveFromPantry(){
	
	// gets name of item in pantry to remove from user
	fmt.Println("What Item are you looking to remove from your pantry?")
	var name string
	fmt.Scanln(name)

	// cycles through slice to find item to remove
	var foundIndex int = 0
	for i, value := range p.foodInPantry{
		// finds and saves index of item getting removed
		if(value.name == name){
			foundIndex = i
		}
	}

	// if index is in range, remove from slice
	if foundIndex < len(p.foodInPantry) && foundIndex >= 0{
        return append(p.foodInPantry[0:foundIndex], p.foodInPantry[foundIndex+1:len(p.foodInPantry)]...);
    }
}