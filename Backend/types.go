package main

import "time"

type Account struct {
	UserName   	string `json:"UserName"`
	FirstName  	string `json:"FirstName"`
	LastName   	string `json:"LastName"`
	Email      	string `json:"Email"`
	Password   	string `json:"Password"`
	DateJoined 	time.Time
}

type SignupRequest struct {
	UserName   	string `json:"UserName"`
	FirstName  	string `json:"FirstName"`
	LastName   	string `json:"LastName"`
	Email      	string `json:"Email"`
	Password   	string `json:"Password"`
}

type SignupResponse struct {
	Response	string `json:"Response"`
}

type LoginRequest struct{
	UserName	string `json:"UserName"`
	Password 	string `json:"Password"`
}

type LoginResponse struct {
	Cookie		string `json:"Cookie"`
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
	UserName   		string `json:"UserName"`
}

type RecipesResponse struct {
	R RecomendedRecipes `json:"RecomendedRecipes"`
}

type Ingredient struct {
	Name 		string 	`json:"Name"`
	FoodType 	string 	`json:"FoodType"`
	SaleDetails string 	`json:"SaleDetails"`
	Quantity 	int 	`json:"Quantity"`
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
