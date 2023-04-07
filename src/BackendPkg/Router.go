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

func ListenPantry(endLink string, port string) {

    // creates new router
    route := mux.NewRouter()
    route.HandleFunc(endLink, PantryItemPostResponse).Methods("POST")

    // enables alternate hosts for CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:4200"},
        AllowCredentials: true,
    })

    // log.Println("Listening...")
    handler := c.Handler(route)
    http.ListenAndServe(port, handler)

}

func PantryItemPostResponse(response http.ResponseWriter, request *http.Request) {

	// fmt.Println(string(jsonResponse)) // used to test

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("hi"))

    // Decode JSON payload into Ingredient struct
    var foodItem FoodItem
    err := json.NewDecoder(request.Body).Decode(&foodItem)
    if err != nil {
        http.Error(response, err.Error(), http.StatusBadRequest)
        return
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

    testDatabase.InsertPantryItemPost(foodItem)

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

func ListenForAllPosts(){
	// all listen functions go in here
	for true{
	
		// listens for new pantry item
		ListenPantry("/api/NewPantryItem", "8083")

	}
}

func ListenForNewUser(){

	// reads from signup page
	resp, _ := http.Get("http://localhost:8085/api/Signup")

	// stores data as new user
	if(resp != nil){
		var user User
		json.NewDecoder(resp.Body).Decode(&user)
		defer resp.Body.Close()

		// creates database object to store info in MealDealz.sb
		newUserDatabase := Database{
			Name: "MealDealz Database",
		}

		// store the new user in the database
		newUserDatabase.StoreUserDatabase(user)
	}
}
