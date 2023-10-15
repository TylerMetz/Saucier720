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

type Recipe struct {
	RecipeId	string `json:"RecipeId"`
	Title		string `json:"Title"`
	Ingredients []string `json:"Ingredients"`
	Instructions []string `json:"Instructions"`
	UserName	string	`json:"UserName"`
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