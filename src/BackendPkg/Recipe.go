package BackendPkg

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

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


func ReadInAllRecipes() (Colors, error) {
	//var recipes []Recipe

	file, err := os.Open("colors.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened colors.json")

	byteValue, _ := ioutil.ReadAll(file)
	var colors Colors

	json.Unmarshal(byteValue, &colors)
	fmt.Println("Color: " + colors.Colors[0].Color)
	fmt.Println("Value: " + colors.Colors[0].Value)

	return colors, nil
}


