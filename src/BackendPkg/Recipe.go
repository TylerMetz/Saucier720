package BackendPkg

import (
    "encoding/json"
    _"fmt"
    "io/ioutil"
    "strings"
    //"regexp"
)

type Recipe struct {
    Instructions   string   `json:"instructions"`
    Ingredients    []string `json:"ingredients"`
    Title          string   `json:"title"`
    PictureLink    *string  `json:"pictureLink"`
    RecipeID        string   `json:"recipeID"`
    UserFavorite    bool     `json:"userFavorite"`
    RecipeAuthor    string `json"recipeAuthor"`
}

func GetJSONRecipes() ([]Recipe, error) {
    file, err := ioutil.ReadFile("recipes.json")
    if err != nil {
        return nil, err
    }

    var recipes map[string]Recipe
    if err := json.Unmarshal(file, &recipes); err != nil {
        return nil, err
    }

    // Eliminate subrecipes
    filteredRecipes := make([]Recipe, 0)

    for _, recipe := range recipes {
        containsSubRecipe := false

        for i, ingredient := range recipe.Ingredients {
            if strings.Contains(strings.ToLower(ingredient), "recipe follows") {
                containsSubRecipe = true
                break
            } else {
                recipe.Ingredients[i] = strings.ReplaceAll(ingredient, ",", ";")
            }
        }
        if !containsSubRecipe {
            filteredRecipes = append(filteredRecipes, recipe)
        }
    }
    return filteredRecipes, nil
}


