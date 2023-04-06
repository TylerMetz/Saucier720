package BackendPkg

import (
	"strings"
	"sort"
)

type Reccomendation struct {
	R Recipe
	ItemsInPantry []FoodItem
	ItemsOnSale []FoodItem
}

func BestRecipes(userPantry Pantry, allRecipes []Recipe, deals []FoodItem) []Reccomendation{
	var returnReccomendations []Reccomendation
	
	var scoreList []int
	
	// ranks recipes based off of what is in pantry 
	for i := 0; i < len(userPantry.FoodInPantry); i++{
		for j:= 0; j < len(allRecipes); j++{
			tempScore := 0
			for k:=0; k < len(allRecipes[j].Ingredients); k++{
				if(strings.Contains(allRecipes[j].Ingredients[k],userPantry.FoodInPantry[i].Name)){
					tempScore++
				}
			}
			scoreList = append(scoreList, tempScore)
		}
	}

	// create a slice of indices, with the same length as the scoreList
    indices := make([]int, len(scoreList))
    for i := range indices {
        indices[i] = i
    }

    // sort the indices based on the corresponding scores in scoreList
    sort.Slice(indices, func(i, j int) bool {
        return scoreList[indices[i]] > scoreList[indices[j]]
    })

    // create a slice to hold the highest score indices, with a length of 30
    highestScoreIndices := make([]int, 30)

    // copy the first 30 indices (or all of the indices, if there are fewer than 30) into the highestScoreIndices slice
    copy(highestScoreIndices, indices[:min(30, len(indices))])

	// finding corresponding recipes at highest scoring indices 
	for m := 0; m < len(allRecipes); m++{
		for n := 0; n < len(highestScoreIndices); n++{
			if(m == highestScoreIndices[n]){
				// get items in pantry
				var pantryItemsInRecipe []FoodItem
				var dealsItemsInRecipe []FoodItem

				// check which food items are actually contained in recipe
				for i := 0; i < len(userPantry.FoodInPantry); i++{
					for j:=0; j < len(allRecipes[m].Ingredients); j++{
						if(strings.Contains(allRecipes[m].Ingredients[j],userPantry.FoodInPantry[i].Name)){
							pantryItemsInRecipe = append(pantryItemsInRecipe, userPantry.FoodInPantry[i])
						}
					}
				}
				newRecc := Reccomendation{
					R: allRecipes[m],
					ItemsInPantry: pantryItemsInRecipe,
					ItemsOnSale: dealsItemsInRecipe,
				}
				returnReccomendations = append(returnReccomendations, newRecc)
			}
		} 
	}

	return returnReccomendations

}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}



