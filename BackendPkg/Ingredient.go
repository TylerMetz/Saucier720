package BackendPkg

type Ingredient struct {
	Item              FoodItem
	Amount            float64
	UnitOfMeasurement string
}

func CalculatePrice(food Ingredient) float64 {
	// Take in food item and return its price
	panic("placeholder")
}
