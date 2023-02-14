package BackendPkg

import (
	"fmt"
	"github.com/gocolly/colly"
)

type GroceryStore struct {
	Inventory []FoodItem
	Name      string
	Address   string
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

	c := colly.NewCollector(colly.AllowedDomains("https://www.publix.com"))

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
	c.OnHTML("div.p-card p-savings-card p-card--interactive", func(e *colly.HTMLElement) {
		tempName := ""
		tempSaleDetails := ""
		e.ForEach("div.content-wrapper", func(i int, h *colly.HTMLElement) {
			h.ForEach("div.top-section", func(i int, t *colly.HTMLElement) {
				t.ForEach("div.title-wrapper", func(i int, m *colly.HTMLElement) {
					m.ForEach("span.p-text paragraph-md normal context--default color--null line-clamp title", func(i int, l *colly.HTMLElement) {
						tempName = l.Text //name of product
					})
				})
				t.ForEach("span.p-savings-badge savings-badge bogo", func(i int, m *colly.HTMLElement) {
					m.ForEach("div.p-savings-badge__text", func(i int, l *colly.HTMLElement) {
						l.ForEach("span.p-text paragraph-sm strong context--default color--null", func(i int, n *colly.HTMLElement) {
							tempSaleDetails = n.Text
						})
					})
				})
			})
			g.UpdateInventory(tempName, tempSaleDetails)
		})

	})

	// visit and scrape deals
	c.Visit("https://www.publix.com/savings/weekly-ad/view-all")

}

// Webscrape deals from store
// Make sure it clears each week
func (g *GroceryStore) ScrapeInventory() {

}
func (g GroceryStore) DisplayDeals() {
	// Display
}
