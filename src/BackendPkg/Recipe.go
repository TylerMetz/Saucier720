package BackendPkg

type Recipe struct {
	// List of steps
	// List of inegredients
	Name        string
	Servings    int
	Time        int
	Cost        float64
	ServingSize int
}

func CalcTotalPrice( /*Need to decide data structure*/ ) float64 {
	// Calculates total price of a recipe
	panic("placeholder")
}
func CalcServing( /*Need to decide data structure*/ ) float64 {
	// Calculates total price of a recipe given accounting for size
	panic("placeholder")
}
