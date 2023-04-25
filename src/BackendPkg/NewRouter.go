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

func RoutData(ctx context.Context, currUser User){

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

    // wait for cancellation signal
    <-ctx.Done()

    // shutdown server
    if err := server.Shutdown(context.Background()); err != nil {
        log.Fatal(err)
    }
}

// LISTEN FUNCTIONS

func handleSignup(w http.ResponseWriter, r *http.Request) {

	// verify POST request from frontend
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// verify post response
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// translate body to ASCII
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// declare new struct to match JSON
	type LoginUser struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var currUser LoginUser

	// unmarshal the JSON data
	err = json.Unmarshal(body, &currUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create a new user with the POST data
	activeUser := User{
		Password: currUser.Password,
		UserName: currUser.UserName,
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

	// listen for user infinitely 
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("Error starting server: %s", err.Error())
    }
}
