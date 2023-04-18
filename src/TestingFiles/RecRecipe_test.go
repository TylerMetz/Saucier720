package main

import (
    "testing"
	"BackendPkg"
	"fmt"
)

func TestBestRecipes(t *testing.T) {
    // Create a test pantry
    testPantry := BackendPkg.Pantry{
        FoodInPantry: []BackendPkg.FoodItem{
            {Name: "spaghetti"},
            {Name: "eggs"},
            {Name: "bacon"},
            {Name: "pamesan cheese"},
            {Name: "salt"},
			{Name: "black pepper"},
        },
    }

	d := BackendPkg.Database{}
	d.WriteRecipes()

	testRecipes := d.ReadRecipes()



    // Create some test deals
    testDeals := [] BackendPkg.FoodItem{
        {Name: "eggs"},
        {Name: "cheese"},
        {Name: "bread"},
        {Name: "chocolate chips"},
    }

    // Call the BestRecipes function with the test data
    recommendations :=  BackendPkg.BestRecipes(testPantry, testRecipes, testDeals)
	//recommendations = recommendations[:2]

	fmt.Print(recommendations)

    // Check that the returned recommendations are in the correct order
    expectedOrder := []string{"Spaghetti Carbonara", "Caesar Salad"}
	for i := 1; i >= 0; i-- {
		if recommendations[i].R.Title != expectedOrder[i] {
            t.Errorf("Recommendation %d was %s; expected %s", i, recommendations[i].R.Title, expectedOrder[i])
        }
	}
}