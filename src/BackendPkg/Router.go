package BackendPkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"sync"
	"context"
	//"strings"
	//"bytes"
)

// global mutex
var dataMutex sync.Mutex

// global wait group
var wg sync.WaitGroup

// global var to be routed
var pantryFoodInterface []interface{}
var dealsFoodInterface []interface{}
var recipesFoodInterface []interface{}

// global Servers list
var Servers []*http.Server

type Router struct {
	Name string
}

func (t *Router) RoutPantry(endLink string, port string, ctx context.Context) {
	// create a new context with cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// creates new router
	route := mux.NewRouter()
	route.HandleFunc(endLink, t.sendResponsePantry).Methods("GET")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// define handler
	handler := c.Handler(route)

	// create server to add to global var
	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	// append the server to the global list
	Servers = append(Servers, server)

	// increment the WaitGroup counter
    wg.Add(1)
	// listen and serve until context is cancelled
	func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
}

func (t *Router) RoutDeals(endLink string, port string, ctx context.Context) {
	// create a new context with cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// creates new router
	route := mux.NewRouter()
	route.HandleFunc(endLink, t.sendResponseDeals).Methods("GET")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// define handler
	handler := c.Handler(route)

	// create server to add to global var
	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	// append the server to the global list
	Servers = append(Servers, server)

    // increment the WaitGroup counter
    wg.Add(1)
	// listen and serve until context is cancelled
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
}

func (t *Router) RoutRecipes(endLink string, port string, ctx context.Context) {
	// create a new context with cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// creates new router
	route := mux.NewRouter()
	route.HandleFunc(endLink, t.sendResponseRecipes).Methods("GET")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// define handler
	handler := c.Handler(route)

	// create server to add to global var
	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	// append the server to the global list
	Servers = append(Servers, server)

	// increment the WaitGroup counter
    wg.Add(1)
	// listen and serve until context is cancelled
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
}

func (t *Router) sendResponsePantry(response http.ResponseWriter, request *http.Request) {

	jsonResponse, jsonError := json.Marshal(pantryFoodInterface)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse)) // used to test

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func (t *Router) sendResponseDeals(response http.ResponseWriter, request *http.Request) {

	jsonResponse, jsonError := json.Marshal(dealsFoodInterface)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse)) // used to test

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func (t *Router) sendResponseRecipes(response http.ResponseWriter, request *http.Request) {

	jsonResponse, jsonError := json.Marshal(recipesFoodInterface)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse)) // used to test

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func ListenPantry(currUser User, ctx context.Context) {
	// create a new context with cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Listens and Serves pantry
	route := mux.NewRouter()
	route.HandleFunc("/api/NewPantryItem", func(w http.ResponseWriter, r *http.Request) {
        PantryItemPostResponse(w, r, currUser)
    })

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// creates handler
	handler := c.Handler(route)

	// creates server to be appended to global list
	server := &http.Server{Addr: ":8083", Handler: handler}
	Servers = append(Servers, server)

	// increment the WaitGroup counter
	wg.Add(1)

	// listens and serves in a new goroutine
	func() {
		defer wg.Done() // decrement the counter when this goroutine is done
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for cancellation signal
	<-ctx.Done()

}

func PantryItemPostResponse(w http.ResponseWriter, r *http.Request, currUser User) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	type Ingredient struct {
		FoodItem FoodItem `json:"ingredient"`
	}

	var newItem Ingredient

	err = json.Unmarshal(body, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	funcDatabase := Database{
		Name: "func db",
	}
	funcDatabase.InsertPantryItemPost(currUser, newItem.FoodItem)

	w.WriteHeader(http.StatusOK)

	if http.StatusOK == 200 {
		d := Database{
			Name: "func db",
		}

		var pantryFoodInterfaceRefresh []interface{}
		pantryFoodInterface = pantryFoodInterfaceRefresh
		UpdateData(d, currUser)
	}

}

func ListenUpdatedPantry(currUser User, ctx context.Context) {
	// create a new context with cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Listens and Serves pantry
	route := mux.NewRouter()
	route.HandleFunc("/api/UpdatePantry", func(w http.ResponseWriter, r *http.Request) {
        UpdatedPantryResponse(w, r, currUser)
    }).Methods("POST")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// creates handler
	handler := c.Handler(route)
	fmt.Println("Handler created")

	// creates server to be appended to global list
	server := &http.Server{Addr: ":8086", Handler: handler}
	Servers = append(Servers, server)

	// increment the WaitGroup counter
	wg.Add(1)

	// listens and serves in a new goroutine
	go func() {
		defer wg.Done() // decrement the counter when this goroutine is done
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		
	}()

	// wait for cancellation signal
	<-ctx.Done()

}

func UpdatedPantryResponse(w http.ResponseWriter, r *http.Request, currUser User) {

	fmt.Println("Request method:", r.Method)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	type Ingredient struct {
		FoodItem []FoodItem `json:"pantry"`
	}

	var updatedPantry Ingredient

	err = json.Unmarshal(body, &updatedPantry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	funcDatabase := Database{
		Name: "func db",
	}

	funcDatabase.UpdatePantry(currUser, updatedPantry.FoodItem)


	w.WriteHeader(http.StatusOK)

	if http.StatusOK == 200 {
		d := Database{
			Name: "func db",
		}

		var pantryFoodInterfaceRefresh []interface{}
		pantryFoodInterface = pantryFoodInterfaceRefresh
		UpdateData(d, currUser)
	}

}

func ListenNewUser(ctx context.Context) {

	// create a new context with cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Listens and Serves New User
	route := mux.NewRouter()
	route.HandleFunc("/api/Signup", NewUserResponse).Methods("POST")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// creates handler
	handler := c.Handler(route)

	// creates server to be appended to global list
	server := &http.Server{Addr: ":8085", Handler: handler}
	Servers = append(Servers, server)

	// increment the WaitGroup counter
	wg.Add(1)

	// listens and serves in a new goroutine
	go func() {
		defer wg.Done() // decrement the counter when this goroutine is done
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for cancellation signal
	<-ctx.Done()

}

func NewUserResponse(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	type UserResponse struct {
		User User `json:"user"`
	}

	var newUser UserResponse

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	funcDatabase := Database{
		Name: "func db",
	}

	funcDatabase.StoreUserDatabase(newUser.User)

}

func ListenLogin(sessionCookie* string, cookieChanged* bool, ctx context.Context) {
	// create a new context with cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Listens and Serves New User
	route := mux.NewRouter()
	route.HandleFunc("/api/Login", func(w http.ResponseWriter, r *http.Request) {
        NewLoginResponse(w, r, sessionCookie, cookieChanged)
    }).Methods("POST")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	// creates handler
	handler := c.Handler(route)

	// creates server to be appended to global list
	server := &http.Server{Addr: ":8084", Handler: handler}
	Servers = append(Servers, server)

	// increment the WaitGroup counter
	wg.Add(1)

	// listens and serves in a new goroutine
	func() {
		defer wg.Done() // decrement the counter when this goroutine is done
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for cancellation signal
	<-ctx.Done()

}

func NewLoginResponse(w http.ResponseWriter, r *http.Request, sessionCookie *string, cookieChanged *bool) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	type LoginUser struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	var currUser LoginUser

	err = json.Unmarshal(body, &currUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
		// readInCookie, _ := r.Cookie("sessionID")
		*sessionCookie = cookie.Value

	}

	// fmt.Println(*sessionCookie) // print for testing

}

func ListenForAllPosts(currUser User, cookie string, ctx context.Context){
	// all listen functions go in here

	// listens for new user
	go ListenNewUser(ctx)

	// listens for pantry updates (quantity and deletes)
	go ListenUpdatedPantry(currUser, ctx)

	// listens for new pantry item
	ListenPantry(currUser, ctx)

}

func RoutUserPantry(d Database, u User, ctx context.Context){
	
	// read from .db file and output test user's pantry to frontend

	// lock the data
	dataMutex.Lock()

	var pantryFoodInterfaceRefresh []interface{}
	pantryFoodInterface = pantryFoodInterfaceRefresh
	for i := 0; i < len(d.GetUserPantry(u.UserName).FoodInPantry); i++{
		pantryFoodInterface = append(pantryFoodInterface, d.GetUserPantry(u.UserName).FoodInPantry[i])
	}

	// unlock the data
	dataMutex.Unlock()

	// test router
	programRouter := Router{
		Name:             "testRouter",
	}
	programRouter.RoutPantry("/api/Pantry", ":8080", ctx)
}

func RoutWeeklyDeals(d Database, ctx context.Context){

		// read from .db file and output test user's pantry to frontend
		for i := 0; i < len(d.ReadPublixDatabase()); i++{
			dealsFoodInterface = append(dealsFoodInterface, d.ReadPublixDatabase()[i])
		}
		// test router
		programRouter := Router{
			Name:             "testRouter",
		}
		programRouter.RoutDeals("/api/Deals", ":8081", ctx)

}

func RoutRecommendedRecipes(d Database, currUser User, ctx context.Context){

	userRecList := BestRecipes(d.GetUserPantry(currUser.UserName), d.ReadRecipes(), d.ReadPublixDatabase())
	var recipesFoodInterfaceRefresh []interface{}
	recipesFoodInterface = recipesFoodInterfaceRefresh
	for i := 0; i < len(userRecList); i++ {
		// sends recipes, items in recipe, and deals related 
		recipesFoodInterface = append(recipesFoodInterface, userRecList[i])
	}

	programRouter := Router{
		Name: "testRouter",
	}
	programRouter.RoutRecipes("/api/Recipes", ":8082", ctx)
}

func RoutAllData(d Database, currUser User, ctx context.Context){

	// routs Eddie's pantry, lol
	go RoutUserPantry(d, currUser, ctx) 

	// routs reccommended recipes to recipes page
	go RoutRecommendedRecipes(d, currUser, ctx)

	// routs deals to deals page
	RoutWeeklyDeals(d, ctx)
}

func UpdateData(d Database, u User) {

	// Lock the mutex to update the data
	dataMutex.Lock()
	// Update the global variable with the updated data
	for i := 0; i < len(d.GetUserPantry(u.UserName).FoodInPantry); i++ {
		pantryFoodInterface = append(pantryFoodInterface, d.GetUserPantry(u.UserName).FoodInPantry[i])
	}
	// Unlock the mutex
	dataMutex.Unlock()
}

func deleteAllCookies(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}

func ShutdownServers() {
    for _, server := range Servers {
        // gracefully shut down the server
        if err := server.Shutdown(context.Background()); err != nil {
            log.Fatal("Server shutdown failed:", err)
        }
    }
    // wait for all the servers to shut down before returning
    wg.Wait()
	Servers = nil
	return
}
