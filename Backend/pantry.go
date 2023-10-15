package main

type Ingredient struct {
	Name 		string 	`json:"Name`
	FoodType 	string 	`json:"FoodType`
	SaleDetails string 	`json:"SaleDetails`
	Quantity 	int 	`json:"Quantity`
}

type Pantry struct {
	Ingredients []Ingredient
}
