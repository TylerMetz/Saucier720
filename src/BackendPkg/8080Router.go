package BackendPkg

import (
	"encoding/json"
	_"fmt"
	"net/http"

	_"time"
	_"github.com/gorilla/mux"
	_"github.com/rs/cors"
	_"io/ioutil"
	"log"
	_"sync"
	_"context"
	//"strings"
	//"bytes"
)

// GLOBAL VARIABLES
var pantryInterface []interface{}
var dealsInterface []interface{}
var recipesInterface []interface{}
var backendDatabase Database

// ALL ROUTING FUNCTIONS

func handlePantry(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	response.Header().Set("Access-Control-Allow-Methods", "GET")
    json.NewEncoder(response).Encode(pantryInterface)
}

func handleDeals(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	response.Header().Set("Access-Control-Allow-Methods", "GET")
    json.NewEncoder(response).Encode(dealsInterface)
}

func handleRecipes(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	response.Header().Set("Access-Control-Allow-Methods", "GET")
	json.NewEncoder(response).Encode(recipesInterface)
}

func formatData(currUser User){
	// read from .db file and prepare all data to be routed

	// lock the user pantry data
	dataMutex.Lock()

	// save all user pantry data to global variable
	var pantryInterfaceRefresh []interface{}
	pantryInterface = pantryInterfaceRefresh
	for i := 0; i < len(backendDatabase.GetUserPantry(currUser.UserName).FoodInPantry); i++{
		pantryInterface = append(pantryInterface, backendDatabase.GetUserPantry(currUser.UserName).FoodInPantry[i])
	}

	// unlock the data
	dataMutex.Unlock()

	// save all deals data to global variable
	for i := 0; i < len(backendDatabase.ReadPublixDatabase()); i++{
		dealsInterface = append(dealsInterface, backendDatabase.ReadPublixDatabase()[i])
	}

	// save all recipes data to global variable
	userRecList := BestRecipes(backendDatabase.GetUserPantry(currUser.UserName), backendDatabase.ReadRecipes(), backendDatabase.ReadPublixDatabase())
	var recipesInterfaceRefresh []interface{}
	recipesInterface = recipesInterfaceRefresh
	for i := 0; i < len(userRecList); i++ {
		// sends recipes, items in recipe, and deals related 
		recipesInterface = append(recipesInterface, userRecList[i])
	}
}

func RoutData(currUser User){

	// setup all global variables to be routed
	formatData(currUser)

	// create server
	server := http.Server{
        Addr: ":8080",
    }

	// handle functions
    http.HandleFunc("/api/Pantry", handlePantry)
    http.HandleFunc("/api/Recipes", handleRecipes)
	http.HandleFunc("/api/Deals", handleDeals)

	// listen and serve on server
    go func() {
        if err := server.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

	select {}

}
