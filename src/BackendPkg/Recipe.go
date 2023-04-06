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
	var recipes []Recipe

	fileBytes, _ := ioutil.ReadFile("recipes.json")

	data := make(map[string]Recipe)
	_ = json.Unmarshal(fileBytes, &data)

	for _, recipe := range data {
		recipes = append(recipes, recipe)
	}

	return recipes
}


