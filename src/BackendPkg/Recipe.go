package BackendPkg

import (
    "encoding/json"
    //"fmt"
    "io/ioutil"
    //"strings"
    //"regexp"
)

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

    //re := regexp.MustCompile(`("[^"]+"),([^"]+")`)
    //file = re.ReplaceAll(file, []byte("$1;$2"))

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



