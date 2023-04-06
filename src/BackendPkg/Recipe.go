package BackendPkg

import (

    "encoding/json"
    //"fmt"
    "io/ioutil"
)

type Recipes struct {
	Recipes []Recipe `json:"recipes"`
}

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

    result := make([]Recipe, 0, len(recipes))
    for _, recipe := range recipes {
        result = append(result, recipe)
    }

    return result, nil
}


func ReadInAllRecipes() (Recipes, error) {

	file, err := os.Open("recipes.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened recipes.json")

	byteValue, _ := ioutil.ReadAll(file)
	var recipes Recipes

	json.Unmarshal(byteValue, &recipes)

	for i := 0; i < len(recipes.Recipes); i++ {
        fmt.Println("User Type: " + recipes.Recipes[i].Title)
    }
	return recipes, nil

}

