package main


import (
    _"encoding/json"
    "io/ioutil"
    "os"
    _"strings"
    "testing"
	"BackendPkg"
)
func TestGetRecipes(t *testing.T) {
    // Create a temporary file with recipe data.
    file, err := ioutil.TempFile("", "recipes.json")
    if err != nil {
        t.Fatalf("failed to create temporary file: %v", err)
    }
    defer os.Remove(file.Name())

    // Write recipe data to the temporary file.
    data := `{
        "recipe1": {
            "instructions": "Fake instructions 1",
            "ingredients": [
                "spaghetti",
                "eggs",
                "bacon",
                "parmesan cheese",
                "salt",
                "black pepper"
            ],
			"title": "Spaghetti Carbonara",
			"picture_link": null
        },
        "recipe2": {
            "instructions": "Fake instruction 2",
            "ingredients": [
                "romaine lettuce",
                "croutons",
                "parmesan cheese",
                "lemon juice",
                "garlic",
                "anchovy paste",
                "olive oil",
                "salt",
                "black pepper"
            ],
			"title": "Caesar Salad",
			"picture_link": null
        }
    }`
	
    if _, err := file.Write([]byte(data)); err != nil {
        t.Fatalf("failed to write to temporary file: %v", err)
    }

    // Call the GetRecipes function to read from the temporary file.
    recipes, err := BackendPkg.GetRecipes()
    if err != nil {
        t.Fatalf("GetRecipes returned an error: %v", err)
    }

    // Check that the returned recipes are correct.
    if got, want := len(recipes), 2; got != want {
        t.Errorf("got %d recipes, want %d", got, want)
    }

    recipe1 := recipes[0]
    if got, want := recipe1.Title, "Spaghetti Carbonara"; got != want {
        t.Errorf("got recipe1 name %q, want %q", got, want)
    }
    if got, want := len(recipe1.Ingredients), 6; got != want {
        t.Errorf("got %d recipe1 ingredients, want %d", got, want)
    }
    if got, want := recipe1.Ingredients[2], "bacon"; got != want {
        t.Errorf("got recipe1 ingredient[2] %q, want %q", got, want)
    }

    recipe2 := recipes[1]
    if got, want := recipe2.Title, "Caesar Salad"; got != want {
        t.Errorf("got recipe2 name %q, want %q", got, want)
    }
    if got, want := len(recipe2.Ingredients), 9; got != want {
        t.Errorf("got %d recipe2 ingredients, want %d", got, want)
    }
    if got, want := recipe2.Ingredients[8], "black pepper"; got != want {
        t.Errorf("got recipe2 ingredient[8] %q, want %q", got, want)
    }
}