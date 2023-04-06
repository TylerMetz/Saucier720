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

func ReadInAllRecipes() ([]Recipe, error) {
	var recipes []Recipe

	file, err := ioutil.ReadFile("recipes.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(file), &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}


