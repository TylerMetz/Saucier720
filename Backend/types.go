package main

import (
	"time"
	"fmt"
)
	

type Account struct {
	UserName   	string 		`json:"UserName"`
	FirstName  	string 		`json:"FirstName"`
	LastName   	string 		`json:"LastName"`
	Email      	string 		`json:"Email"`
	Password   	string 		`json:"Password"`
	DateJoined 	time.Time
}

type SignupRequest struct {
	UserName   	string 		`json:"UserName"`
	FirstName  	string 		`json:"FirstName"`
	LastName   	string 		`json:"LastName"`
	Email      	string 		`json:"Email"`
	Password   	string 		`json:"Password"`
}

type SignupResponse struct {
	Response	string 		`json:"Response"`
}

type PantryRequest struct {
	UserName   	string 		`json:"UserName"`
}

type PantryResponse struct { 
	Pantry 		[]Ingredient `json:"Pantry"`
}

type LoginRequest struct{
	UserName	string 		`json:"UserName"`
	Password 	string 		`json:"Password"`
}

type LoginResponse struct {
	Cookie		string 		`json:"Cookie"`
}

type LogoutRequest struct { 
	UserName	string 		`json:"UserName"`
}

type LogoutResponse struct { 
	Response	string 		`json:"Response"`
}

type RecipeFilter struct {
	UserCreatedRecipes 	bool `json:"UserCreatedRecipes"`
	MealDealzRecipes 	bool `json:"MealDealzRecipes"`
	SelfCreatedRecipes 	bool `json:"SelfCreatedRecipes"` 	
}

type RecipesRequest struct {
	UserName   		string `json:"UserName"`
	RecipeFilter 	RecipeFilter `json:"RecipeFilter"`
}

type FavoriteRecipesRequest struct {
	UserName   	string 		`json:"UserName"`
}

type RecipesResponse struct {
	R 			RecomendedRecipes `json:"RecomendedRecipes"`
}

type DealsRequest struct { 
	Zipcode 	int 		`json:"Zipcode"`
}

type DealsByStoreRequest struct { 
	StoreName 	string 		`json:"StoreName"`
	Zipcode 	int 		`json:"Zipcode"`
}

type DealsResponse struct {
	Deals 		[]Ingredient `json:"Deals"`
}

type ListRequest struct {
	UserName 	string 		`json:"UserName"`
}

type ListResponse struct {
	List 		[]Ingredient `json:"List"`
}

type PostPantryRequest struct {
	UserName 	string 		`json:"UserName"`
	Ingredient 	Ingredient 	`json:"Ingredient"`
}

type PostPantryResponse struct { 
	Response 	string 		`json:"Response"`
}

type PostRecipeRequest struct {
	UserName 	string 		`json:"UserName"`
	Recipe 		Recipe 		`json:"Recipe"`
}

type PostRecipeResponse struct {
	Response 	string 		`json:"Response"`
}

type PostListRequest struct {
	UserName 	string 		`json:"UserName"`
	Ingredient 	Ingredient 	`json:"Ingredient"`
}

type PostListResponse struct { 
	Response 	string 		`json:"Response"`
}

type PostCookieRequest struct { 
	UserName 		string 	`json:"UserName"`
}

type PostFavoriteRequest struct { 
	UserName 		string 	`json:"UserName"`
	RecipeID 		int 	`json:"RecipeID"`
}

type PostFavoriteResponse struct {
	Response 	string 		`json:"Response"`
}

type PostCookieResponse struct { 
	Response 	string 		`json:"Response"`
}

type DeletePantryRequest struct {
	UserName	string		`json:"UserName"`
	Ingredient	Ingredient	`json:"Ingredient"`
}

type DeletePantryResponse struct {
	Response	string		`json:"Response"`
}

type DeleteListRequest struct {
	UserName	string 		`json:"UserName"`
	Ingredient	Ingredient	`json:"Ingredient"`
}

type DeleteListResponse struct {
	Response	string		`json:"Response"`
}

type DeleteFavoriteRequest struct {
	UserName	string		`json:"UserName"`
	RecipeID	int			`json:"RecipeID"`
}

type DeleteFavoriteResponse struct {
	Response	string		`json:"Response"`
}

type DeleteRecipeRequest struct {
	UserName	string		`json:"UserName"`
	RecipeID	int			`json:"RecipeID"`
}

type DeleteRecipeResponse struct {
	Response	string		`json:"Response"`
}

type UpdatePantryRequest struct {
	UserName	string		`json:"UserName"`
	Pantry 		Pantry		`json:"Pantry"`
}

type UpdatePantryResponse struct {
	Response	string		`json:"Response"`
}

type UpdateRecipeRequest struct { 
	UserName	string		`json:"UserName"`
	Recipe		Recipe		`json:"Recipe"`
}

type UpdateRecipeResponse struct { 
	Response	string		`json:"Response"`
}

type UpdateListRequest struct { 
	UserName	string		`json:"UserName"`
	List		Pantry		`json:"List"`
}

type UpdateListResponse struct { 
	Response	string		`json:"Response"`
}

type Ingredient struct {
	Name 		string 		`json:"Name"`
	FoodType 	string 		`json:"FoodType"`
	SaleDetails string 		`json:"SaleDetails"`
	Quantity 	int 		`json:"Quantity"`
}

type Pantry struct {
	Ingredients []Ingredient
}

func NewAccount(userName, firstName, lastName, email, password string) (*Account, error){
	return &Account{
		UserName: userName,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		DateJoined: time.Now(),
	}, nil
}

func CreateCookieForUser(userName string) (string, error){
	//create hash from username and time and store in database and THEN have it expire after 7 days
	uniqueToken := fmt.Sprintf("%s-%s", userName, time.Now().String())

	return uniqueToken, nil
}
