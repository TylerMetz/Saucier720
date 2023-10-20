package main

import (
	"encoding/json"
	"fmt"
	"time"
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

	//GETS
	router.HandleFunc("/Login", makeHTTPHandleFunc(s.handleLogin)) // we need to generate the cookie here ?
	router.HandleFunc("/Pantry", makeHTTPHandleFunc(s.handleGetPantry))
	router.HandleFunc("/Recipes", makeHTTPHandleFunc(s.handleGetRecipes))
	router.HandleFunc("/Recipes/Favorite", makeHTTPHandleFunc((s.handleGetFavRecipes)))
	router.HandleFunc("/Deals", makeHTTPHandleFunc((s.handleGetDeals)))
	router.HandleFunc("/Deals/Store", makeHTTPHandleFunc((s.handleGetDealsByStore)))
	router.HandleFunc("/List", makeHTTPHandleFunc((s.handleGetList)))
	// PUTS WILL BE NEXT
	router.HandleFunc("/Signup", makeHTTPHandleFunc(s.handleSignup)) // we need to generate the cookie here ?
	router.HandleFunc("/NewPantryItem", makeHTTPHandleFunc(s.handlePostPantryIngredient))
	router.HandleFunc("/NewRecipe", makeHTTPHandleFunc(s.handlePostRecipe))
	router.HandleFunc("/NewListIngredient", makeHTTPHandleFunc(s.handlePostList))
	router.HandleFunc("/NewFavoriteRecipe", makeHTTPHandleFunc(s.handlePostFavoriteRecipe))
	router.HandleFunc("/PostCookie", makeHTTPHandleFunc(s.handlePostCookie))
	// THEN DELETE 
	router.HandleFunc("/Logout", makeHTTPHandleFunc(s.handleLogout)) // we delete the cookie here ?
	router.HandleFunc("/DeletePantryIngredient", makeHTTPHandleFunc(s.handleDeletePantryIngredient))
	router.HandleFunc("/DeleteListIngredient", makeHTTPHandleFunc(s.handleDeleteListIngredient))
	router.HandleFunc("/DeleteFavoriteRecipe", makeHTTPHandleFunc(s.handleDeleteFavoriteRecipe))
	router.HandleFunc("/DeleteUserRecipe", makeHTTPHandleFunc(s.handleDeleteUserRecipe))
	// (UPDATES WILL PROB HAPPEN WITH PUTS)
	router.HandleFunc("/UpdatePantry", makeHTTPHandleFunc(s.handleUpdatePantry))
	router.HandleFunc("/UpdateList", makeHTTPHandleFunc(s.handleUpdateList))
	router.HandleFunc("/UpdateRecipe", makeHTTPHandleFunc(s.handleUpdateRecipe))

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

	verify := CheckPassword(s.store, req.UserName, req.Password)
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
	req := new(PantryRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	pantry, err := s.store.GetPantryByUserName(req.UserName)
	if err != nil {
		return err
	}

	resp := PantryResponse {
		Pantry: pantry,
	}

	return WriteJSON(w, http.StatusOK, resp)
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
		mealDealzRecipes, err := s.store.GetRecipesByUserName("MealDealz Classic Recipe")
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

// handleGetFavRecipes
func (s *APIServer) handleGetFavRecipes(w http.ResponseWriter, r *http.Request) error{ 
	req := new(FavoriteRecipesRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	//get fav recipes
	favRecipes, err := s.store.GetFavoriteRecipes(req.UserName)
	if err != nil { 
		fmt.Println("error getting fav recipes")
		return err
	}
	//Get User Pantry
	pantry, err := s.store.GetPantryByUser(req.UserName)
	if err != nil {
		fmt.Println("error getting pantry")
	}

	//getting deals to pass in, i think we could make all of these go func things and have them run concurrently?
	deals, err := s.store.GetDeals()
	if err != nil { 
		fmt.Println("error getting deals")
		return err
	}

	resp := new(RecipesResponse)

	resp.R.Recommendations = ReturnRecipesWithHighestPercentageOfOwnedIngredients(pantry, favRecipes, len(favRecipes), deals)

	return WriteJSON(w, http.StatusOK, resp)
}

// handleGetDeals - we should add a zipcode type to this? or go off the current user's zipcode setting (we also need to implement settings)
func (s *APIServer) handleGetDeals(w http.ResponseWriter, r *http.Request) error { 
	req := new(DealsRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	deals, err := s.store.GetDeals()
	if err != nil { 
		fmt.Println("error getting deals")
		return err
	}

	resp := new(DealsResponse)
	resp.Deals = deals

	return WriteJSON(w, http.StatusOK, resp)
}

//handleGetDealsByStore
func (s *APIServer) handleGetDealsByStore(w http.ResponseWriter, r *http.Request) error { 
	req := new(DealsByStoreRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	deals, err := s.store.GetDealsByStore(req.StoreName)
	if err != nil { 
		fmt.Println("error getting deals")
		return err
	}

	resp := new(DealsResponse)
	resp.Deals = deals

	return WriteJSON(w, http.StatusOK, resp)
}

// handleGetList
func (s *APIServer) handleGetList(w http.ResponseWriter, r *http.Request) error {
	req := new(ListRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	deals, err := s.store.GetShoppingListByUserName(req.UserName)
	if err != nil { 
		fmt.Println("error getting deals")
		return err
	}

	resp := new(ListResponse)
	resp.List = deals

	return WriteJSON(w, http.StatusOK, resp)
}


// POSTS
func (s *APIServer) handlePostPantryIngredient(w http.ResponseWriter, r *http.Request) error { 
	req := new(PostPantryRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	if err := s.store.PostPantryIngredient(req.UserName, req.Ingredient); err != nil {
		return err
	}

	resp := PostPantryResponse{
		Response: "Ingredient Successfully Posted!",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handlePostRecipe(w http.ResponseWriter, r *http.Request) error { 
	req := new(PostRecipeRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.PostRecipe(req.UserName, req.Recipe); err != nil {
		return err
	}

	resp := PostRecipeResponse{
		Response: "Recipe Successfully Posted!",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handlePostList(w http.ResponseWriter, r *http.Request) error { 
	req := new(PostListRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.PostListIngredient(req.UserName, req.Ingredient); err != nil {
		return err
	}

	resp := PostListResponse {
		Response: "Ingredient Successfully Posted!",
	}
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeletePantryIngredient(w http.ResponseWriter, r *http.Request) error {
	req := new(DeletePantryRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.DeletePantryIngredient(req.UserName, req.Ingredient); err != nil{
		return err
	}

	resp := DeletePantryResponse {
		Response: "Ingredient Successfully Removed From Pantry",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handlePostFavoriteRecipe(w http.ResponseWriter, r *http.Request) error { 
	req := new(PostFavoriteRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.PostFavoriteRecipe(req.UserName, req.RecipeID); err != nil {
		return err
	}

	resp := PostFavoriteResponse {
		Response: "Recipe Successfully Posted!",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handlePostCookie(w http.ResponseWriter, r *http.Request) error { 
	req := new(PostCookieRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	//Generate Cookie Here
	//helper function calling CreateCookieForUser
	cookieToken, err := CreateCookieForUser(req.UserName); if err != nil { 
		fmt.Println("error creating cookie")
		return err
	}

	httpCookie := &http.Cookie{
		Name:     "Jason's Cookie",
		Value:    cookieToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour), // Set expiration to 7 days in the future.
		HttpOnly: true,
	}

	http.SetCookie(w, httpCookie)

	if err := s.store.PostCookieByUserName(req.UserName, cookieToken); err != nil {
		fmt.Println("error posting cookie")
		return err
	}

	resp := PostCookieResponse {
		Response: "Cookie Successfully Posted!",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

// Deletes

func (s *APIServer) handleDeleteListIngredient(w http.ResponseWriter, r *http.Request) error {
	req := new(DeleteListRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.DeleteListIngredient(req.UserName, req.Ingredient); err != nil{
		return err
	}

	resp := DeletePantryResponse {
		Response: "Ingredient Successfully Removed From List",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeleteFavoriteRecipe(w http.ResponseWriter, r *http.Request) error {
	req := new(DeleteFavoriteRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.DeleteFavorite(req.UserName, req.RecipeID); err != nil{
		return err
	}

	resp := DeleteFavoriteResponse {
		Response: "Successfully Removed From Favorites",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeleteUserRecipe(w http.ResponseWriter, r *http.Request) error {
	req := new(DeleteRecipeRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.DeleteRecipe(req.UserName, req.RecipeID); err != nil{
		return err
	}

	resp := DeleteRecipeResponse {
		Response: "Successfully Removed Recipe",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleLogout(w http.ResponseWriter, r *http.Request) error {
	req := new(LogoutRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.DeleteCookieByUserName(req.UserName); err != nil{
		return err
	}

	resp := LogoutResponse {
		Response: "Successfully Logged Out",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleUpdatePantry(w http.ResponseWriter, r *http.Request) error {
	req := new(UpdatePantryRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.UpdatePantryByUserName(req.UserName, req.Pantry); err != nil{
		return err
	}

	resp := UpdatePantryResponse {
		Response: "Successfully Updated Pantry",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleUpdateList(w http.ResponseWriter, r *http.Request) error {
	req := new(UpdateListRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.UpdateListByUserName(req.UserName, req.List); err != nil{
		return err
	}

	resp := UpdateListResponse {
		Response: "Successfully Updated List",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleUpdateRecipe(w http.ResponseWriter, r *http.Request) error {
	req := new(UpdateRecipeRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{ 
		return err
	}

	if err := s.store.UpdateRecipeByUserName(req.UserName, req.Recipe); err != nil{
		return err
	}

	resp := UpdateRecipeResponse {
		Response: "Successfully Updated Recipe",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func CheckPassword(s Storage, username, password string) bool {
	dbPassword, _ := s.GetPasswordByUserName(username)
	return password == dbPassword
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	// do we need those other CORS headers?
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
