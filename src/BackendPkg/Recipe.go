package BackendPkg

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Recipes struct {
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	Instructions   string   `json:"instructions"`
	Ingredients    []string `json:"ingredients"`
	Title          string   `json:"title"`
	PictureLink    string   `json:"picture_link"`
}

type Colors struct {
	Colors []Color `json:"colors"`
}

type Color struct {
	Color          string   `json:"color"`
	Value          string   `json:"value"`

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


