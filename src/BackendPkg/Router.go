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

	fmt.Println(body)

	type Ingredient struct {
		FoodItem FoodItem `json:"Ingredient"`
	}

    var newItem Ingredient;
	
    err = json.Unmarshal(body, &newItem)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	fmt.Println(newItem)

	funcDatabase := Database{
		Name: "func db",
	}
    funcDatabase.InsertPantryItemPost(currUser, newItem.FoodItem)

    w.WriteHeader(http.StatusOK)
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

    w.WriteHeader(http.StatusOK)
}

func ListenForAllPosts(currUser User){
	// all listen functions go in here

	// listens for new user
	go ListenNewUser()

	// listens for new pantry item
	ListenPantry(currUser)

}


