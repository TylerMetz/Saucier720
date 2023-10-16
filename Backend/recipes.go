package main

import (
	"fmt"
	"sort"
	"strings"
	"golang.org/x/exp/slices"
)

type Recommendation struct {
	R             Recipe
	ItemsInPantry []Ingredient
	ItemsOnSale   []Ingredient
}

type RecomendedRecipes struct {
	Recipes 		[]Recipe 		`json:"R"`
	ItemsInPantry	[]Ingredient 	`json:"ItemsInPantry"`
	ItemsOnSale		[]Ingredient 	`json:"ItemsOnSale"`
}

type Recipe struct {
    RecipeID        string   	`json:"recipeID"`
    Title          	string   	`json:"title"`
    Ingredients    	[]string 	`json:"ingredients"`
    Instructions   	string   	`json:"instructions"`
    RecipeAuthor   	string 		`json"recipeAuthor"`
}

func GetRecomendedRecipes(username string) (RecomendedRecipes, error) {
	
	
	return RecomendedRecipes{}, nil
}


func AllRecipesWithRelatedItems(userPantry Pantry, allRecipes []Recipe, deals []Ingredient) []Recommendation {
	
	// return val
	var returnRecommendations []Recommendation
	
	for m := 0; m < len(allRecipes); m++ {
		// get items in pantry
		var pantryItemsInRecipe []Ingredient
		var dealsItemsInRecipe []string

		// check which food items are actually contained in recipe
		for i := 0; i < len(userPantry.Ingredients); i++ {
			for j := 0; j < len(allRecipes[m].Ingredients); j++ {
				if strings.Contains(allRecipes[m].Ingredients[j], userPantry.Ingredients[i].Name) {
					if !slices.Contains(pantryItemsInRecipe, userPantry.Ingredients[i]) {
						pantryItemsInRecipe = append(pantryItemsInRecipe, userPantry.Ingredients[i])
					}
				}
			}
		}

		// check which deals can be recommended
		for i := 0; i < len(deals); i++ {
			dealWords := strings.Split((deals[i].Name), " ")
			if len(dealWords) > 3 {
				dealWords = dealWords[len(dealWords)-3:]
			}
			for j := 0; j < len(allRecipes[m].Ingredients); j++ {
				for k := 0; k < len(dealWords); k++ {
					if strings.Contains(allRecipes[m].Ingredients[j], (" " + dealWords[k] + " ")) {
						if (" "+dealWords[k]+" ") != " or " && (" "+dealWords[k]+" ") != " and " && (" "+dealWords[k]+" ") != " the " && (" "+dealWords[k]+" ") != " 1 " && (" "+dealWords[k]+" ") != " 2 " && (" "+dealWords[k]+" ") != " 3 " && (" "+dealWords[k]+" ") != " ground " && (" "+dealWords[k]+" ") != " AND " && (" "+dealWords[k]+" ") != " Any " && (" "+dealWords[k]+" ") != " ANY " && (" "+dealWords[k]+" ") != " Sauce " && (" "+dealWords[k]+" ") != " gallon " && (" "+dealWords[k]+" ") != " mix " && (" "+dealWords[k]+" ") != " organic " && (" "+dealWords[k]+" ") != " size " && (" "+dealWords[k]+" ") != " own " && (" "+dealWords[k]+" ") != " alternative " {
							if !slices.Contains(dealsItemsInRecipe, deals[i].Name) {
								dealsItemsInRecipe = append(dealsItemsInRecipe, deals[i].Name)
							}
						}
					}
				}
			}
		}

		var realDealz []Ingredient
		for i := 0; i < len(dealsItemsInRecipe); i++ {
			tempItem := Ingredient{
				Name: dealsItemsInRecipe[i],
			}
			realDealz = append(realDealz, tempItem)
		}

		newRecc := Recommendation{
			R:             allRecipes[m],
			ItemsInPantry: pantryItemsInRecipe,
			ItemsOnSale:   realDealz,
		}
		returnRecommendations = append(returnRecommendations, newRecc)
	}

	// I realized it was returning the recommendations slice backwards so I inverted it in backend
	invertSlice(returnRecommendations)
	return returnRecommendations
}

func ReturnRecipesWithHighestPercentageOfOwnedIngredients(userPantry Pantry, recipesArr []Recipe, numRecipesToReturn int, deals []Ingredient) []Recommendation {
	recipeRatios := make([]struct {
		recipe      Recipe
		ingredientRatio float64
	}, 0)

	for _, recipe := range recipesArr {
		ownedIngredients := 0
		for _, ingredient := range recipe.Ingredients {
			for _, pantryItem := range userPantry.Ingredients {
				if strings.Contains(ingredient, pantryItem.Name) {
					ownedIngredients++
					break
				}
			}
		}

		ingredientRatio := float64(ownedIngredients) / float64(len(recipe.Ingredients))
		recipeRatios = append(recipeRatios, struct {
			recipe         Recipe
			ingredientRatio float64
		}{recipe: recipe, ingredientRatio: ingredientRatio})
	}

	// Sort recipes by ingredient ratio in descending order
	for i := 0; i < len(recipeRatios)-1; i++ {
		for j := i + 1; j < len(recipeRatios); j++ {
			if recipeRatios[i].ingredientRatio < recipeRatios[j].ingredientRatio {
				recipeRatios[i], recipeRatios[j] = recipeRatios[j], recipeRatios[i]
			}
		}
	}

	// Select top numRecipesToReturn recipes
	var resultRecipes []Recipe
	for i := 0; i < numRecipesToReturn && i < len(recipeRatios); i++ {
		resultRecipes = append(resultRecipes, recipeRatios[i].recipe)
	}

	// pass highest rated recipes into func to get related deals and pantry items
	returnRecommendation := AllRecipesWithRelatedItems(userPantry, resultRecipes, deals)
	
	invertSlice(returnRecommendation)
	return returnRecommendation
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// flip the slice
func invertSlice(s []Recommendation) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func OutputRecommendations(r []Recommendation) {
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i].R.Title)
		fmt.Println("From Pantry:")
		for j := 0; j < len(r[i].ItemsInPantry); j++ {
			fmt.Println(r[i].ItemsInPantry[j].Name)
		}
		fmt.Println("From Deals:")
		for k := 0; k < len(r[i].ItemsOnSale); k++ {
			fmt.Println(r[i].ItemsOnSale[k].Name)
		}
	}
}
