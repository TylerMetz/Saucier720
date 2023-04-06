package BackendPkg

import (
)

type Reccomendation struct {
	r Recipe
	itemsInPantry []FoodItem
	itemsOnSale []FoodItem
}

func BestRecipes(userPantry Pantry, allRecipes []Recipe, deals []FoodItem) []Reccomendation{
	maxScore := len(userPantry.FoodInPantry)
	var scoreList []int
	
	for i := 0; i < len(userPantry.FoodInPantry); i++{
		
	}
}