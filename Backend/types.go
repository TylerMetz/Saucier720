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