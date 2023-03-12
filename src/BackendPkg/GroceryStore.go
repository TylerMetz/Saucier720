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

func (g *GroceryStore) OrganizeDeals(deals string, start, end int) {
	// testing to see what the string reads as 'words'
	words := strings.Fields(deals)
	newRange := words[start : len(words)-1]
	count := 0
	var name string
	var deal string
	newStart := start
	var countHelp int


	for {
	
		var nextStep int = 0
		// Find item name
		// Most of the names end after we find the loadinglazy string
		for i := 0; i < len(newRange); i++ {
			if newRange[i] == "loading=\"lazy\"" {
				name = strings.Join(newRange[0:i], " ")
				newStart = newStart + i
				break
			}
		}
		// Find item deal
		// the deal is usually between color--null and span 
		newRange = words[newStart : len(words)-1]
		for i := 0; i < len(newRange); i++ {
			if newRange[i] == "color--null\">" {
				for j := 0; j < len(newRange); j++ {
					if newRange[i+j] == "</span>" {
						countHelp = j
						break
					}
				}
				deal = strings.Join(newRange[i:i+countHelp], " ")
				newStart = newStart + i + countHelp
				newRange = words[newStart : len(words)-1]
				break
			}
		}
	
		// clean up
		deal = deal[14:]
		name = name[5:]
		name = name[:len(name)-1]

		if(name == "Paper Coupon"){
			break
		}
	
		// find next starting point
		for i:= 0; i < len(newRange); i++ {
			if newRange[i] == "data-v-cfc9b7ee=\"\""{
				nextStep++
			}
			if(nextStep == 4){
				newStart = newStart + i
				newRange = words[newStart + 1: len(words)-1]
				break
			}
			
		}
		fmt.Println(name)
		fmt.Println(deal)
		count++
	}
	fmt.Print(count)
	// Next steps: 
	// Need to make it recursive or loop until we reach the end of the list 
	// Once it consistently works, must add each item into the inventory 
	// Push to database after 
	//return deal, name
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
