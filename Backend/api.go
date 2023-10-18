package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store Storage
}

func NewMealDealzServer(listenAddr string, store Storage) *APIServer{
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/Signup", makeHTTPHandleFunc(s.handleSignup))
	router.HandleFunc("/Login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/Pantry", makeHTTPHandleFunc(s.handleGetPantry))
	router.HandleFunc("/Recipes", makeHTTPHandleFunc(s.handleGetRecipes))
	

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleSignup(w http.ResponseWriter, r *http.Request) error {
	req := new(SignupRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	account, err := NewAccount(req.UserName, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil{
		return err
	}
	if err := s.store.PostSignup(account); err != nil {
		return err
	}

	resp := SignupResponse{
		Response: "User Successfully Created!",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	req := new(LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	verify := s.store.CheckPassword(req.UserName, req.Password)
	if(verify){
		resp := LoginResponse{
			// WE ACTUALLY NEED TO GENERATE A COOKIE
			Cookie: "GeneratedCookie",
		}
		return WriteJSON(w, http.StatusOK, resp)
	}

	return WriteJSON(w, http.StatusBadRequest, 0)
}

func (s *APIServer) handleGetPantry(w http.ResponseWriter, r *http.Request) error {
	pantry, err := s.store.GetPantry()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, pantry)
}

func (s *APIServer) handleGetRecipes(w http.ResponseWriter, r *http.Request) error{
	req := new(RecipesRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	var recipes []Recipe

	//get recipes based on filters
	if req.RecipeFilter.UserCreatedRecipes {
		//get user created recipes
		userCreatedRecipes, err := s.store.GetUserCreatedRecipes()
		if err != nil { 
			fmt.Println("error getting user created recipes")
			return err
		}
		// add to recipes array
		recipes = append(recipes, userCreatedRecipes...)
	}
	if req.RecipeFilter.MealDealzRecipes {
		//get meal dealz recipes
		mealDealzRecipes, err := s.store.GetRecipesByUserName("MealDealz Classic Recipes")
		if err != nil { 
			fmt.Println("error getting mealdealz classic recipes")
			return err
		}
		// add to recipes array
		recipes = append(recipes, mealDealzRecipes...)
	}
	if req.RecipeFilter.SelfCreatedRecipes {	
		//get self created recipes
		selfCreatedRecipes, err := s.store.GetRecipesByUserName(req.UserName)
		// add to recipes array
		if err != nil { 
			fmt.Println("error getting own users recipes")
			return err
		}
		// add to recipes array
		recipes = append(recipes, selfCreatedRecipes...)
	}

	//Get User Pantry
	pantry, err := s.store.GetPantryByUser(req.UserName)
	if err != nil {
		fmt.Println("error getting pantry")
	}

	//rate them based on recomendation funcs

	recomendedRecipes := ReturnRecipesWithHighestPercentageOfOwnedIngredients(pantry, recipes, 50, []Ingredient{})

	//return recipes request

	resp := new(RecipesResponse)

	resp.R = RecomendedRecipes{ 
		Recommendations: recomendedRecipes,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
