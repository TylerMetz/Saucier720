package Backend

type FoodItem struct {
	name      string
	storeCost float64
	onSale    bool
	salePrice float64
}

// generates new FoodItem with defaults set to 0 values, used for pantry items since only name matters
func NewPantryItem(name string) *FoodItem {

	item := FoodItem{name: name}
	item.storeCost = 0
	item.onSale = false
	item.salePrice = 0

	return &item
}
