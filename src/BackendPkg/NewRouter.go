package BackendPkg

import (
	"encoding/json"
	_"fmt"
	"net/http"

	"time"
	_"github.com/gorilla/mux"
	_"github.com/rs/cors"
	"io/ioutil"
	"log"
	"sync"
	"context"
	//"strings"
	//"bytes"
)

// GLOBAL VARIABLES
var pantryInterface []interface{}
var dealsInterface []interface{}
var recipesInterface []interface{}
var backendDatabase Database
var dataMut sync.Mutex
var NewServers []*http.Server
var wait sync.WaitGroup
var UpdatingPantry bool
var CurrentUser User

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

func FormatData(){
	// read from .db file and prepare all data to be routed

	// lock the user pantry data
	dataMutex.Lock()

	// save all user pantry data to global variable
	var pantryInterfaceRefresh []interface{}
	pantryInterface = pantryInterfaceRefresh
	for i := 0; i < len(backendDatabase.GetUserPantry(CurrentUser.UserName).FoodInPantry); i++{
		pantryInterface = append(pantryInterface, backendDatabase.GetUserPantry(CurrentUser.UserName).FoodInPantry[i])
	}

	// unlock the data
	dataMutex.Unlock()

	// save all recipes data to global variable
	userRecList := BestRecipes(backendDatabase.GetUserPantry(CurrentUser.UserName), backendDatabase.ReadRecipes(), backendDatabase.ReadPublixDatabase())
	var recipesInterfaceRefresh []interface{}
	recipesInterface = recipesInterfaceRefresh
	for i := 0; i < len(userRecList); i++ {
		// sends recipes, items in recipe, and deals related 
		recipesInterface = append(recipesInterface, userRecList[i])
	}
}

func RoutData(ctx context.Context){

    // setup all global variables to be routed
	go func(){
		// save all deals data to global variable
		for i := 0; i < len(backendDatabase.ReadPublixDatabase()); i++{
			dealsInterface = append(dealsInterface, backendDatabase.ReadPublixDatabase()[i])
		}
		// infinitely write pantry and recipe data
		for{
			if(!UpdatingPantry){FormatData()}
		}
	}()
    

    // create server
    server := &http.Server{
        Addr: ":8080",
    }

    // handle functions
    http.HandleFunc("/api/Pantry", handlePantry)
    http.HandleFunc("/api/Recipes", handleRecipes)
    http.HandleFunc("/api/Deals", handleDeals)

    // append the server to the global list
	NewServers = append(NewServers, server)

	// increment the WaitGroup counter
    wait.Add(1)
	// listen and serve until context is cancelled
	func() {
		defer wait.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
}

// LISTEN FUNCTIONS

func handleSignup(w http.ResponseWriter, r *http.Request) {

	// verify POST request from frontend
    if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.WriteHeader(http.StatusOK)
        return
    }

	// set correct headers
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Credentials", "true")

	// translate body into ASCII
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// define type to match JSON data from frontend
	type UserResponse struct {
		User User `json:"user"`
	}
	var newUser UserResponse

	// unmarshal JSON data
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// store the new user
	backendDatabase.StoreUserDatabase(newUser.User)

}

func handleLogin(w http.ResponseWriter, r *http.Request, sessionCookie *string, cookieChanged *bool) {

	// verify POST request from frontend
    if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.WriteHeader(http.StatusOK)
        return
    }

	// set correct headers
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Credentials", "true")


	// translate body to ASCII
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// declare new struct to match JSON
	type LoginUser struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var CurrentUser LoginUser

	// unmarshal the JSON data
	err = json.Unmarshal(body, &CurrentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create a new user with the POST data
	activeUser := User{
		Password: CurrentUser.Password,
		UserName: CurrentUser.UserName,
	}

	// checks if validate user function returned an empty cookie, if not then setts the cookies
	if ValidateUser(activeUser) == "" {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	} else {
		// Set the cookie
		cookie := &http.Cookie{
			Name:     "sessionID",
			Value:    ValidateUser(activeUser),
			Path:     "/",
			Expires: time.Now().Add(7 * 24 * time.Hour), 
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			Domain: "localhost",
		}
		//http.SetCookie(w, cookie)

		// sets cookie changed to true
		*cookieChanged = true

		// write a response to the client
		type response struct {
			Message string `json:"message"`
			Value   string `json:"value"`
		}
		returnResponse := response {
			Message: "Cookie set successfully",
			Value: cookie.Value,
		}
	

		// Allow the 'Set-Cookie' header to be exposed to the frontend
		w.Header().Set("Content-Type","application/json",)

		// Return a response to the frontend
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(returnResponse)

		// Get the new "sessionID" cookie value
		*sessionCookie = cookie.Value

	}

}

func ListenForUser(ctx context.Context, sessionCookie* string, cookieChanged* bool){
	
	// handle the listening functions
	http.HandleFunc("/api/Signup", handleSignup)
    http.HandleFunc("/api/Login", func(response http.ResponseWriter, request *http.Request) {
        handleLogin(response, request, sessionCookie, cookieChanged)
    })

	// creare server
	server := &http.Server{Addr: ":8081"}

    // append the server to the global list
	NewServers = append(NewServers, server)

	// increment the WaitGroup counter
    wait.Add(1)

	// listen for user infinitely 
	func() {
		defer wait.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

}

func handlePantryUpdate(w http.ResponseWriter, r *http.Request) {

	// verify POST request from frontend
    if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.WriteHeader(http.StatusOK)
        return
    }

	// set correct headers
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Credentials", "true")


	// translate POST data to ASCII
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// define type to match JSON data from frontend
	type Ingredient struct {
		FoodItem []FoodItem `json:"pantry"`
	}
	var updatedPantry Ingredient

	// unmarshal JSON data
	err = json.Unmarshal(body, &updatedPantry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update the current user's pantry
	UpdatingPantry = true
	backendDatabase.UpdatePantry(CurrentUser, updatedPantry.FoodItem)
	UpdatingPantry = false

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		FormatData()
	}

}

func handleNewPantryItem(w http.ResponseWriter, r *http.Request) {

	// verify POST request from frontend
    if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.WriteHeader(http.StatusOK)
        return
    }

	// set correct headers
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Credentials", "true")


	// translate POST data to ASCII
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// define type to match JSON data from frontend
	type Ingredient struct {
		FoodItem FoodItem `json:"ingredient"`
	}
	var newItem Ingredient

	// unmarshal JSON data
	err = json.Unmarshal(body, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update the current user's pantry
	UpdatingPantry = true
	backendDatabase.InsertPantryItemPost(CurrentUser, newItem.FoodItem)
	UpdatingPantry = false

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		FormatData()
	}

}

func ListenForData(ctx context.Context){
	
	// handle the listening functions
	http.HandleFunc("/api/UpdatePantry", func(response http.ResponseWriter, request *http.Request) {
        handlePantryUpdate(response, request)
    })
	http.HandleFunc("/api/NewPantryItem", func(response http.ResponseWriter, request *http.Request) {
        handleNewPantryItem(response, request)
    })

	// creare server
	server := &http.Server{Addr: ":8082"}

    // append the server to the global list
	NewServers = append(NewServers, server)

	// increment the WaitGroup counter
    wait.Add(1)

	// listen for user infinitely 
	func() {
		defer wait.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
}

func ShutdownNewServers() {
    for _, server := range NewServers {
        // gracefully shut down the server
        if err := server.Shutdown(context.Background()); err != nil {
            log.Fatal("Server shutdown failed:", err)
        }
    }
    // wait for all the servers to shut down before returning
    wait.Wait()
	NewServers = nil
	return
}


