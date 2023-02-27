package BackendPkg

type FoodItem struct {
	Name        string
	StoreCost   float64
	OnSale      bool
	SaleDetails string
	Quantity int
}

// generates new FoodItem with defaults set to 0 values, used for pantry items since only name matters
func NewPantryItem(name string) FoodItem {

	item := FoodItem{Name: name}
	item.StoreCost = 0
	item.OnSale = false
	item.SaleDetails = ""

	return item
}
