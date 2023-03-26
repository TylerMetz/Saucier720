package BackendPkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Router struct {
	Name             string
	ItemsToBeEncoded []interface{}
}

func (t *Router) Rout(endLink string, port string) {

	// creates new router
	route := mux.NewRouter()
	route.HandleFunc(endLink, t.sendResponse).Methods("GET")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// log.Println("Listening...")
	handler := c.Handler(route)
	http.ListenAndServe(port, handler)
}

func (t *Router) sendResponse(response http.ResponseWriter, request *http.Request) {

	jsonResponse, jsonError := json.Marshal(t.ItemsToBeEncoded)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse)) // used to test

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func (t *Router) Listen(endLink string, port string) {

    // creates new router
    route := mux.NewRouter()
    route.HandleFunc(endLink, t.sendPostResponse).Methods("POST")

    // enables alternate hosts for CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:4200"},
        AllowCredentials: true,
    })

    // log.Println("Listening...")
    handler := c.Handler(route)
    http.ListenAndServe(port, handler)

}

func (t *Router) sendPostResponse(response http.ResponseWriter, request *http.Request) {
    // Decode JSON payload into Ingredient struct
    var foodItem FoodItem
    err := json.NewDecoder(request.Body).Decode(&foodItem)
    if err != nil {
        http.Error(response, err.Error(), http.StatusBadRequest)
        return
    }

	newFoodItem := FoodItem{
        Name:        foodItem.Name,
        StoreCost:   foodItem.StoreCost,
        OnSale:      foodItem.OnSale,
        SalePrice:   foodItem.SalePrice,
        SaleDetails: foodItem.SaleDetails,
        Quantity:    foodItem.Quantity,
    }

    // Do something with the ingredient struct, e.g. store it in a database
    fmt.Println(newFoodItem.Name)
	//InsertPantryItemPost() -- insert into backend database

    testDatabase := Database{
		Name: "MealDealz Database",
	}

    testDatabase.InsertPantryItemPost(newFoodItem)

    // Return a 200 OK response
    response.WriteHeader(http.StatusOK)
}
