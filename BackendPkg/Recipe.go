package Backend

type Recipe struct{
	// List of steps
	// List of inegredients
	name string 
	servings int
	time int
	cost float64
}

func CalcTotalPrice(/*Need to decide data structure*/) float64{
	// Calculates total price of a recipe
}
func CalcServing(/*Need to decide data structure*/) float64{
	// Calculates total price of a recipe given accounting for size 
}