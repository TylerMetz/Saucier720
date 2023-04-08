package BackendPkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"time"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"sync"
)

// global mutex
var dataMutex sync.Mutex

// global var to be routed
var testFoodInterface []interface{}

type Router struct {
	Name             string
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

	jsonResponse, jsonError := json.Marshal(testFoodInterface)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse)) // used to test

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func ListenPantry(currUser User) {

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

	handler := c.Handler(route)
    log.Fatal(http.ListenAndServe(":8083", handler))

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

    var newItem Ingredient;
	
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

	if http.StatusOK == 200{
		d := Database{
			Name: "func db",
		} 
		
		var testFoodInterfaceRefresh []interface{}
		testFoodInterface = testFoodInterfaceRefresh
		UpdateData(d, currUser)
		fmt.Println("new item unlocked!")
	}

}

func ListenNewUser() {

	// Listens and Serves New User
	route := mux.NewRouter()
    route.HandleFunc("/api/Signup", NewUserResponse).Methods("POST")

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	handler := c.Handler(route)
    log.Fatal(http.ListenAndServe(":8085", handler))

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

func ListenForAllPosts(currUser User){
	// all listen functions go in here

	// listens for new user
	go ListenNewUser()

	// listens for new pantry item
	ListenPantry(currUser)

}

func RoutUserPantry(d Database, u User){
	
	// read from .db file and output test user's pantry to frontend
	for{
		// lock the data
		dataMutex.Lock()

		for i := 0; i < len(d.GetUserPantry(u.UserName).FoodInPantry); i++{
			testFoodInterface = append(testFoodInterface, d.GetUserPantry(u.UserName).FoodInPantry[i])
		}

		// unlock the data
		dataMutex.Unlock()

		// test router
		programRouter := Router{
			Name:             "testRouter",
		}
		programRouter.Rout("/api/Pantry", ":8080")
	}
}

func RoutWeeklyDeals(d Database){
	
	// read from .db file and output test user's pantry to frontend
	var testFoodInterface []interface{}
	for i := 0; i < len(d.ReadPublixDatabase()); i++{
		testFoodInterface = append(testFoodInterface, d.ReadPublixDatabase()[i])
	}
	// test router
	programRouter := Router{
		Name:             "testRouter",
	}
	programRouter.Rout("/api/Deals", ":8081")
}

func RoutRecommendedRecipes(d Database, currUser User){

	userRecList := BestRecipes(d.GetUserPantry(currUser.UserName), d.ReadRecipes(), d.ReadPublixDatabase())
	var testFoodInterface []interface{}
	for i := 0; i < len(userRecList); i++{
		testFoodInterface = append(testFoodInterface, userRecList[i].R)
		testFoodInterface = append(testFoodInterface, "Pantry Data:")
		for j := 0; j < len(userRecList[i].ItemsInPantry); j++{
			testFoodInterface = append(testFoodInterface, userRecList[i].ItemsInPantry[j].Name)
		}
		testFoodInterface = append(testFoodInterface, "Publix Data:")
		for k := 0; k < len(userRecList[i].ItemsOnSale); k++{
			testFoodInterface = append(testFoodInterface, userRecList[i].ItemsOnSale[k].Name)
		}
	}

	programRouter := Router{
		Name:             "testRouter",
	}
	programRouter.Rout("/api/Recipes", ":8082")
}

func RoutAllData(d Database, currUser User){

	// routs Eddie's pantry, lol
	go RoutUserPantry(d, currUser)

	// routs deals to deals page
	go RoutWeeklyDeals(d)

	// routs reccommended recipes to recipes page
	RoutRecommendedRecipes(d, currUser)
}

func UpdateData(d Database, u User) {

    // Lock the mutex to update the data
    dataMutex.Lock()
    // Update the global variable with the updated data
	for i := 0; i < len(d.GetUserPantry(u.UserName).FoodInPantry); i++{
		testFoodInterface = append(testFoodInterface, d.GetUserPantry(u.UserName).FoodInPantry[i])
	}
    // Unlock the mutex
    dataMutex.Unlock()
}
