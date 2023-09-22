package BackendPkg

import (
	"encoding/json"
	"net/http"
	"time"
	"io/ioutil"
	"log"
	"sync"
	"context"
	_"fmt"
)

// GLOBAL VARIABLES
var pantryInterface []interface{}
var dealsInterface []interface{}
var recipesInterface []interface{}
var listInterface []interface{}
var userDashboardInterface []interface{}
var backendDatabase Database
var dataMutex sync.Mutex
var NewServers []*http.Server
var wait sync.WaitGroup
var UpdatingData bool
var CurrentUser User
var StoreSelection string
var StoreDeals []FoodItem
var RoutingRecipesType RecipeType
var RoutingUserDashboardData UserDashboardData
var RecipesRecommendationPool []Recipe

// RECIPE TYPE ENUM
type RecipeType int
const (
	RecommendedRecipes RecipeType = iota
	UserRecipes
	FavoriteRecipes
)

// USER DASHBOARD DATA ENUM
type UserDashboardData int
const (
	// making examples here for what we could output to user dashboard
	TopFiveFavoritesEasiestToMake UserDashboardData = iota
	TopTenFavoritesEasiestToMake
	RecentRecipesViewed
	RecentUsersViewed
)

// ROUTING FUNCTIONS

func handlePantry(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	response.Header().Set("Access-Control-Allow-Methods", "GET")
	// Encode the items as JSON and send the response
    json.NewEncoder(response).Encode(pantryInterface)
}

func handleDeals(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	response.Header().Set("Access-Control-Allow-Methods", "GET")
	// Encode the items as JSON and send the response
    json.NewEncoder(response).Encode(dealsInterface)
}

func handleRecipes(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	response.Header().Set("Access-Control-Allow-Methods", "GET")
	// Encode the items as JSON and send the response
	json.NewEncoder(response).Encode(recipesInterface)
}

func handleListPage(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200") 
	response.Header().Set("Access-Control-Allow-Methods", "GET")
	// Encode the items as JSON and send the response
	json.NewEncoder(response).Encode(listInterface)
}

func handleUserDashboard(response http.ResponseWriter, request *http.Request) {
	// set header and encode items
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200") 
	response.Header().Set("Access-Control-Allow-Methods", "GET")
	// Encode the items as JSON and send the response
	json.NewEncoder(response).Encode(userDashboardInterface)
}

func RoutData(){

    // setup all global variables to be routed
	UpdateAllData();
	
    // create server
    server := &http.Server{
        Addr: ":8080",
    }

    // handle functions
    http.HandleFunc("/api/Pantry", handlePantry)
    http.HandleFunc("/api/Recipes", handleRecipes)
    http.HandleFunc("/api/Deals", handleDeals)
	http.HandleFunc("/api/List", handleListPage)
	http.HandleFunc("/api/UserDashboard", handleUserDashboard)	

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
	UpdatingData = true;
	StoreUserDatabase(newUser.User)
	UpdatingData = false;

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
	if ValidateUser(activeUser) == "" { // RILEY what is this validate user
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	} else {
		// set updating data to true
		UpdatingData = true

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

		///////////////////////////////////////////////////////////////////////////////////// DEFAULTS /////////////////////////////////////////////////////////////////////////////////////////////////

		// set recommended recipes to pull from MealDealz classsic recipes by default
		RecipesRecommendationPool, _ = backendDatabase.ReadJSONRecipes()

		// set recipes that are routed to recommended by default
		RoutingRecipesType = RecommendedRecipes

		// set user dashboard default to be top 5 recipes
		RoutingUserDashboardData = TopFiveFavoritesEasiestToMake

		// set default deals to Publix 
		StoreSelection = "Publix"

		////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		// allow data to be routed again
		UpdatingData = false;

	}

}

func handleLogout(w http.ResponseWriter, r *http.Request, sessionCookie *string) {

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

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {

		// set cookie to be null
		*sessionCookie = ""

		// set UpdatingData to false so that data stops being acitvely updated
		UpdatingData = true;

		// set all routing data to be empty
		var interfaceRefresh []interface{}
		pantryInterface = interfaceRefresh
		dealsInterface = interfaceRefresh
		recipesInterface = interfaceRefresh
		listInterface = interfaceRefresh
		userDashboardInterface = interfaceRefresh
		
	}

}

func ListenUserInfo(sessionCookie* string, cookieChanged* bool){
	
	// handle the listening functions
	http.HandleFunc("/api/Signup", handleSignup)
    http.HandleFunc("/api/Login", func(response http.ResponseWriter, request *http.Request) {
        handleLogin(response, request, sessionCookie, cookieChanged)
    })
	http.HandleFunc("/api/Logout", func(response http.ResponseWriter, request *http.Request) {
        handleLogout(response, request, sessionCookie)
    })

	// create server
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
	UpdatingData = true
	backendDatabase.UpdatePantry(CurrentUser, updatedPantry.FoodItem)
	UpdatingData = false

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
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
	UpdatingData = true
	backendDatabase.InsertPantryItemPost(CurrentUser, newItem.FoodItem)
	UpdatingData = false

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleNewDealsStore(w http.ResponseWriter, r *http.Request) {

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
	type Store struct{
		Name string `json:"Name"`
	}
	type DealsStore struct {
		Store Store `json:"store"`
	}
	var storeChange DealsStore

	// unmarshal JSON data
	err = json.Unmarshal(body, &storeChange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// change store selection global var
	UpdatingData = true;
	StoreSelection = storeChange.Store.Name
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleRecommendedRecipesSelect(w http.ResponseWriter, r *http.Request) {

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

	// change store selection global var
	UpdatingData = true;
	RoutingRecipesType = RecommendedRecipes
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleUserRecipesSelect(w http.ResponseWriter, r *http.Request) {

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

	// change store selection global var
	UpdatingData = true;
	RoutingRecipesType = UserRecipes
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleFavoriteRecipesSelect(w http.ResponseWriter, r *http.Request) {

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

	// change store selection global var
	UpdatingData = true;
	RoutingRecipesType = FavoriteRecipes
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleNewUserRecipe(w http.ResponseWriter, r *http.Request) {

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
	var newItem Recipe

	// unmarshal JSON data
	err = json.Unmarshal(body, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// change store selection global var
	UpdatingData = true;
	backendDatabase.WriteNewUserRecipe(CurrentUser, newItem)
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleAddFavorite(w http.ResponseWriter, r *http.Request) {

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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// save ASCII as string
	recipeID := string(body)

	// change store selection global var
	UpdatingData = true;
	backendDatabase.FavoriteRecipe(CurrentUser, recipeID)
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleRemoveFavorite(w http.ResponseWriter, r *http.Request) {

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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// save ASCII as string
	recipeID := string(body)

	// change store selection global var
	UpdatingData = true;
	backendDatabase.UnfavoriteRecipe(CurrentUser, recipeID)
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleNewListItem(w http.ResponseWriter, r *http.Request) {

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
	UpdatingData = true
	backendDatabase.WriteList(newItem.FoodItem, CurrentUser)
	UpdatingData = false

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleRecipeFilters(w http.ResponseWriter, r *http.Request) {

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
	type FilterValues struct {
		MyRecipesCheckbox bool `json:"myRecipesCheckbox"`
		UserRecipesCheckbox bool `json:"userRecipesCheckbox"`
		MdRecipesCheckbox bool `json:"mdRecipesCheckbox"`
	}
	var newFilters FilterValues

	// unmarshal JSON data
	err = json.Unmarshal(body, &newFilters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update the recommended recipes pool of recipes depending on the filters
	UpdatingData = true

	// reset the pool to empty
	RecipesRecommendationPool = nil

	//check each checkbox value and add accordingly
	if newFilters.MyRecipesCheckbox {
		temp, _ := backendDatabase.ReadCurrUserRecipes(CurrentUser)
		RecipesRecommendationPool = append(RecipesRecommendationPool, temp...)
	}
	if newFilters.UserRecipesCheckbox {
		temp, _ := backendDatabase.ReadAllUserRecipes()
		RecipesRecommendationPool = append(RecipesRecommendationPool, temp...)
	}
	if newFilters.MdRecipesCheckbox {
		temp, _ := backendDatabase.ReadJSONRecipes()
		RecipesRecommendationPool = append(RecipesRecommendationPool, temp...)
	}

	UpdatingData = false

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func handleDeleteUserRecipe(w http.ResponseWriter, r *http.Request) {

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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// save ASCII as string
	recipeID := string(body)

	// change store selection global var
	UpdatingData = true;
	backendDatabase.DeleteUserRecipe(recipeID)
	UpdatingData = false;

	// write a successful header
	w.WriteHeader(http.StatusOK)

	// if the header was successful, change the recipe data
	if http.StatusOK == 200 {
		// get new data for routing
		UpdateAllData()
	}

}

func ListenForData(){
	
	// handle the listening functions
	http.HandleFunc("/api/UpdatePantry", func(response http.ResponseWriter, request *http.Request) {
        handlePantryUpdate(response, request)
    })
	http.HandleFunc("/api/NewPantryItem", func(response http.ResponseWriter, request *http.Request) {
        handleNewPantryItem(response, request)
    })
	http.HandleFunc("/api/DealsStore", func(response http.ResponseWriter, request *http.Request) {
        handleNewDealsStore(response, request)
    })
	http.HandleFunc("/api/RecommendedRecipesSelect", func(response http.ResponseWriter, request *http.Request) {
        handleRecommendedRecipesSelect(response, request)
    })
	http.HandleFunc("/api/UserRecipesSelect", func(response http.ResponseWriter, request *http.Request) {
        handleUserRecipesSelect(response, request)
    })
	http.HandleFunc("/api/FavoriteRecipesSelect", func(response http.ResponseWriter, request *http.Request) {
        handleFavoriteRecipesSelect(response, request)
    })
	http.HandleFunc("/api/NewUserRecipe", func(response http.ResponseWriter, request *http.Request) {
        handleNewUserRecipe(response, request)
    })
	http.HandleFunc("/api/AddFavoriteRecipe", func(response http.ResponseWriter, request *http.Request) {
        handleAddFavorite(response, request)
    })
	http.HandleFunc("/api/RemoveFavoriteRecipe", func(response http.ResponseWriter, request *http.Request) {
        handleRemoveFavorite(response, request)
	})
	http.HandleFunc("/api/NewItem", func(response http.ResponseWriter, request *http.Request) {
        handleNewListItem(response, request)
    })
	http.HandleFunc("/api/RecommendedRecipesFilters", func(response http.ResponseWriter, request *http.Request) {
        handleRecipeFilters(response, request)
    })
	http.HandleFunc("/api/DeleteUserRecipe", func(response http.ResponseWriter, request *http.Request) {
        handleDeleteUserRecipe(response, request)
    })
	

	// create server
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
}

// DATA UPDATE FUNCTIONS

func UpdateDealsData(){
	// determine which store to take deals from
	if StoreSelection == "Publix"{
		StoreDeals, _ = backendDatabase.ReadPublixDatabase() 
	} else if StoreSelection == "Walmart"{
		StoreDeals, _ = backendDatabase.ReadWalmartDatabase()
	}

	// lock the deals data
	dataMutex.Lock()

	// set all deals to global variable
	var dealsInterfaceRefresh []interface{}
	dealsInterface = dealsInterfaceRefresh

	// Create a temp food item to store the name of the store
	storeItem := FoodItem{Name: StoreSelection}
	dealsInterface = append(dealsInterface, storeItem)
	for i := 0; i < len(StoreDeals); i++{
		dealsInterface = append(dealsInterface, StoreDeals[i])
	}

	// unlock the data
	dataMutex.Unlock()
}

func UpdatePantryData(){
	// lock the user pantry data
	dataMutex.Lock()

	currUserPantry, _ := backendDatabase.GetUserPantry(CurrentUser.UserName)

	// save all user pantry data to global variable
	var pantryInterfaceRefresh []interface{}
	pantryInterface = pantryInterfaceRefresh
	for i := 0; i < len(currUserPantry.FoodInPantry); i++{
		pantryInterface = append(pantryInterface, currUserPantry.FoodInPantry[i])
	}

	// unlock the data
	dataMutex.Unlock()
}

func UpdateRecipeData(){
	var routingRecipes []Recommendation

	currUserRecipes, _ := backendDatabase.ReadCurrUserRecipes(CurrentUser)
	currUserFavRecipes, _ := backendDatabase.ReadFavoriteRecipes(CurrentUser)
	currUserPantry, _ := backendDatabase.GetUserPantry(CurrentUser.UserName)

	// save all recipes data to global variable
	if RoutingRecipesType == RecommendedRecipes{
		routingRecipes = BestRecipes(currUserPantry, RecipesRecommendationPool, StoreDeals)
	} else if RoutingRecipesType == UserRecipes{
		routingRecipes = AllRecipesWithRelatedItems(currUserPantry, currUserRecipes, StoreDeals)
	} else if RoutingRecipesType == FavoriteRecipes {
		routingRecipes = AllRecipesWithRelatedItems(currUserPantry, currUserFavRecipes, StoreDeals)
	}
	
	// find which recipes are user favorites
	routingRecipes = backendDatabase.FindFavoriteRecipes(CurrentUser, routingRecipes)

	// lock the recipe data
	dataMutex.Lock()

	var recipesInterfaceRefresh []interface{}
	recipesInterface = recipesInterfaceRefresh
	for i := 0; i < len(routingRecipes); i++ {
		// sends recipes, items in recipe, and deals related 
		recipesInterface = append(recipesInterface, routingRecipes[i])
	}

	// unlock the data
	dataMutex.Unlock()
}
 
func UpdateListData(){

	// lock the recipe data
	dataMutex.Lock()

	var listInterfaceRefresh []interface{}
	listInterface = listInterfaceRefresh
	userListOverall, _ := backendDatabase.ReadList(CurrentUser)
	userList := userListOverall.ShoppingList
	for i := 0; i < len(userList); i++ {
		// sends shopping list food item slice, time last updated, and user connected to list
		listInterface = append(listInterface, userList[i])
	}

	// unlock the data
	dataMutex.Unlock()
}

func UpdateUserDashboard(){

	// selects what type of data to send to the user dashboard
	if RoutingUserDashboardData == TopFiveFavoritesEasiestToMake{
		
		// get recipe recommendations
		var routingUserDashboardRecipes []Recommendation
		routingUserDashboardRecipes = ReturnRecipesWithHighestPercentageOfOwnedIngredients(backendDatabase.GetUserPantry(CurrentUser.UserName), backendDatabase.ReadFavoriteRecipes(CurrentUser), 5, StoreDeals)
		routingUserDashboardRecipes = backendDatabase.FindFavoriteRecipes(CurrentUser, routingUserDashboardRecipes) // used to determine if star is darkened in full recipe card

		// lock the recipe data
		dataMutex.Lock()

		var userDashboardInterfaceRefresh []interface{}
		userDashboardInterface = userDashboardInterfaceRefresh
		for i := 0; i < len(routingUserDashboardRecipes); i++ {
			// sends recipes, items in recipe, and deals related 
			userDashboardInterface = append(userDashboardInterface, routingUserDashboardRecipes[i])
		}

		// unlock the data
		dataMutex.Unlock()

	} else if RoutingUserDashboardData == TopTenFavoritesEasiestToMake{
		
		// get recipe recommendations
		var routingUserDashboardRecipes []Recommendation
		routingUserDashboardRecipes = ReturnRecipesWithHighestPercentageOfOwnedIngredients(backendDatabase.GetUserPantry(CurrentUser.UserName), backendDatabase.ReadFavoriteRecipes(CurrentUser), 10, StoreDeals)
		routingUserDashboardRecipes = backendDatabase.FindFavoriteRecipes(CurrentUser, routingUserDashboardRecipes) // used to determine if star is darkened in full recipe card

		// lock the recipe data
		dataMutex.Lock()

		var userDashboardInterfaceRefresh []interface{}
		userDashboardInterface = userDashboardInterfaceRefresh
		for i := 0; i < len(routingUserDashboardRecipes); i++ {
			// sends recipes, items in recipe, and deals related 
			userDashboardInterface = append(userDashboardInterface, routingUserDashboardRecipes[i])
		}

		// unlock the data
		dataMutex.Unlock()
	}
	
}


func UpdateAllData(){
	// wait if any data is being altered
	for UpdatingData {}

	// updates all data that's being routed
	UpdateUserDashboard() // first because loaded first
	UpdatePantryData()
	UpdateRecipeData()
	UpdateDealsData()
	UpdateListData()
		
}

// SHUTDOWN FUNCTIONS

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

