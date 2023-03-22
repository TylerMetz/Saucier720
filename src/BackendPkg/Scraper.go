package BackendPkg

import (
	"fmt"
	"runtime"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Scraper struct {
	Store                GroceryStore
	TimeLastDealsScraped time.Time
	DealsHTML            string
	InventoryHTML        string
}

func (s *Scraper) Scrape() {

	// calls function based on store
	if s.Store.Name == "Publix" {
		s.PublixScrapeDeals()
	} else if s.Store.Name == "Walmart" {
		s.WalmartScrapeDeals()
	}

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
		err = inputBox.SendKeys(s.Store.ZipCode)
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
		if !alternateLayoutThree{
			err = inputBoxTwo.SendKeys(s.Store.ZipCode)
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
		} else{
			inputBoxThree, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div.search-container > form > input[type=search]")
			err = inputBoxThree.SendKeys("32601")
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

	html, err := wd.PageSource()
	if err != nil {
		panic(err)
	}

	// saves html from page
	s.DealsHTML = html
	//fmt.Println("Deals Scraped Successfully!")

}

// will scrape entire publix inventory
func (s *Scraper) PublixScrapeInventory() {

	// init chrome driver
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService("src/SeleniumDrivers/chromedriver_mac64/chromedriver", 9515, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

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
	err = wd.Get("https://www.publix.com/d/all-categories")
	if err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("opened inventory page")

	// input desired zipcode
	inputBoxNewPage, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div > form > input[type=search]")
	if err != nil {
		panic(err)
	}
	err = inputBoxNewPage.SendKeys(s.Store.ZipCode)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("zip inputed")
	}

	// search for stores
	searchStoreButtonNewPage, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div.search-container > form > button")
	if err != nil {
		panic(err)
	}
	err = searchStoreButtonNewPage.Click()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("search button pressed")
	}
	time.Sleep(20 * time.Second) // wait for page to load

	// select the first store from the results list
	chooseStoreResultNewPage, err := wd.FindElement(selenium.ByCSSSelector, "#\\31 560 > div > div > div.buttons-links > div.p-button-group__wrapper.buttons-wrapper > div > button")
	if err != nil {
		panic(err)
	}
	err = chooseStoreResultNewPage.Click()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("store selected")
	}
	time.Sleep(10 * time.Second) // wait for page to load

	// sets InventoryHTML as first page's html
	html, err := wd.PageSource()
	if err != nil {
		panic(err)
	}
	// adds this page inventory to the html
	s.InventoryHTML = html

	fmt.Println(s.InventoryHTML) //used to check if correct page

	morePages := true
	// loop until all pages are taken in
	for morePages {
		nextPageButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div.search-results-super-container.v4.mar-top-md.search-page-content > div > div.search-content-column > div.card-loader.search-results-section > div:nth-child(2) > nav.pagination.mobile-only.condensed > button:nth-child(3)")
		err = nextPageButton.Click()
		if err != nil {
			// no more pages to load
			morePages = false
		} else {
			// takes in html from page
			html, err = wd.PageSource()
			if err != nil {
				panic(err)
			}
			// adds this page inventory to the html
			s.InventoryHTML += html
			fmt.Println("Page Done!")
			time.Sleep(1 * time.Second)
		}
	}

}

func (s *Scraper) WalmartScrapeDeals() {
	// will scrape entire Walmart inventory
}

func (s *Scraper) WalmartScrapeInventory() {
	// will scrape entire Walmart inventory
}
