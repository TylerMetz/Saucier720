package Backend

import "fmt"
import "time"

type Pantry struct{
	foodInPantry []FoodItem; //slice for all items in pantry
	timeLastUpdated time.Time; //MMDDYYYY
}

// displays all pantry using fmt
func (p Pantry) DisplayPantry(){
	fmt.Sprintf("User's Current Pantry: (Last Updated on", timeLastUpdated.Format("February 6, 2023, at 15:04"), ")")
	for i, value := range p.foodInPantry{
		fmt.Println("Item", i+1, ":", value.name)
	}
}

// update ingredients list by reference, add new ingredient parament to list
func (p Pantry) AddToPantry(){
	
	// gets name of item in pantry from user
	fmt.Println("What Item are you looking to add to your pantry?")
	var name string
	fmt.Scanln(name)

	// adds new item to fooditem slice
	p.foodInPantry = append(p.foodInPantry, NewPantryItem(new))
	
	//get updated time
	timeLastUpdated = time.Now()
}

// update ingredients list by reference, remove specific parameter ingredient from list
func (p Pantry) RemoveFromPantry(){
	
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
	if foundIndex != len(p.foodInPantry) && foundIndex >= 0{
        p.foodInPantry = append(p.foodInPantry[0:foundIndex], p.foodInPantry[foundIndex+1:len(p.foodInPantry)]...);
		// get time updated
		timeLastUpdated = time.Now()
    } else{
		fmt.Println("Item couldn't be found in your pantry!")
	}

}
