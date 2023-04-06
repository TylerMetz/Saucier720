package BackendPkg

import (
	"encoding/json"
	"io/ioutil"
)

type Recipe struct {
	Instructions   string   `json:"instructions"`
	Ingredients    []string `json:"ingredients"`
	Title          string   `json:"title"`
	PictureLink    string   `json:"picture_link"`
}

func ReadInAllRecipes() []Recipe {
	// read the JSON data from the file
    file, _ := ioutil.ReadFile("recipes.json")

    // unmarshal the JSON data into a map[string]Recipe
    var recipes []Recipe
	_ = json.Unmarshal(file, &recipes)

	return recipes 
}


