package BackendPkg

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
    "github.com/gorilla/mux"
)

type Router struct{
	Name string
	ItemsToBeEncoded []interface{}
}

func (t *Router) Rout() { 
    
	// creates new router
	route := mux.NewRouter()
	route.HandleFunc("/", t.sendResponse)
	log.Println("Listening...")
	http.ListenAndServe(":8080", route) 
}

func (t *Router) sendResponse(response http.ResponseWriter, request *http.Request) {
	
	jsonResponse, jsonError := json.Marshal(t.ItemsToBeEncoded)
	
	if jsonError != nil {
	fmt.Println("Unable to encode JSON")
	}
	
	fmt.Println(string(jsonResponse))
	
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

