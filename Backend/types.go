package main

import "time"

type Account struct {
	UserName   string `json:"UserName`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	DateJoined time.Time
}