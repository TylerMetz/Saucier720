package BackendPkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"time"

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
	jsonResponse, jsonError := json.Marshal(t.ItemsToBeEncoded)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse)) // used to test

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)

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
    //fmt.Println(newFoodItem.Name)
	//InsertPantryItemPost() -- insert into backend database

    testDatabase := Database{
		Name: "MealDealz Database",
	}

	//foodSlice := []FoodItem {newFoodItem};

    // testUser := User{
	// 	FirstName: "Eddie",
	// 	LastName: "Menello",
	// 	Email: "Edward@gmail.com",
	// 	UserName: "Eddiefye69",
	// 	Password: "ILoveGraham420",
	// 	UserPantry: Pantry{
	// 		FoodInPantry: foodSlice,
	// 		TimeLastUpdated: time.Now(),
	// 	},
	// }

    testDatabase.InsertPantryItemPost(newFoodItem)

    //go RoutUserPantry(testDatabase, testUser)

    // Return a 200 OK response
    response.WriteHeader(http.StatusOK)
}

func RoutUserPantry(d Database, u User){
	
	// read from .db file and output test user's pantry to frontend
	var testFoodInterface []interface{}
	for i := 0; i < len(d.GetUserPantry(u.UserName).FoodInPantry); i++{
		testFoodInterface = append(testFoodInterface, d.GetUserPantry(u.UserName).FoodInPantry[i])
	}
	// test router
	programRouter := Router{
		Name:             "testRouter",
		ItemsToBeEncoded: testFoodInterface,
	}
	programRouter.Rout("/api/Pantry", ":8080")
}

func ListenForNewUser(){

	// create a router to output items to the port
	testRouter := BackendPkg.Router{
		Name:             "NewUser",
	}
	resp, err := http.Get("http://localhost:8080/api/Signup")

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)

}