package BackendPkg

import (
	"fmt"
	_ "regexp"
	"strings"
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

func (g *GroceryStore) OrganizeDeals(deals string, start, end int) (string, string) {
	// testing to see what the string reads as 'words'
	words := strings.Fields(deals)
	
	newRange := words[48634:len(words)-1]
	var name string
	var deal string
	var newStart int
	var countHelp int
	// Find item name
	for i := 0; i < len(newRange); i++ {
		if(newRange[i] == "loading=\"lazy\""){
			name = strings.Join(newRange[0:i], " ")
			newStart = start + i 
			break
		}
	}
	// Find item deal
	newRange = words[newStart:len(words)-1]
	for i := 0; i < len(newRange); i++ {
		if(newRange[i] == "color--null\">"){
			for j := 0; j < len(newRange); j++{
				if(newRange[i + j] == "</span>"){
					countHelp = j
					break
				} 
			}
			deal = strings.Join(newRange[0:i+countHelp]," ")
			break
		}
	}
	return deal,name

	/*var itemName string 
    if end > len(words) {
        end = len(words)
    }
    if start < 0 || start > end {
        start = 0
    }*/

    //itemName = strings.Join(words[start:end], " ")
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
