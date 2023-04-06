package BackendPkg

import "strings"

type Reccomendation struct {
	r Recipe
	itemsInPantry []FoodItem
	itemsOnSale []FoodItem
}

func BestRecipes(userPantry Pantry, allRecipes []Recipe, deals []FoodItem) []Reccomendation{
	maxScore := len(userPantry.FoodInPantry)
	var scoreList []int
	
	for i := 0; i < len(userPantry.FoodInPantry); i++{
		for j:= 0; j < len(allRecipes); j++{
			tempScore := 0
			for k:=0; k < len(allRecipes[j].Ingredients); k++{
				if(strings.Contains(allRecipes[j].Ingredients[k],userPantry.FoodInPantry[i].Name)){
					tempScore++
				}
			}
			scoreList[j] = tempScore
		}
	}
}