package BackendPkg

import (
	"time"
)

type List struct{
	ListOwner User 
	TimeUpdated time.Time 
	ShoppingList []FoodItem
}
