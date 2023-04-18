package BackendPkg

import (
    "encoding/json"
    //"fmt"
    "io/ioutil"
    "strings"
    //"regexp"
)

type Recipe struct {
    Instructions   string   `json:"instructions"`
    Ingredients    []string `json:"ingredients"`
    Title          string   `json:"title"`
    PictureLink    *string  `json:"picture_link"`
}

func GetRecipes() ([]Recipe, error) {
    file, err := ioutil.ReadFile("recipes.json")
    if err != nil {
        return nil, err
    }

    var recipes map[string]Recipe
    if err := json.Unmarshal(file, &recipes); err != nil {
        return nil, err
    }

    for _, recipe := range recipes {
        for i, ingredient := range recipe.Ingredients {
            recipe.Ingredients[i] = strings.ReplaceAll(ingredient, ",", ";")
        }
    }

    result := make([]Recipe, 0, len(recipes))
    for _, recipe := range recipes {
        result = append(result, recipe)
    }

    return result, nil
}



