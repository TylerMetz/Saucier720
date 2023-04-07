package BackendPkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"time"

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

func ListenPantry() {

	// Listens and Serves pantry
    http.HandleFunc("/api/NewPantryItem", PantryItemPostResponse)
    log.Fatal(http.ListenAndServe(":8083", nil))

}

func PantryItemPostResponse(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var newItem FoodItem;

    err := json.NewDecoder(r.Body).Decode(&newItem)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	funcDatabase := Database{
		Name: "func db",
	}
    funcDatabase.InsertPantryItemPost(newItem)

    w.WriteHeader(http.StatusOK)
}

func ListenNewUser() {

	// Listens and Serves pantry
    http.HandleFunc("/api/Signup", NewUserResponse)
    log.Fatal(http.ListenAndServe(":8085", nil))

}

func NewUserResponse(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var newUser User;

    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	funcDatabase := Database{
		Name: "func db",
	}
    funcDatabase.StoreUserDatabase(newUser)

    w.WriteHeader(http.StatusOK)
}

func ListenForAllPosts(){
	// all listen functions go in here

	// listens for new user
	go ListenNewUser()

	// listens for new pantry item
	ListenPantry()

}


