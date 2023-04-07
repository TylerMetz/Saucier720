package BackendPkg

type FoodItem struct {
	Name      string `json:"Name"`
	StoreCost float64 `json:"StoreCost"`
	OnSale    bool `json:"OnSale"`
	SalePrice float64 `json:"SalePrice"`
	SaleDetails string `json:"SaleDetails"`
	Quantity int `json:"Quantity"`
}

// generates new FoodItem with defaults set to 0 values, used for pantry items since only name matters
func NewPantryItem(name string) FoodItem {

	item := FoodItem{Name: name}
	item.StoreCost = 0
	item.OnSale = false
	item.SalePrice = 0

	return item
}
