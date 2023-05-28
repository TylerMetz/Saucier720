package BackendPkg

import (
	"fmt"
	"runtime"
	"time"
	"strings"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var UserZipCodePlaceholder string = "32601"

type Scraper struct {
	TimeLastDealsScraped time.Time
	PublixDeals			[]FoodItem
	WalmartDeals 		[]FoodItem
}

func (s *Scraper) Scrape() {
	// scrapes data for all stores

	// scrapes walmart deals
	s.WalmartScrapeDeals()
	fmt.Println("Walmart Deals Scraped!")

	// scrapes publix deals
	s.PublixScrapeDeals()
	fmt.Println("Publix Deals Scraped!")

	// saves current time to ref later
	s.TimeLastDealsScraped = time.Now()
}

func (s *Scraper) PublixScrapeDeals() {
	// init chrome driver
	opts := []selenium.ServiceOption{}
	if runtime.GOOS == "windows" {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_win32/chromedriver.exe", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	} else {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_mac64/chromedriver", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	}

	// init headless browser
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
		},
	}
	caps.AddChrome(chromeCaps)

	// run headless chrome browser
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// open desired page
	err = wd.Get("https://www.publix.com/savings/weekly-ad/view-all")
	if err != nil {
		panic(err)
	}

	// select "Choose a Store button"
	chooseStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div.savings-content-wrapper.skeleton-spacer > div > div.savings-container.full-bleed > div > div > button > span")
	if err != nil {
		panic(err)
	}
	err = chooseStoreButton.Click()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println("landing page button selected")
	}

	alternateLayout := false
	// input desired zipcode
	inputBox, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div > form > input[type=search]")
	if err != nil {
		alternateLayout = true //checks if the alternate windowed version is running (runs on some networks with the window on the right side of the screen)
	}
	if alternateLayout == false {
		time.Sleep(10 * time.Second)
		err = inputBox.SendKeys(UserZipCodePlaceholder)
		if err != nil {
			panic(err)
		} else {
			//fmt.Println("zip inputed")
		}

		// search for stores
		searchStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div.search-container > form > button")
		if err != nil {
			panic(err)
		}
		err = searchStoreButton.Click()
		if err != nil {
			panic(err)
		} else {
			//fmt.Println("search button pressed")
		}
	} else {
		alternateLayoutThree := false
		// sets the input box of the alternate window as the input box
		inputBoxTwo, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div > form > input[type=search]")
		if err != nil {
			alternateLayoutThree = true //checks if the alternate windowed version is running (runs on some networks with the window on the right side of the screen)
		}
		if !alternateLayoutThree {
			time.Sleep(10 * time.Second)
			err = inputBoxTwo.SendKeys(UserZipCodePlaceholder)
			// search for stores
			searchStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div > form > button")
			if err != nil {
				panic(err)
			}
			err = searchStoreButton.Click()
			if err != nil {
				panic(err)
			} else {
				//fmt.Println("search button pressed")
			}
		} else {
			inputBoxThree, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div.search-container > form > input[type=search]")
			time.Sleep(10 * time.Second)
			err = inputBoxThree.SendKeys(UserZipCodePlaceholder)
			if err != nil {
				fmt.Println("not third layout")
			}
			// search for stores
			searchStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div.search-container > form > button")
			if err != nil {
				panic(err)
			}
			err = searchStoreButton.Click()
			if err != nil {
				panic(err)
			} else {
				//fmt.Println("search button pressed")
			}
		}
	}
	time.Sleep(20 * time.Second) // wait for page to load

	// select the first store from the results list
	chooseStoreResult, err := wd.FindElement(selenium.ByCSSSelector, "#\\31 560 > div > div > div.buttons-links > div.p-button-group__wrapper.buttons-wrapper > div > button")
	if err != nil {
		panic(err)
	}
	err = chooseStoreResult.Click()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println("store selected")
	}
	time.Sleep(5 * time.Second) // wait for page to load

	// switch to list view
	listViewButton, err := wd.FindElement(selenium.ByCSSSelector, "#listing")
	if err != nil {
		panic(err)
	}
	err = listViewButton.Click()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println("list view enabled")
	}

	time.Sleep(20 * time.Second) // wait for page to load

	// Keeps hitting "Load More" button until all of the data is loaded
	moreLoadingNeeded := true
	// triggers page to load more
	for moreLoadingNeeded {
		loadMoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div.savings-content-wrapper > div > div.savings-container > div.card-loader.savings-content.search-results-section.-coupons > div.button-container > button")
		if err != nil {
			moreLoadingNeeded = false
		} else {
			_ = loadMoreButton.Click()
			time.Sleep(3 * time.Second)
		}
	}

	// stores expanded deals page html
	html, err := wd.PageSource()
	if err != nil {
		panic(err)
	}

	// organize and store publix deals as a slice of FoodItems
	s.PublixDeals = s.OrganizePublixDeals(html)

}

func (s *Scraper) WalmartScrapeDeals() {
	// init chrome driver
	opts := []selenium.ServiceOption{}
	if runtime.GOOS == "windows" {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_win32/chromedriver.exe", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	} else {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_mac64/chromedriver", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	}

	// init headless browser
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
		},
	}
	caps.AddChrome(chromeCaps)

	// run headless chrome browser
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Open the Walmart webpage
	if err := wd.Get("https://www.walmart.com/browse/grocery-deals/c2hlbGZfaWQ6MjQ1NTI0NQieie?affinityOverride=default"); err != nil {
		panic(err)
	}

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// make pages array
	pages := make([]string, 0)

	// stores first deals page html
	
	html, err := wd.PageSource()
	if err != nil {
		panic(err)
	}
	pages = append(pages, html)

	// Keeps hitting next page button until all of the data is loaded
	moreLoadingNeeded := true
	// triggers page to load more
	for moreLoadingNeeded {
		loadMoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#maincontent > main > div > div > div > div > div:nth-child(8) > nav > ul > li:nth-child(4) > a")
		if err != nil {
			moreLoadingNeeded = false
		} else {
			_ = loadMoreButton.Click()
			time.Sleep(3 * time.Second)
			// stores next deals page html
			html2, err := wd.PageSource()
			if err != nil {
				panic(err)
			}
			pages = append(pages, html2)
		}
	}

	fmt.Println(pages)

	//s.WalmartDeals = ???
}

func FindStart(phrase, s string) (string) {
    i := strings.Index(s, phrase)
    if i == -1 {
        return ""
    }
    return s[i:]
}

func (s *Scraper) OrganizePublixDeals(deals string) []FoodItem {
	// testing to see what the string reads as 'words'
	words := strings.Fields(deals)
	newRange := words[0 : len(words)-1]
	//count := 0
	var name string
	var deal string
	newStart := 0
	var countHelp int
	dealSlice := make([]FoodItem, 0)

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
		// need to check for interesting deals and clean them into their own spot 
		//bigDeal := strings.Fields(name)
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
		/*fmt.Println(name)
		fmt.Println(deal)
		count++*/
		
		item := FoodItem{
				Name:        name,
				StoreCost:   100,
				OnSale:      true,
				SaleDetails: deal,
				Quantity:    0,
		}

		dealSlice = append(dealSlice, item)
	}
	//fmt.Print(count)
	// Once it consistently works, must add each item into the inventory 
	// Push to database after 
	// Cleaning up edge case
	dealSlice = dealSlice[1:]
	return dealSlice
}
