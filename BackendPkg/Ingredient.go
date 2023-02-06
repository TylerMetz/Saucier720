package BackendPkg

type Ingredient struct {
	item              FoodItem
	amount            float64
	unitOfMeasurement string
}

func CalculatePrice(food Ingredient) float64 {
	// Take in food item and return its price
	panic("placeholder")
}
