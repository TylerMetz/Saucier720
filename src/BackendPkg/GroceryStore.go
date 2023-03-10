package BackendPkg

import (
	"fmt"
)

type GroceryStore struct {
	Inventory []FoodItem
	Name      string
	ZipCode   string
}

func (g *GroceryStore) UpdateInventory(name_, saleDetails_ string) {
	for i := 0; i < len(g.Inventory); i++ {
		if g.Inventory[i].Name == name_ {
			g.Inventory[i].Name = name_
			g.Inventory[i].SaleDetails = saleDetails_
			i = len(g.Inventory) // end loop
		}
	}
}

func (g *GroceryStore) DisplaySales() {
	for i := 0; i < len(g.Inventory); i++ {
		fmt.Println("Item: ", g.Inventory[i].Name, ", Sale Details: ", g.Inventory[i].SaleDetails)
	}
}

// Take in Inventory list & change by reference
func (g *GroceryStore) ScrapeDeals() {
	// find and print all of the deals

	// scrapes deals
	/*
		Name:
		div.p-card p-savings-card p-card--interactive
			div.content-wrapper
				div.top-section
					div.title-wrapper
						span.p-text paragraph-md normal context--default color--null line-clamp title
							TEXT

		Sale Details:
		div.p-card p-savings-card p-card--interactive
			div.content-wrapper
				div.top-section
					span.p-savings-badge savings-badge bogo
						div.p-savings-badge__text
							span.p-text paragraph-sm strong context--default color--null
								TEXT
	*/

}

// Webscrape deals from store
// Make sure it clears each week
func (g *GroceryStore) ScrapeInventory() {

}
func (g GroceryStore) DisplayDeals() {
	// Display
}

func (g *GroceryStore) PassItemPage(itemURL string) {
	//parses through pages and using Colly web scrapes the inventory
}
