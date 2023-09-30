package BackendPkg

type FoodItem struct {
	Name      string `json:"Name"`
	FoodType string `json:"FoodType"`
	SaleDetails string `json:"SaleDetails"`
	Quantity int `json:"Quantity"`
}
