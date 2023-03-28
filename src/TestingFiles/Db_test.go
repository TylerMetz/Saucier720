package main

import (
	"BackendPkg"
	"testing"
	"time"
)

func TestThree(t *testing.T){
	testingDatabase := BackendPkg.Database{
		Name: "Testing Database",
	}

	testFoodItem := BackendPkg.FoodItem{
		Name:        "peanut butter",
		StoreCost:   369.99,
		OnSale:      true,
		SaleDetails: "BOGO",
		Quantity:    10,
	}
	testFoodItem2 := BackendPkg.FoodItem{
		Name:        "jelly",
		StoreCost:   1.0,
		OnSale:      false,
		SaleDetails: "N/A",
		Quantity:    30,
	}
	testFoodItem3 := BackendPkg.FoodItem{
		Name:        "bread",
		StoreCost:   10.69,
		OnSale:      true,
		SaleDetails: "$2 for 2",
		Quantity:    2,
	}
	testUserFoodSlice := []BackendPkg.FoodItem{testFoodItem, testFoodItem2, testFoodItem3}

	testUser := BackendPkg.User{
		FirstName: "Sam",
		LastName: "Forstot",
		Email: "samuel@gmail.com",
		UserName: "SameHatesBigWordsXXX",
		Password: "ILoveJess420",
		UserPantry: BackendPkg.Pantry{
			FoodInPantry: testUserFoodSlice,
			TimeLastUpdated: time.Now(),
		},
	}

	testingDatabase.StoreUserDatabase(testUser)
	testingDatabase.StoreUserPantry(testUser)
	
	returnPantry := testingDatabase.GetUserPantry(testUser.UserName)
	
	if(testUser.UserPantry.TimeLastUpdated == returnPantry.TimeLastUpdated){
		if(len(testUser.UserPantry.FoodInPantry) != len(returnPantry.FoodInPantry)){
			t.Errorf("Pantries do not match")
		} else {
			for i := range returnPantry.FoodInPantry {
				if testUser.UserPantry.FoodInPantry[i] != returnPantry.FoodInPantry[i] {
					t.Errorf("Pantries do not match")
				}
			}
		}
	}

	
}


