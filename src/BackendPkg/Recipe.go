package BackendPkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Recipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Instructions string  `json:"instructions"`
	PictureLink string   `json:"picture_link"`
}

type Recipes struct {
	Recipes map[string]Recipe `json:"recipes"`
}

func GetRecipes() ([]Recipe, error) {
	var recipes []Recipe

	file, err := ioutil.ReadFile("recipes.json")
	if err != nil {
		log.Fatal(err)
	}

	var data Recipes
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	for _, recipe := range data.Recipes {
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

