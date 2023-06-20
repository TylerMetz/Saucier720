package BackendPkg

import (
	"time"
)

type List struct{
	listOwner User 
	TimeAdded time.Time 
	ShoppingList []FoodItem
}
