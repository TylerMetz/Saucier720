package main

import (
	"BackendPkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestOne(t *testing.T) {

	testFoodItem := BackendPkg.FoodItem{
		Name:        "peanut butter",
		StoreCost:   369.99,
		OnSale:      true,
		SaleDetails: "BOGO",
		Quantity:    10,
	}
	testFoodItem2 := BackendPkg.FoodItem{
		Name:        "jelly",
		StoreCost:   1.0,
		OnSale:      false,
		SaleDetails: "N/A",
		Quantity:    30,
	}
	testFoodItem3 := BackendPkg.FoodItem{
		Name:        "bread",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}

	testItems := []interface{}{testFoodItem3, testFoodItem2, testFoodItem}

	// convert items to be translated into json
	itemsInJson, _ := json.Marshal(testItems)

	// create a router to output items to the port
	testRouter := BackendPkg.Router{
		Name:             "test1",
		ItemsToBeEncoded: testItems,
	}

	// display item on port in background
	go testRouter.Rout("/api/Pantry","8080")

	//pull data from local host port
	resp, err := http.Get("http://localhost:8080/api/Pantry")
	if err != nil {
		fmt.Println("Error!")
	}
	defer resp.Body.Close()

	// make website data a string
	body, err := ioutil.ReadAll(resp.Body)
	portValue := string(body)

	// if the test doesn't pass
	if portValue != string(itemsInJson) {
		t.Errorf("Result was incorrect, got: %s, want: %s.", string(itemsInJson), portValue)
	}

}
