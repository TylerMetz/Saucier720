package BackendPkg

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/rs/cors"
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

	// enables alternate hosts for CORS
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:4200"},
        AllowCredentials: true,
    })

	log.Println("Listening...")
	handler := c.Handler(route)
	http.ListenAndServe(":8080", handler) 
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

