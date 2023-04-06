package BackendPkg

import (
)

type Reccomendation struct {
	r Recipe
	itemsInPantry []FoodItem
	itemsOnSale []FoodItem
}

func BestRecipes(userPantry Pantry, allRecipes []Recipe, deals []FoodItem) []Reccomendation{
	
}